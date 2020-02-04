package nogo

import (
	"os"
	"time"
)

// FileInfo corresponds to os.FileInfo.
type FileInfo struct {
	FileName    string
	FileSize    int64
	FileMode    os.FileMode
	FileModTime time.Time
	FileIsDir   bool
}

// ensure FileInfo corresponds to os.FileInfo.
var _ os.FileInfo = FileInfo{}

// Name returns the file's name.
func (info FileInfo) Name() string {
	return info.FileName
}

// Size returns the file's size.
func (info FileInfo) Size() int64 {
	return info.FileSize
}

// Mode returns the file's mode.
func (info FileInfo) Mode() os.FileMode {
	return info.FileMode
}

// ModTime returns the file's modification time.
func (info FileInfo) ModTime() time.Time {
	return info.FileModTime
}

// IsDir returns true if the file is a directory.
func (info FileInfo) IsDir() bool {
	return info.FileIsDir
}

// Sys returns nil.
func (info FileInfo) Sys() interface{} {
	return nil
}
