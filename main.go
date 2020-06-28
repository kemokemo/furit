package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	furit "github.com/kemokemo/furit/lib"
)

// exitCode
const (
	exitCodeOK = iota
	exitCodeInvalidArgs
	exitCodeInternalOperation
)

const (
	usage = `Usage: furit [<option>...] <1st path> <2nd path>...
 you can set mutiple paths to search invalid images.`
)

var (
	out     io.Writer = os.Stdout
	outerr  io.Writer = os.Stderr
	exit              = os.Exit
	cmdArgs []string
)

// flags
var (
	help bool
)

func init() {
	testing.Init() // require Go 1.13 or later
	flag.BoolVar(&help, "h", false, "display help")
	flag.Parse()
	cmdArgs = flag.Args()
}

func main() {
	exit(run(cmdArgs))
}

func run(args []string) int {
	if help {
		fmt.Fprintf(out, "%s\n\n", usage)
		flag.PrintDefaults()
		return exitCodeOK
	}
	if len(args) == 0 {
		fmt.Fprintf(outerr, "path is empty. please set it.\n\n%v", usage)
		return exitCodeInvalidArgs
	}

	exitCode := exitCodeOK
	for _, root := range args {
		_, err := os.Lstat(root)
		if err != nil {
			fmt.Fprintf(outerr, "path is invalid: %v\n", err)
			exitCode = exitCodeInvalidArgs
			continue
		}

		links, err := furit.Markdown.Find(root)
		if err != nil {
			fmt.Fprintf(outerr, "failed to find links: %v", err)
			exitCode = exitCodeInternalOperation
			continue
		}

		imgPaths, err := furit.Image.Find(root)
		if err != nil {
			fmt.Fprintf(outerr, "failed to find image paths: %v", err)
			exitCode = exitCodeInternalOperation
			continue
		}

		imgMap := make(map[string](bool), len(links))
		for _, link := range links {
			imgMap[link] = false
		}

		for _, imPath := range imgPaths {
			_, ok := imgMap[imPath]
			if !ok {
				fmt.Fprintln(out, imPath)
			}
		}
	}

	return exitCode
}
