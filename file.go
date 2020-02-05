package nogo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// File represents a file's content and it's FileInfo.
// It implements the http.File interface.
type File struct {
	//File os.File
	FileInfo FileInfo
	Content  []byte
	DirInfos []FileInfo
}

// ensure File corresponds to http.File.
var _ http.File = File{}

// Close implements the io.Closer interface.
// It does not do anything at the moment.
func (f File) Close() error {
	return nil
}

// Seek implements the io.Seeker interface.
func (f File) Seek(offset int64, whence int) (int64, error) {
	return bytes.NewReader(f.Content).Seek(offset, whence)
}

// Read implements the io.Reader interface.
func (f File) Read(b []byte) (int, error) {
	return bytes.NewReader(f.Content).Read(b)
}

// Readdir corresponds to the http.file interface.
func (f File) Readdir(count int) ([]os.FileInfo, error) {
	var infos []os.FileInfo
	for _, info := range f.DirInfos {
		infos = append(infos, info)
	}
	if count > len(f.DirInfos) {
		return infos, io.EOF
	}
	return infos, nil
}

// Stat corresponds to the http.file interface.
func (f File) Stat() (os.FileInfo, error) {
	return f.FileInfo, nil
}

// convert os.FileInfo to nogo.FileInfo
func fileInfoFromOS(info os.FileInfo) FileInfo {
	return FileInfo{
		FileName:    info.Name(),
		FileSize:    info.Size(),
		FileMode:    info.Mode(),
		FileModTime: info.ModTime(),
		FileIsDir:   info.IsDir(),
	}
}

// LoadFile creates a new nogo file.
func LoadFile(name string) (File, error) {
	res := File{}
	f, err := os.Open(name)
	if err != nil {
		return res, fmt.Errorf("%v: could not open file named %v", err, name)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return res, fmt.Errorf("%v: could not get file info from file named %v", err, name)
	}
	res.FileInfo = fileInfoFromOS(info)
	if info.IsDir() {
		infos, err := f.Readdir(0)
		if err != nil {
			return res, fmt.Errorf("%v: could not readdir for file named %v", err, name)
		}
		for _, info := range infos {
			res.DirInfos = append(res.DirInfos, fileInfoFromOS(info))
		}
	} else {
		res.Content, err = ioutil.ReadFile(f.Name())
		if err != nil {
			return res, fmt.Errorf("%v: could not read file named %v", err, name)
		}
	}
	return res, err
}
