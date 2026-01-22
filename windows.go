//go:build windows

package main

import (
	"net/url"
	"path/filepath"
)

func getInstallInstructions() string {
	return "You can install Universal Ctags with:\n" +
		"- Chocolatey: choco install universal-ctags\n" +
		"- Scoop: scoop install universal-ctags\n" +
		"Or download from: https://github.com/universal-ctags/ctags-win32/releases"
}

// pathToFileURI expects an absolute, cleaned filesystem path.
func pathToFileURI(path string) string {
	slashPath := "/" + filepath.ToSlash(path) // Turns invalid "file://C:/" into valid "file:///C:/"
	return (&url.URL{Scheme: "file", Path: slashPath}).String()
}
