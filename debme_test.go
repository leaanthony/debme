package main

import (
	"embed"
	"github.com/leaanthony/slicer"
	"github.com/matryer/is"
	"io"
	"io/fs"
	"testing"
)

//go:embed fixtures
var testfixtures embed.FS

func TestReadFile(t *testing.T) {
	is2 := is.New(t)
	d, err := Sub(testfixtures, "fixtures")
	is2.NoErr(err)
	file, err := d.ReadFile("test1/onefile.txt")
	is2.NoErr(err)
	is2.Equal(string(file), "test")
	test2, err := d.Sub("test2")
	is2.NoErr(err)
	file, err = test2.ReadFile("inner/deeper/three.txt")
	is2.NoErr(err)
	is2.Equal(string(file), "3")

	_, err = Sub(testfixtures, "badfixture")
	is2.True(err != nil)
	println(err.Error())

	_, err = d.ReadFile("badfile")
	is2.True(err != nil)
	println(err.Error())
}

func TestReadDir(t *testing.T) {
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

	_, err = d.ReadDir("baddir")
	is2.True(err != nil)
	println(err.Error())
}

func TestOpen(t *testing.T) {
	is2 := is.New(t)
	d, err := Sub(testfixtures, "fixtures/test1")
	is2.NoErr(err)
	file, err := d.Open("onefile.txt")
	is2.NoErr(err)
	data, err := io.ReadAll(file)
	is2.NoErr(err)
	is2.Equal(string(data), "test")

	_, err = d.Open("badfile")
	is2.True(err != nil)
	println(err.Error())

}

func TestCompatibility(t *testing.T) {
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
}

func TestSub(t *testing.T) {
	is2 := is.New(t)
	_, err := Sub(testfixtures, "baddir")
	is2.True(err != nil)
	println(err.Error())
}
