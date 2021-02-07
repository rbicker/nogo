package nogo

import (
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
