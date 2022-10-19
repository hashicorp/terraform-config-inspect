package tfconfig

import (
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

// FS is an interface used by [LoadModuleFromFilesystem].
//
// Unfortunately this package implemented a draft version of the io/fs.FS
// API before it was finalized and so this interface is not compatible with
// the final design. To use this package with the final filesystem API design,
// use [WrapFS] to wrap a standard filesystem implementation so that it
// implements this interface.
type FS interface {
	Open(name string) (File, error)
	ReadFile(name string) ([]byte, error)
	ReadDir(dirname string) ([]os.FileInfo, error)
}

// File represents an open file in [FS].
type File interface {
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

// wrapFS is a rather regrettable adapter from the standard library filesystem
// interfaces to the one we have designed in this package, since we adopted
// a draft of that API before it was finalized and the final is incompatible.
type wrapFS struct {
	wrapped fs.FS
}

// WrapFS wraps a standard library filesystem implementation so that it
// implements this package's own (slightly-incompatible) interface [FS].
func WrapFS(wrapped fs.FS) FS {
	return wrapFS{wrapped}
}

func (wfs wrapFS) Open(name string) (File, error) {
	return wfs.wrapped.Open(name)
}

func (wfs wrapFS) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(wfs.wrapped, name)
}

func (wfs wrapFS) ReadDir(dirname string) ([]os.FileInfo, error) {
	entries, err := fs.ReadDir(wfs.wrapped, dirname)
	var ret []os.FileInfo
	if len(entries) != 0 {
		ret = make([]os.FileInfo, len(entries))
		for i, entry := range entries {
			ret[i] = wrapFileInfoDirEntry{entry}
		}
	}
	return ret, err
}

type wrapFileInfoDirEntry struct {
	wrapped fs.DirEntry
}

func (d wrapFileInfoDirEntry) IsDir() bool {
	return d.wrapped.IsDir()
}

func (d wrapFileInfoDirEntry) ModTime() time.Time {
	// this package doesn't actually care about modification times,
	// so we don't need to implement this.
	panic("unimplemented")
}

func (d wrapFileInfoDirEntry) Mode() fs.FileMode {
	// this package doesn't actually care about file modes,
	// so we don't need to implement this.
	panic("unimplemented")
}

func (d wrapFileInfoDirEntry) Name() string {
	return d.wrapped.Name()
}

func (d wrapFileInfoDirEntry) Size() int64 {
	// this package doesn't actually care about file sizes,
	// so we don't need to implement this.
	panic("unimplemented")
}

func (d wrapFileInfoDirEntry) Sys() any {
	return nil
}

type osFs struct{}

func (fs *osFs) Open(name string) (File, error) {
	return os.Open(name)
}

func (fs *osFs) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

func (fs *osFs) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

// NewOsFs provides a basic implementation of FS for an OS filesystem
func NewOsFs() FS {
	return &osFs{}
}
