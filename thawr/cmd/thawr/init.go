package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func initWorkspace() {
	root := ".thawr"

	if _, err := os.Stat(root); err == nil {
		fmt.Println("Workspace already initialized.")
		return
	}

	dirs := []string{
		root,
		filepath.Join(root, "organizations"),
		filepath.Join(root, "services"),
		filepath.Join(root, "logs"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println(err)
			return
		}
	}

	config := []byte("version: 1\n")

	err := os.WriteFile(filepath.Join(root, "config.yaml"), config, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Thawr workspace initialized successfully.")
}
