//go:build windows

package main

func getInstallInstructions() string {
	return "You can install Universal Ctags with:\n" +
		"- Chocolatey: choco install universal-ctags\n" +
		"- Scoop: scoop install universal-ctags\n" +
		"Or download from: https://github.com/universal-ctags/ctags-win32/releases"
}
