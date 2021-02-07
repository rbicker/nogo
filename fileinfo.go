package nogo

import (
	"os"
	"time"
)

// FileInfo corresponds to os.FileInfo.
type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

// ensure FileInfo corresponds to os.FileInfo.
var _ os.FileInfo = FileInfo{}

func NewFileInfo(info os.FileInfo) FileInfo {
	return FileInfo{
		name:    info.Name(),
		size:    info.Size(),
		mode:    info.Mode(),
		modTime: info.ModTime(),
		isDir:   info.IsDir(),
	}
}

// Name returns the file's name.
func (info FileInfo) Name() string {
	return info.name
}

// Size returns the file's size.
func (info FileInfo) Size() int64 {
	return info.size
}

// Mode returns the file's mode.
func (info FileInfo) Mode() os.FileMode {
	return info.mode
}

// ModTime returns the file's modification time.
func (info FileInfo) ModTime() time.Time {
	return info.modTime
}

// IsDir returns true if the file is a directory.
func (info FileInfo) IsDir() bool {
	return info.isDir
}

// Sys returns nil.
func (info FileInfo) Sys() interface{} {
	return nil
}
