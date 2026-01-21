//go:build windows

package main

import (
	"net/url"
	"path/filepath"
)

func encodePathForTest(path string) string {
	slashPath := "/" + filepath.ToSlash(path)
	return (&url.URL{Scheme: "file", Path: slashPath}).String()[len("file://"):]
}
