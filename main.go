package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	os.Exit(run())
}

func run() int {
	wd, _ := os.Getwd()
	gd, err := findGitDir(wd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(gd)

	return 0
}

func findGitDir(path string) (string, error) {
	for {
		gitDirPath := filepath.Join(path, ".git")
		_, err := os.Stat(gitDirPath)
		if err != nil {
			if path == "/" {
				return "", errors.New("fatal: not a git repository (or any of the parent directories): .git")
			}
			path = filepath.Dir(path)
			continue
		}
		return gitDirPath, nil
	}
}
