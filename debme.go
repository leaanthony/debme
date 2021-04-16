package main

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
)

// Debme is an embed.FS compatible wrapper, providing Sub() functionality
type Debme struct {
	basedir string
	embed.FS
}

// Sub creates an embed.FS compatible struct, anchored to the given basedir.
func Sub(fs embed.FS, basedir string) (Debme, error) {
	result := Debme{
		FS:      fs,
		basedir: basedir,
	}
	_, err := result.ReadDir(".")
	if err != nil {
		return Debme{}, fmt.Errorf("cannot create Sub: invalid basedir '%s'", basedir)
	}
	return result, nil
}

// Open opens the named file for reading and returns it as an fs.File.
func (d Debme) Open(name string) (fs.File, error) {
	return d.FS.Open(filepath.Join(d.basedir, name))
}

// ReadDir reads and returns the entire named directory.
func (d Debme) ReadDir(name string) ([]fs.DirEntry, error) {
	return d.FS.ReadDir(filepath.Join(d.basedir, name))
}

// ReadFile reads and returns the content of the named file.
func (d Debme) ReadFile(name string) ([]byte, error) {
	return d.FS.ReadFile(filepath.Join(d.basedir, name))
}

// Sub returns a new Debme anchored at the given subdirectory.
func (d Debme) Sub(subDir string) (Debme, error) {
	_, err := d.ReadDir(subDir)
	if err != nil {
		return Debme{}, err
	}
	return Debme{
		basedir: filepath.Join(d.basedir, subDir),
		FS:      d.FS,
	}, nil
}
