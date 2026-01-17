package main

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func TestListWorkspaceFilesCommandOrder(t *testing.T) {
	t.Run("fd wins", func(t *testing.T) {
		workspaceRoot := t.TempDir()
		binDir := t.TempDir()
		writeScript(t, binDir, "fd", "echo \"fd_file.go\"")
		writeScript(t, binDir, "rg", "exit 1")
		writeScript(t, binDir, "git", "exit 1")
		t.Setenv("PATH", binDir)

		files, err := listWorkspaceFiles(workspaceRoot)
		if err != nil {
			t.Fatalf("list workspace files: %v", err)
		}
		if !reflect.DeepEqual(files, []string{"fd_file.go"}) {
			t.Fatalf("expected fd output, got %v", files)
		}
	})

	t.Run("rg fallback", func(t *testing.T) {
		workspaceRoot := t.TempDir()
		binDir := t.TempDir()
		writeScript(t, binDir, "fd", "exit 1")
		writeScript(t, binDir, "rg", "printf \"%s\\n\" \"rg_file.go\" \"rg_other.go\"")
		writeScript(t, binDir, "git", "exit 1")
		t.Setenv("PATH", binDir)

		files, err := listWorkspaceFiles(workspaceRoot)
		if err != nil {
			t.Fatalf("list workspace files: %v", err)
		}
		if !reflect.DeepEqual(files, []string{"rg_file.go", "rg_other.go"}) {
			t.Fatalf("expected rg output, got %v", files)
		}
	})

	t.Run("fd empty output returns error", func(t *testing.T) {
		workspaceRoot := t.TempDir()
		binDir := t.TempDir()
		writeScript(t, binDir, "fd", "exit 0")
		writeScript(t, binDir, "rg", "exit 1")
		writeScript(t, binDir, "git", "exit 1")
		t.Setenv("PATH", binDir)

		if _, err := listWorkspaceFiles(workspaceRoot); err == nil {
			t.Fatalf("expected error for empty workspace output")
		}
	})

	t.Run("walkdir fallback", func(t *testing.T) {
		workspaceRoot := t.TempDir()
		binDir := t.TempDir()
		writeScript(t, binDir, "fd", "exit 1")
		writeScript(t, binDir, "rg", "exit 1")
		writeScript(t, binDir, "git", "exit 1")
		t.Setenv("PATH", binDir)

		firstPath := filepath.Join(workspaceRoot, "first.go")
		if err := os.WriteFile(firstPath, []byte("package main\n"), 0o644); err != nil {
			t.Fatalf("write first file: %v", err)
		}
		subDir := filepath.Join(workspaceRoot, "subdir")
		if err := os.MkdirAll(subDir, 0o755); err != nil {
			t.Fatalf("mkdir subdir: %v", err)
		}
		secondPath := filepath.Join(subDir, "second.go")
		if err := os.WriteFile(secondPath, []byte("package main\n"), 0o644); err != nil {
			t.Fatalf("write second file: %v", err)
		}

		files, err := listWorkspaceFiles(workspaceRoot)
		if err != nil {
			t.Fatalf("list workspace files: %v", err)
		}

		sort.Strings(files)
		expected := []string{firstPath, secondPath}
		sort.Strings(expected)
		if !reflect.DeepEqual(files, expected) {
			t.Fatalf("expected walkdir output %v, got %v", expected, files)
		}
	})

	t.Run("walkdir empty workspace returns error", func(t *testing.T) {
		workspaceRoot := t.TempDir()
		binDir := t.TempDir()
		writeScript(t, binDir, "fd", "exit 1")
		writeScript(t, binDir, "rg", "exit 1")
		writeScript(t, binDir, "git", "exit 1")
		t.Setenv("PATH", binDir)

		if _, err := listWorkspaceFiles(workspaceRoot); err == nil {
			t.Fatalf("expected error for empty workspace")
		}
	})
}

func writeScript(t *testing.T, dir, name, body string) string {
	t.Helper()

	path := filepath.Join(dir, name)
	contents := "#!/bin/sh\n" + body + "\n"
	if err := os.WriteFile(path, []byte(contents), 0o755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	return path
}
