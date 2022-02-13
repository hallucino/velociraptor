package ntfs

// This is an accessor which represents an NTFS filesystem
/*
   Velociraptor - Hunting Evil
   Copyright (C) 2019 Velocidex Innovations.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
// A Raw NTFS accessor for disks.

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Velocidex/ordereddict"
	errors "github.com/pkg/errors"
	ntfs "www.velocidex.com/golang/go-ntfs/parser"
	"www.velocidex.com/golang/velociraptor/accessors"
	"www.velocidex.com/golang/velociraptor/accessors/ntfs/readers"
	"www.velocidex.com/golang/velociraptor/json"
	"www.velocidex.com/golang/velociraptor/third_party/cache"
	"www.velocidex.com/golang/velociraptor/uploads"
	"www.velocidex.com/golang/velociraptor/utils"
	"www.velocidex.com/golang/vfilter"
)

const (
	// Scope cache tag for the NTFS parser
	NTFSFileSystemTag = "_NTFS"
)

type AccessorContext struct {
	mu sync.Mutex

	// The context is reference counted and will only be destroyed
	// when all users have closed it.
	refs      int
	cached_fd *os.File
	is_closed bool // Keep track if the file needs to be re-opened.
	ntfs_ctx  *ntfs.NTFSContext

	path_listing *cache.LRUCache
}

func (self *AccessorContext) GetNTFSContext() *ntfs.NTFSContext {
	self.mu.Lock()
	defer self.mu.Unlock()

	return self.ntfs_ctx
}

func (self *AccessorContext) IsClosed() bool {
	self.mu.Lock()
	defer self.mu.Unlock()

	return self.is_closed
}

func (self *AccessorContext) IncRef() {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.refs++
}

func (self *AccessorContext) Close() {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.refs--
	if self.refs <= 0 {
		self.cached_fd.Close()
		self.is_closed = true
	}
}

type NTFSFileInfo struct {
	info       *ntfs.FileInfo
	_full_path *accessors.OSPath
}

func (self *NTFSFileInfo) IsDir() bool {
	return self.info.IsDir
}

func (self *NTFSFileInfo) Size() int64 {
	return self.info.Size
}

func (self *NTFSFileInfo) Data() *ordereddict.Dict {
	result := ordereddict.NewDict().
		Set("mft", self.info.MFTId).
		Set("name_type", self.info.NameType).
		Set("fn_btime", self.info.FNBtime).
		Set("fn_mtime", self.info.FNMtime)
	if self.info.ExtraNames != nil {
		result.Set("extra_names", self.info.ExtraNames)
	}

	return result
}

func (self *NTFSFileInfo) Name() string {
	return self.info.Name
}

func (self *NTFSFileInfo) Mode() os.FileMode {
	var result os.FileMode = 0755
	if self.IsDir() {
		result |= os.ModeDir
	}
	return result
}

func (self *NTFSFileInfo) ModTime() time.Time {
	return self.info.Mtime
}

func (self *NTFSFileInfo) FullPath() string {
	return self._full_path.String()
}

func (self *NTFSFileInfo) OSPath() *accessors.OSPath {
	return self._full_path
}

func (self *NTFSFileInfo) Btime() time.Time {
	return self.info.Btime
}

func (self *NTFSFileInfo) Mtime() time.Time {
	return self.info.Mtime
}

func (self *NTFSFileInfo) Ctime() time.Time {
	return self.info.Ctime
}

func (self *NTFSFileInfo) Atime() time.Time {
	return self.info.Atime
}

// Not supported
func (self *NTFSFileInfo) IsLink() bool {
	return false
}

func (self *NTFSFileInfo) GetLink() (*accessors.OSPath, error) {
	return nil, errors.New("Not implemented")
}

type NTFSFileSystemAccessor struct {
	scope vfilter.Scope

	// The delegate accessor we use to open the underlying volume.
	accessor string
	device   string

	root *accessors.OSPath
}

func NewNTFSFileSystemAccessor(
	scope vfilter.Scope, device, accessor string) *NTFSFileSystemAccessor {
	root_path, _ := accessors.NewGenericOSPath("")
	device = strings.TrimSuffix(device, "\\")
	return &NTFSFileSystemAccessor{
		scope:    scope,
		accessor: accessor,
		device:   device,
		root:     root_path,
	}
}

func (self NTFSFileSystemAccessor) New(scope vfilter.Scope) (
	accessors.FileSystemAccessor, error) {
	// Create a new cache in the scope.
	return &NTFSFileSystemAccessor{
		scope:    scope,
		device:   self.device,
		accessor: self.accessor,
		root:     self.root,
	}, nil
}

func (self *NTFSFileSystemAccessor) getRootMFTEntry(ntfs_ctx *ntfs.NTFSContext) (
	*ntfs.MFT_ENTRY, error) {
	return ntfs_ctx.GetMFT(5)
}

func (self NTFSFileSystemAccessor) ParsePath(path string) (
	*accessors.OSPath, error) {
	return accessors.NewWindowsNTFSPath(path)
}

func (self *NTFSFileSystemAccessor) ReadDir(path string) (
	res []accessors.FileInfo, err error) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("PANIC %v\n", r)
			err, _ = r.(error)
		}
	}()

	// Normalize the path
	fullpath, err := self.ParsePath(path)
	if err != nil {
		return nil, err
	}
	result := []accessors.FileInfo{}

	device := self.device
	accessor := self.accessor
	if device == "" {
		pathspec := fullpath.PathSpec()
		device = pathspec.GetDelegatePath()
		accessor = pathspec.GetDelegateAccessor()
	}

	ntfs_ctx, err := readers.GetNTFSContext(self.scope, device, accessor)
	if err != nil {
		return nil, err
	}

	root, err := ntfs_ctx.GetMFT(5)
	if err != nil {
		return nil, err
	}

	// Open the device path from the root.
	dir, err := Open(self.scope, root, ntfs_ctx, device, accessor, fullpath)
	if err != nil {
		return nil, err
	}

	// Only process each mft id once.
	seen := []int64{}
	in_seen := func(id int64) bool {
		for _, i := range seen {
			if i == id {
				return true
			}
		}
		return false
	}

	// List the directory.
	for _, node := range dir.Dir(ntfs_ctx) {
		node_mft_id := int64(node.MftReference())
		if in_seen(node_mft_id) {
			continue
		}

		seen = append(seen, node_mft_id)

		node_mft, err := ntfs_ctx.GetMFT(node_mft_id)
		if err != nil {
			continue
		}
		// Emit a result for each filename
		for _, info := range ntfs.Stat(ntfs_ctx, node_mft) {
			// Skip . files - they are pretty useless.
			if info == nil || info.Name == "." || info.Name == ".." {
				continue
			}
			result = append(result, &NTFSFileInfo{
				info:       info,
				_full_path: fullpath.Append(info.Name),
			})
		}
	}
	return result, nil
}

type readAdapter struct {
	sync.Mutex

	info   accessors.FileInfo
	reader ntfs.RangeReaderAt
	pos    int64
}

func (self *readAdapter) Ranges() []uploads.Range {
	result := []uploads.Range{}
	for _, rng := range self.reader.Ranges() {
		result = append(result, uploads.Range{
			Offset:   rng.Offset,
			Length:   rng.Length,
			IsSparse: rng.IsSparse,
		})
	}
	return result
}

func (self *readAdapter) Read(buf []byte) (res int, err error) {
	self.Lock()
	defer self.Unlock()

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("PANIC %v\n", r)
			err, _ = r.(error)
		}
	}()

	res, err = self.reader.ReadAt(buf, self.pos)
	self.pos += int64(res)

	// If ReadAt is unable to read anything it means an EOF.
	if res == 0 {
		return res, io.EOF
	}

	return res, err
}

func (self *readAdapter) ReadAt(buf []byte, offset int64) (int, error) {
	self.Lock()
	defer self.Unlock()
	self.pos = offset

	return self.reader.ReadAt(buf, offset)
}

func (self *readAdapter) Close() error {
	return nil
}

func (self *readAdapter) Seek(offset int64, whence int) (int64, error) {
	self.Lock()
	defer self.Unlock()

	self.pos = offset
	return self.pos, nil
}

func (self *NTFSFileSystemAccessor) Open(path string) (res accessors.ReadSeekCloser, err error) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("PANIC %v\n", r)
			err, _ = r.(error)
		}
	}()

	fullpath, err := self.ParsePath(path)
	if err != nil {
		return nil, err
	}

	device := self.device
	accessor := self.accessor
	if device == "" {
		pathspec := fullpath.PathSpec()
		device = pathspec.GetDelegatePath()
		accessor = pathspec.GetDelegateAccessor()
	}

	// We dont want to open a subpath of the filesyste, instead we
	// special case this as openning the raw device.
	if len(fullpath.Components) == 0 {
		accessor, err := accessors.GetAccessor(accessor, self.scope)
		if err != nil {
			return nil, err
		}

		file, err := accessor.Open(device)
		if err != nil {
			return nil, err
		}

		reader, err := ntfs.NewPagedReader(
			utils.ReaderAtter{file}, 0x1000, 10000)
		if err != nil {
			return nil, err
		}

		return utils.NewReadSeekReaderAdapter(reader), nil

	}

	ntfs_ctx, err := readers.GetNTFSContext(self.scope, device, accessor)
	if err != nil {
		return nil, err
	}

	root, err := self.getRootMFTEntry(ntfs_ctx)
	if err != nil {
		return nil, err
	}

	data, err := ntfs.GetDataForPath(ntfs_ctx, fullpath.Path())
	if err != nil {
		return nil, err
	}

	dirname := fullpath.Dirname()
	basename := strings.ToLower(fullpath.Basename())

	dir, err := Open(self.scope, root, ntfs_ctx, device, accessor, dirname)
	if err != nil {
		return nil, err
	}

	for _, info := range ntfs.ListDir(ntfs_ctx, dir) {
		if strings.ToLower(info.Name) == basename {
			return &readAdapter{
				info: &NTFSFileInfo{
					info:       info,
					_full_path: dirname.Append(info.Name),
				},
				reader: data,
			}, nil
		}
	}

	return nil, errors.New("File not found")
}

func (self *NTFSFileSystemAccessor) Lstat(path string) (res accessors.FileInfo, err error) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("PANIC %v\n", r)
			err, _ = r.(error)
		}
	}()

	fullpath, err := self.ParsePath(path)
	if err != nil {
		return nil, err
	}
	device := self.device
	accessor := self.accessor
	if device == "" {
		pathspec := fullpath.PathSpec()
		device = pathspec.GetDelegatePath()
		accessor = pathspec.GetDelegateAccessor()
	}

	ntfs_ctx, err := readers.GetNTFSContext(self.scope, device, accessor)
	if err != nil {
		return nil, err
	}

	root, err := self.getRootMFTEntry(ntfs_ctx)
	if err != nil {
		return nil, err
	}

	dirname := fullpath.Dirname()
	basename := strings.ToLower(fullpath.Basename())
	dir, err := Open(self.scope, root, ntfs_ctx, device, accessor, dirname)
	if err != nil {
		return nil, err
	}
	for _, info := range ntfs.ListDir(ntfs_ctx, dir) {
		if strings.ToLower(info.Name) == basename {
			res := &NTFSFileInfo{
				info:       info,
				_full_path: dirname.Append(info.Name),
			}
			return res, nil

		}
	}

	return nil, errors.New("File not found")
}

// Open the MFT entry specified by a path name. Walks all directory
// indexes in the path to find the right MFT entry.
func Open(scope vfilter.Scope, self *ntfs.MFT_ENTRY,
	ntfs_ctx *ntfs.NTFSContext, device, accessor string,
	filename *accessors.OSPath) (*ntfs.MFT_ENTRY, error) {

	components := filename.Components

	// Path is the relative path from the root of the device we want to list
	// component: The name of the file we want (case insensitive)
	// dir: The MFT entry to search.
	get_path_in_dir := func(path string, component string, dir *ntfs.MFT_ENTRY) (
		*ntfs.MFT_ENTRY, error) {

		key := device + path
		path_cache := GetNTFSPathCache(scope, device, accessor)
		item, pres := path_cache.GetComponentMetadata(key, component)
		if pres {
			return ntfs_ctx.GetMFT(item.MftId)
		}

		lru_map := make(map[string]*CacheMFT)

		// Populate the directory cache with all the mft ids.
		lower_component := strings.ToLower(component)
		for _, idx_record := range dir.Dir(ntfs_ctx) {
			file := idx_record.File()
			name_type := file.NameType().Name
			if name_type == "DOS" {
				continue
			}
			item_name := file.Name()
			mft_id := int64(idx_record.MftReference())

			lru_map[strings.ToLower(item_name)] = &CacheMFT{
				MftId:     mft_id,
				Component: item_name,
				NameType:  name_type,
			}
		}
		path_cache.SetLRUMap(key, lru_map)

		for _, v := range lru_map {
			if strings.ToLower(v.Component) == lower_component {
				return ntfs_ctx.GetMFT(v.MftId)
			}
		}

		return nil, errors.New("Not found")
	}

	// NOTE: This refreshes each parent directory in the LRU.
	directory := self
	path := ""
	for _, component := range components {
		if component == "" {
			continue
		}
		next, err := get_path_in_dir(
			path, component, directory)
		if err != nil {
			return nil, err
		}
		directory = next
		path = path + "\\" + component
	}

	return directory, nil
}

func init() {
	accessors.Register("raw_ntfs", &NTFSFileSystemAccessor{},
		`Access the NTFS filesystem by parsing NTFS structures.`)

	json.RegisterCustomEncoder(&NTFSFileInfo{}, accessors.MarshalGlobFileInfo)
}