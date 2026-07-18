package main

import (
	"fmt"
	"os"
)

const Version = "0.0.1-dev"

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {
	case "init":
		initWorkspace()
	case "version":
		fmt.Println("Thawr", Version)
	default:
		help()
	}
}

func help() {
	fmt.Println("Thawr")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  thawr init")
	fmt.Println("  thawr version")
}
