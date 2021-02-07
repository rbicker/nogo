package nogo

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"

	"encoding/gob"
)

// map for storing the nogos files as slice of bytes
var nogos = make(map[string][]byte)

// Dir corresponds to http.FileSystem.
type Dir string

// ensure Dir corresponds to http.FileSystem
var _ http.FileSystem = Dir("")

// Get decodes the file with the given name.
func Get(name string) (File, error) {
	f := File{}
	if b, ok := nogos[name]; ok {
		r := bytes.NewReader(b)
		dec := gob.NewDecoder(r)
		if err := dec.Decode(&f); err != nil {
			return f, err
		}
		return f, nil
	}
	return f, os.ErrNotExist
}

// Open returns a http.File based on the given name.
// The function corresponds to http.FileSystem.
func (d Dir) Open(name string) (http.File, error) {
	var res http.File
	res, err := Get(filepath.Join(string(d), name))
	return res, err
}
