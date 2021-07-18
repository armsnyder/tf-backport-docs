package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if err := run(os.DirFS(".")); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

func run(fs fs.FS) error {
	return nil
}
