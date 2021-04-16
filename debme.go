package main

import (
	"embed"
	"io/fs"
	"path/filepath"
)

type embedFS = embed.FS

// Debme is an embed.FS compatible wrapper, providing Sub() functionality
type Debme struct {
	basedir string
	embedFS
}

// FS creates an embed.FS compatible struct, anchored to the given basedir.
func FS(fs embed.FS, basedir string) (Debme, error) {
	result := Debme{embedFS: fs, basedir: basedir}
	_, err := result.ReadDir(".")
	if err != nil {
		return Debme{}, err
	}
	return result, nil
}

// Open opens the named file for reading and returns it as an fs.File.
func (d Debme) Open(name string) (fs.File, error) {
	return d.embedFS.Open(filepath.Join(d.basedir, name))
}

// ReadDir reads and returns the entire named directory.
func (d Debme) ReadDir(name string) ([]fs.DirEntry, error) {
	return d.embedFS.ReadDir(filepath.Join(d.basedir, name))
}

// ReadFile reads and returns the content of the named file.
func (d Debme) ReadFile(name string) ([]byte, error) {
	return d.embedFS.ReadFile(filepath.Join(d.basedir, name))
}

// FS returns a new Debme anchored at the given subdirectory.
func (d Debme) FS(subDir string) (Debme, error) {
	return FS(d.embedFS, filepath.Join(d.basedir, subDir))
}
