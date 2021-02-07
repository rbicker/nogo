package nogo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// File represents a file's Content and it's FileInfo.
// It implements the http.File interface.
type File struct {
	//File os.File
	FileInfo     FileInfo
	Content      []byte
	DirInfos     []FileInfo
	reader       *bytes.Reader
	readDirIndex int
}

// ensure File corresponds to http.File.
var _ http.File = &File{}

// Close implements the io.Closer interface.
// It does not do anything at the moment.
func (f *File) Close() error {
	return nil
}

// Seek implements the io.Seeker interface.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	if f.reader == nil {
		f.reader = bytes.NewReader(f.Content)
	}
	return f.reader.Seek(offset, whence)
}

// Read implements the io.Reader interface.
func (f *File) Read(b []byte) (int, error) {
	if f.reader == nil {
		f.reader = bytes.NewReader(f.Content)
	}
	return f.reader.Read(b)
}

// Readdir corresponds to the http.file interface.
func (f *File) Readdir(n int) ([]os.FileInfo, error) {
	if !f.FileInfo.IsDir() {
		return nil, os.ErrInvalid
	}
	var infos []os.FileInfo
	// If n <= 0, Readdir returns all the FileInfo from the directory in a single slice.
	// In this case, if Readdir succeeds (reads all the way to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory, Readdir returns the FileInfo read until that point and a non-nil error.
	if n <= 0 {
		// reset index to 0 and set the number to include all the directories files.
		f.readDirIndex = 0
		n = len(f.DirInfos)
	}
	for ; f.readDirIndex < f.readDirIndex+n && f.readDirIndex < len(f.DirInfos); f.readDirIndex++ {
		infos = append(infos, &f.DirInfos[f.readDirIndex])
	}
	if f.readDirIndex >= len(f.DirInfos)-1 {
		return infos, io.EOF
	}
	return infos, nil
}

// Stat corresponds to the http.file interface.
func (f *File) Stat() (os.FileInfo, error) {
	return &f.FileInfo, nil
}

// LoadFile creates a new nogo file.
func LoadFile(name string) (File, error) {
	res := File{}
	f, err := os.Open(name)
	if err != nil {
		return res, fmt.Errorf("could not open file named %s: %w", name, err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return res, fmt.Errorf("could not get file FileInfo from file named %s: %w", name, err)
	}
	res.FileInfo = NewFileInfo(info)
	if info.IsDir() {
		infos, err := f.Readdir(0)
		if err != nil {
			return res, fmt.Errorf("could not readdir for file named %s: %w", name, err)
		}
		for _, info := range infos {
			res.DirInfos = append(res.DirInfos, NewFileInfo(info))
		}
	} else {
		res.Content, err = ioutil.ReadFile(f.Name())
		if err != nil {
			return res, fmt.Errorf("could not read file named %s: %w", name, err)
		}
	}
	return res, err
}
