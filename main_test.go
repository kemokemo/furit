package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	emptyArgs   = "path is empty. please set it.\n\nUsage: furit [<option>...] <1st path> <2nd path>...\n you can set multiple paths to search invalid images.\n"
	promptStr   = "Are you sure to delete these unlinked images? [y/n]: "
	canceledStr = "the file deletion process has been canceled: failed to read user input: EOF\n"
)

var (
	foundFilesArray = []string{
		filepath.Join("lib", "test-data", "markdown", "assets", "画面.bmp"),
		filepath.Join("lib", "test-data", "markdown", "テスト.gif"),
	}
	foundFiles         = strings.Join(foundFilesArray, "\n") + "\n"
	multiplePathsArray = []string{
		filepath.Join("lib", "test-data", "markdown", "assets", "画面.bmp"),
		filepath.Join("lib", "test-data", "markdown", "テスト.gif"),
		filepath.Join("lib", "test-data", "image-files", "blank.jpg"),
		filepath.Join("lib", "test-data", "image-files", "sample.png"),
		filepath.Join("lib", "test-data", "image-files", "画像", "テスト.gif"),
		filepath.Join("lib", "test-data", "image-files", "画面.bmp"),
	}
	multiplePaths = strings.Join(multiplePathsArray, "\n") + "\n"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		exitCode int
		out      string
		outerr   string
	}{
		{"not found", []string{filepath.Join("lib", "test-data", "markdown0")}, exitCodeOK, "", ""},
		{"found files", []string{filepath.Join("lib", "test-data", "markdown")}, exitCodeFoundUnreferencedImages, foundFiles, ""},
		{"empty args", []string{}, exitCodeInvalidArgs, "", emptyArgs},
		{"multiple paths", []string{filepath.Join("lib", "test-data", "markdown"), filepath.Join("lib", "test-data", "image-files")}, exitCodeFoundUnreferencedImages, multiplePaths, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			err := flagSet.Parse(tt.args)
			if err != nil {
				t.Errorf("failed to flagSet.Parse, %v", err)
			}
			cmdArgs = flagSet.Args()

			var exitCodeTest int
			exit = func(n int) {
				exitCodeTest = n
			}
			out = new(bytes.Buffer)
			outerr = new(bytes.Buffer)

			main()

			if exitCodeTest != tt.exitCode {
				t.Errorf("exit code = %v, want %v", exitCodeTest, tt.exitCode)
			}
			gotout := out.(*bytes.Buffer).String()
			if gotout != tt.out {
				t.Errorf("gotout = %v, want %v", gotout, tt.out)
				if diff := cmp.Diff(gotout, tt.out); diff != "" {
					fmt.Printf("diff:\n%s\n", diff)
				}
			}

			goterr := outerr.(*bytes.Buffer).String()
			if goterr != tt.outerr {
				t.Errorf("goterr = %v, want %v", goterr, tt.outerr)
				if diff := cmp.Diff(goterr, tt.outerr); diff != "" {
					fmt.Printf("diff:\n%s\n", diff)
				}
			}
		})
	}
}

func Test_main_help_ver(t *testing.T) {
	type args struct {
		help bool
		ver  bool
	}
	tests := []struct {
		name     string
		args     args
		exitCode int
		out      string
		outerr   string
	}{
		{"help", args{true, false}, exitCodeOK, fmt.Sprintf("%s\n\n%s\n", usage, flags), ""},
		{"version", args{false, true}, exitCodeOK, fmt.Sprintf("%s version %s.%s\n", Name, Version, Revision), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var exitCodeTest int
			exit = func(n int) {
				exitCodeTest = n
			}
			out = new(bytes.Buffer)
			outerr = new(bytes.Buffer)

			help = tt.args.help
			ver = tt.args.ver
			main()

			if exitCodeTest != tt.exitCode {
				t.Errorf("exit code = %v, want %v", exitCodeTest, tt.exitCode)
			}
			gotout := out.(*bytes.Buffer).String()
			if gotout != tt.out {
				t.Errorf("gotout = %v, want %v", gotout, tt.out)
			}
			goterr := outerr.(*bytes.Buffer).String()
			if goterr != tt.outerr {
				t.Errorf("goterr = %v, want %v", goterr, tt.outerr)
			}
		})
	}
	help = false
	ver = false
}

func Test_main_delete(t *testing.T) {
	help = false
	ver = false

	dir, err := os.MkdirTemp("test-data", "furit-test")
	if err != nil {
		t.Errorf("failed to create temp dir: %v", err)
	}
	defer func() {
		e := os.RemoveAll(dir)
		if err != nil {
			t.Errorf("failed to remove files: %v", e)
		}
	}()

	mdF, err := os.CreateTemp(dir, "test*.md")
	if err != nil {
		log.Printf("failed to create temporary file: %v", err)
	}
	defer func() {
		e := mdF.Close()
		if err != nil {
			log.Printf("failed to close temporary file: %v", e)
		}
	}()

	imgF, err := os.CreateTemp(dir, "sample*.png")
	if err != nil {
		log.Printf("failed to create temporary file: %v", err)
	}
	defer func() {
		e := imgF.Close()
		if err != nil {
			log.Printf("failed to close temporary file: %v", e)
		}
	}()

	type args struct {
		root  []string
		del   bool
		force bool
	}
	tests := []struct {
		name     string
		args     args
		exitCode int
		out      string
		outerr   string
	}{
		{"delete with prompt", args{[]string{dir}, true, false}, exitCodeFoundUnreferencedImages, fmt.Sprintf("%s\n%s", imgF.Name(), promptStr), canceledStr},
		// TODO: failed to run on Windows.
		//{"delete forcedly", args{[]string{dir}, true, true}, exitCodeFoundUnreferencedImages, fmt.Sprintf("%v\n", imgF.Name()), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			err := flagSet.Parse(tt.args.root)
			if err != nil {
				t.Errorf("failed to flagSet.Parse, %v", err)
			}
			cmdArgs = flagSet.Args()

			var exitCodeTest int
			exit = func(n int) {
				exitCodeTest = n
			}
			out = new(bytes.Buffer)
			outerr = new(bytes.Buffer)

			delFlag = tt.args.del
			forceFlag = tt.args.force
			main()

			if exitCodeTest != tt.exitCode {
				t.Errorf("exit code = %v, want %v", exitCodeTest, tt.exitCode)
			}
			gotout := out.(*bytes.Buffer).String()
			if gotout != tt.out {
				t.Errorf("gotout = %v, want %v", gotout, tt.out)
			}
			goterr := outerr.(*bytes.Buffer).String()
			if goterr != tt.outerr {
				t.Errorf("goterr = %v, want %v", goterr, tt.outerr)
			}
		})
	}
	delFlag = false
	forceFlag = false
}
