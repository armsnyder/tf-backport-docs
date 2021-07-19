package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	workDir := testMakeTempDir(t)
	testCopyDir(t, filepath.Join("testdata", "in"), workDir)
	if err := run(workDir); err != nil {
		t.Fatal(err)
	}
	testAssertDirsEqual(t, filepath.Join("testdata", "out"), workDir)
}

func testMakeTempDir(t *testing.T) string {
	t.Helper()
	dir, err := ioutil.TempDir("", "tf-backport-docs")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	})
	return dir
}

func testCopyDir(t *testing.T, from, to string) {
	t.Helper()
	err := filepath.WalkDir(from, func(path string, d fs.DirEntry, err error) error {
		dest := filepath.Join(to, strings.TrimPrefix(path, from))
		if d.IsDir() {
			return os.MkdirAll(dest, 0755)
		}
		reader, err := os.Open(path)
		if err != nil {
			return err
		}
		defer reader.Close()
		writer, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer writer.Close()
		_, err = io.Copy(writer, reader)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
}

func testAssertDirsEqual(t *testing.T, want, got string) {
	t.Helper()
	err := filepath.WalkDir(want, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		shortPath := strings.TrimPrefix(path, want)
		t.Run(shortPath, func(t *testing.T) {
			wantText, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			gotText, err := os.ReadFile(filepath.Join(got, shortPath))
			if errors.Is(err, os.ErrNotExist) {
				t.Fatalf("file %s does not exist", shortPath)
			}
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, string(wantText), string(gotText))
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
