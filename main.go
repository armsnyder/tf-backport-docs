package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	if err := run(dir); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(rootDir string) error {
	providerName, specs, err := parse(rootDir)
	if err != nil {
		return err
	}

	return render(rootDir, providerName, specs)
}
