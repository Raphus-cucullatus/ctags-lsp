//go:build !darwin && !linux && !windows

package main

func getInstallInstructions() string {
	return "Please visit https://github.com/universal-ctags/ctags for installation instructions"
}
