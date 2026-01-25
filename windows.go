//go:build windows

package main

import (
	"net/url"
	"path/filepath"
)

// pathToFileURI expects path was run through `normalizePath()`.
func pathToFileURI(path string) string {
	slashPath := "/" + filepath.ToSlash(path) // Turns invalid "file://C:/" into valid "file:///C:/"
	return (&url.URL{Scheme: "file", Path: slashPath}).String()
}
