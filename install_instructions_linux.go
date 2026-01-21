//go:build linux

package main

func getInstallInstructions() string {
	return "You can install Universal Ctags with:\n" +
		"- Ubuntu/Debian: sudo apt-get install universal-ctags\n" +
		"- Fedora: sudo dnf install ctags\n" +
		"- Arch Linux: sudo pacman -S ctags"
}
