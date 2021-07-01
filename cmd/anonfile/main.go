package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wabarc/go-anonfile"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  anonfile [options] [file1] ... [fileN]\n")

		flag.PrintDefaults()
	}
	var basePrint = func() {
		fmt.Print("A CLI tool help upload files to anonfiles.\n\n")
		flag.Usage()
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		basePrint()
		os.Exit(0)
	}

}

func main() {
	files := flag.Args()

	anon := anonfile.NewAnonfile(nil)
	for _, path := range files {
		if _, err := os.Stat(path); err != nil {
			fmt.Fprintf(os.Stderr, "anonfile: %s: no such file or directory\n", path)
			continue
		}

		if url, err := anon.Upload(path); err != nil {
			fmt.Fprintf(os.Stderr, "anonfile: %v\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s  %s\n", url.Short(), path)
		}
	}
}
