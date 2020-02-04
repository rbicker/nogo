package nogo

import (
	"bytes"
	"fmt"
	"net/http"

	"encoding/gob"
)

// map for storing the nogos files as slice of bytes
var nogos map[string][]byte

// Dir corresponds to http.FileSystem
type Dir string

// ensure Dir corresponds to http.FileSystem
var _ http.FileSystem = Dir("")

func init() {
	nogos = make(map[string][]byte)
}

// Add adds the given bytes to nogo under the given name.
func Add(name string, b []byte) {
	nogos[name] = b
}

// Open decodes the file with the given name.
func Open(name string) (File, error) {
	f := File{}
	if b, ok := nogos[name]; ok {
		r := bytes.NewReader(b)
		dec := gob.NewDecoder(r)
		if err := dec.Decode(&f); err != nil {
			return f, err
		}
		return f, nil
	}
	return f, fmt.Errorf("file named '%v' not found", name)
}

// Open returns a http.File based on the given name.
// The function corresponds to http.FileSystem.
func (Dir) Open(name string) (http.File, error) {
	var res http.File
	res, err := Open(name)
	return res, err
}
