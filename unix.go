//go:build !windows

package main

import (
	"net/url"
	"path/filepath"
)

// pathToFileURI expects an absolute, cleaned filesystem path.
func pathToFileURI(path string) string {
	slashPath := filepath.ToSlash(path)
	return (&url.URL{Scheme: "file", Path: slashPath}).String()
}
