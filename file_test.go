package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func mkfile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func TestFiles(t *testing.T) {
	testdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("cannot create testdir: %s", err)
	}
	defer os.RemoveAll(testdir)

	expectedFiles := map[string]string{
		"a.go": "f",
		"tmp":  "d",
	}

	for f, typ := range expectedFiles {
		tmp := filepath.Join(testdir, f)
		if typ == "f" {
			err := mkfile(tmp)
			if err != nil {
				t.Fatalf("create error: %s", err)
			} else if typ == "d" {
				err := os.Mkdir(tmp, 0666)
				if err != nil {
					t.Fatalf("create error: %s", err)
				}
			}
		}
	}

	files, err := Files(testdir)
	if err != nil {
		t.Fatalf("cannot get files: %s", err)
	}

	fileName := files[0].Name()
	if fileName != "a.go" {
		t.Fatalf("cannot get files: %s", err)
	}
}

func TestFilesFail(t *testing.T) {
	if _, err := Files("xxx"); err == nil {
		t.Fatalf("failed test: err is nil")
	}
}
