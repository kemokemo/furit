package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	furit "github.com/kemokemo/furit/lib"
	"gopkg.in/yaml.v3"
)

// exitCode
const (
	exitCodeOK = iota
	exitCodeFoundUnreferencedImages
	exitCodeInvalidArgs
	exitCodeInternalOperation
	exitCodeFailedToRemoveFiles
)

const (
	usage = `Usage: furit [<option>...] <1st path> <2nd path>...
 you can set multiple paths to search invalid images.`

	flags = `-d, -delete:
    delete unlinked image files (with confirmation)

-f, -force:
    delete unlinked image files without prompting for confirmation

-s, -settings:
    specify the settings file path to exclude files etc..

-t, -type:
    specify the target text format (markdown, html are available)

-h, -help:
    display help

-v, -version:
    display version`
)

var (
	out     io.Writer = os.Stdout
	outerr  io.Writer = os.Stderr
	exit              = os.Exit
	cmdArgs []string
)

// flags
var (
	help         bool
	ver          bool
	delFlag      bool
	forceFlag    bool
	typeFlag     string
	settingsPath string
)

func init() {
	testing.Init() // require Go 1.13 or later
	flag.BoolVar(&help, "help", false, "display help")
	flag.BoolVar(&help, "h", false, "display help")
	flag.BoolVar(&ver, "version", false, "display version")
	flag.BoolVar(&ver, "v", false, "display version")
	flag.BoolVar(&delFlag, "delete", false, "delete unlinked image files")
	flag.BoolVar(&delFlag, "d", false, "delete unlinked image files")
	flag.BoolVar(&forceFlag, "force", false, "delete unlinked image files without prompting for confirmation")
	flag.BoolVar(&forceFlag, "f", false, "delete unlinked image files without prompting for confirmation")
	flag.StringVar(&typeFlag, "type", "markdown", "file type to check links")
	flag.StringVar(&typeFlag, "t", "markdown", "file type to check links")
	flag.StringVar(&settingsPath, "settings", "", "settings file path")
	flag.StringVar(&settingsPath, "s", "", "settings file path")
	flag.Parse()
	cmdArgs = flag.Args()
}

func main() {
	exit(run(cmdArgs))
}

func run(args []string) int {
	if help {
		fmt.Fprintf(out, "%s\n\n%s\n", usage, flags)
		return exitCodeOK
	}
	if ver {
		fmt.Fprintf(out, "%s version %s.%s\n", Name, Version, revision)
		return exitCodeOK
	}
	if len(args) == 0 {
		fmt.Fprintf(outerr, "path is empty. please set it.\n\n%v\n", usage)
		return exitCodeInvalidArgs
	}
	linkFinder, err := getFinderByTypeFlag(typeFlag)
	if err != nil {
		fmt.Fprintf(outerr, "type error, %v\n", err)
		return exitCodeInvalidArgs
	}
	var settings settings
	if settingsPath != "" {
		b, err := os.ReadFile(settingsPath)
		if err != nil {
			fmt.Fprintf(outerr, "failed to read settings error, %v\n", err)
			return exitCodeInvalidArgs
		}
		err = yaml.Unmarshal(b, &settings)
		if err != nil {
			fmt.Fprintf(outerr, "failed to unmarshal settings error, %v\n", err)
			return exitCodeInvalidArgs
		}
	}

	exitCode := exitCodeOK
	for _, root := range args {
		_, err := os.Lstat(root)
		if err != nil {
			fmt.Fprintf(outerr, "path is invalid: %v\n", err)
			exitCode = exitCodeInvalidArgs
			continue
		}

		links, err := linkFinder.Find(root)
		if err != nil {
			fmt.Fprintf(outerr, "failed to find links: %v\n", err)
			exitCode = exitCodeInternalOperation
			continue
		}

		imgPaths, err := furit.Image.Find(root)
		if err != nil {
			fmt.Fprintf(outerr, "failed to find image paths: %v\n", err)
			exitCode = exitCodeInternalOperation
			continue
		}

		imgMap := make(map[string](bool), len(links))
		for _, link := range links {
			imgMap[link] = false
		}

		var delPaths []string
		for _, imPath := range imgPaths {
			_, ok := imgMap[imPath]
			if ok {
				continue
			}
			if isExcludePath(imPath, root, settings.Excludes) {
				continue
			}
			delPaths = append(delPaths, imPath)
			fmt.Fprintln(out, imPath)
		}

		if len(delPaths) > 0 {
			exitCode = exitCodeFoundUnreferencedImages
		}

		if !delFlag || len(delPaths) == 0 {
			continue
		}

		if !forceFlag {
			res, err := askForConfirmation("Are you sure to delete these unlinked images?", os.Stdin, out, 3)
			if !res {
				if err != nil {
					fmt.Fprintf(outerr, "the file deletion process has been canceled: %s\n", err)
					continue
				} else {
					fmt.Fprintln(outerr, "the file deletion process has been canceled by user input")
					continue
				}
			}
		}
		for _, delPath := range delPaths {
			e := os.Remove(delPath)
			if e != nil {
				fmt.Fprintf(outerr, "failed to remove file: %s\n", e)
				exitCode = exitCodeFailedToRemoveFiles
			}
		}
	}

	return exitCode
}

func getFinderByTypeFlag(tf string) (furit.ImageLinkFinder, error) {
	tf = strings.ToLower(tf)
	if tf == "markdown" {
		return furit.Markdown, nil
	} else if tf == "html" {
		return furit.HTML, nil
	} else {
		return nil, fmt.Errorf("unknown type flag: %v", tf)
	}
}

func isExcludePath(path string, root string, excludes []string) bool {
	for _, ex := range excludes {
		exPath := filepath.Join(root, ex)
		if exPath == path {
			return true
		}
	}
	return false
}
