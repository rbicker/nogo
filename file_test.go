package nogo

import (
	"io"
	"testing"
)

func TestLoadFile(t *testing.T) {
	f, err := LoadFile("file.go")
	if err != nil {
		t.Error(err)
	}
	if f.FileInfo.Name() != "file.go" {
		t.Errorf("LoadFile with file.go failed, expected fileinfo name %v, got %v", "file.go", f.FileInfo.Name())
	}
	if f.FileInfo.IsDir() {
		t.Errorf("LoadFile with file.go failed, expected fileinfo is dir %v, got %v", false, true)
	}
	d, err := LoadFile("internal")
	if err != nil {
		t.Error(err)
	}
	if !d.FileInfo.IsDir() {
		t.Errorf("LoadFile with file.go failed, expected fileinfo is dir %v, got %v", true, false)
	}
}

func TestRead(t *testing.T) {
	s := "123456789"
	read := ""
	f := File{Content: []byte(s)}
	b := make([]byte, 1)
	for {
		_, err := f.Read(b)
		if err == io.EOF {
			break
		}
		read = read + string(b)
		if err != nil {
			t.Fatal(err)
		}
	}
	if read != s {
		t.Errorf("expected %s, got %s", s, read)
	}
}

func TestReaddir(t *testing.T) {
	d, err := LoadFile("internal")
	if err != nil {
		t.Error(err)
	}
	infos, err := d.Readdir(0)
	if err != nil {
		t.Error(err)
	}
	name := infos[0].Name()
	if name != "nogogen" {
		t.Errorf("Readdir with internal failed, expected first fileinfo to be %v, got %v", "nogogen", name)

	}
}
