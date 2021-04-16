package main

import (
	"embed"
	"github.com/leaanthony/slicer"
	"github.com/matryer/is"
	"io/fs"
	"testing"
)

//go:embed fixtures
var testfixtures embed.FS

func Test1(t *testing.T) {
	is2 := is.New(t)
	d, err := Sub(testfixtures, "fixtures")
	is2.NoErr(err)
	files, err := d.ReadDir("test1")
	is2.NoErr(err)
	is2.Equal(len(files), 1)
	is2.Equal(files[0].Name(), "onefile.txt")
	file, err := d.ReadFile("test1/onefile.txt")
	is2.NoErr(err)
	is2.Equal(string(file), "test")
}

func Test2(t *testing.T) {
	is2 := is.New(t)
	d, err := Sub(testfixtures, "fixtures/test2")
	is2.NoErr(err)
	files, err := d.ReadDir("inner")
	is2.NoErr(err)
	is2.Equal(len(files), 3)

	expectedFiles := slicer.String([]string{
		"deeper",
		"one.txt",
		"two.txt",
	})
	for _, file := range files {
		is2.True(expectedFiles.Contains(file.Name()))
	}

	file, err := d.ReadFile("inner/deeper/three.txt")
	is2.NoErr(err)
	is2.Equal(string(file), "3")
}

func TestSub(t *testing.T) {
	is2 := is.New(t)
	inner, err := Sub(testfixtures, "fixtures/test2/inner")
	is2.NoErr(err)
	expectedFiles := slicer.String([]string{
		".",
		"deeper",
		"deeper/three.txt",
		"one.txt",
		"two.txt",
	})
	err = fs.WalkDir(inner, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		is2.True(expectedFiles.Contains(path))
		return nil
	})
	is2.NoErr(err)
	deeper, err := inner.Sub("deeper")
	is2.NoErr(err)
	files, err := deeper.ReadDir(".")
	is2.NoErr(err)
	is2.Equal(len(files), 1)
	is2.Equal(files[0].Name(), "three.txt")
	file, err := deeper.ReadFile("three.txt")
	is2.NoErr(err)
	is2.Equal(string(file), "3")
}

func TestBad(t *testing.T) {
	is2 := is.New(t)
	_, err := Sub(testfixtures, "baddir")
	is2.True(err != nil)

	root, err := Sub(testfixtures, "fixtures")
	is2.NoErr(err)

	_, err = root.Sub("baddir")
	is2.True(err != nil)

}
