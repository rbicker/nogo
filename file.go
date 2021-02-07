package nogo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// File represents a file's Content and it's FileInfo.
// It implements the http.File interface.
type File struct {
	//File os.File
	FileInfo *FileInfo
	Content  []byte
	DirInfos []FileInfo
	reader   *bytes.Reader
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
		log.Println("creating new reader")
		f.reader = bytes.NewReader(f.Content)
	}
	return f.reader.Read(b)
}

// Readdir corresponds to the http.file interface.
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	var infos []os.FileInfo
	for _, info := range f.DirInfos {
		infos = append(infos, &info)
	}
	if count > len(f.DirInfos) {
		return infos, io.EOF
	}
	return infos, nil
}

// Stat corresponds to the http.file interface.
func (f *File) Stat() (os.FileInfo, error) {
	return f.FileInfo, nil
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
			res.DirInfos = append(res.DirInfos, *NewFileInfo(info))
		}
	} else {
		res.Content, err = ioutil.ReadFile(f.Name())
		if err != nil {
			return res, fmt.Errorf("could not read file named %s: %w", name, err)
		}
	}
	return res, err
}
