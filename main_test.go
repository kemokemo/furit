package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	foundFiles    = "lib/test-data/markdown/assets/画面.bmp\nlib/test-data/markdown/テスト.gif\n"
	emptyArgs     = "path is empty. please set it.\n\nUsage: furit [<option>...] <1st path> <2nd path>...\n you can set mutiple paths to search invalid images.\n"
	emptyString   = "path is invalid: lstat : no such file or directory\n"
	multiplePaths = "lib/test-data/markdown/assets/画面.bmp\nlib/test-data/markdown/テスト.gif\nlib/test-data/image-files/blank.jpg\nlib/test-data/image-files/sample.png\nlib/test-data/image-files/画像/テスト.gif\nlib/test-data/image-files/画面.bmp\n"
	promptStr     = "Are you sure to delete these unlinked images? [y/n]: "
	canceledStr   = "the file deletion process has been canceled: failed to read user input: EOF\n"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		exitCode int
		out      string
		outerr   string
	}{
		{"not found", []string{"lib/test-data/markdown0"}, exitCodeOK, "", ""},
		{"found files", []string{"lib/test-data/markdown"}, exitCodeFoundUnreferencedImages, foundFiles, ""},
		{"empty args", []string{}, exitCodeInvalidArgs, "", emptyArgs},
		{"empty string", []string{""}, exitCodeInvalidArgs, "", emptyString},
		{"multiple paths", []string{"lib/test-data/markdown", "lib/test-data/image-files"}, exitCodeFoundUnreferencedImages, multiplePaths, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			flagSet.Parse(tt.args)
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
			}
			goterr := outerr.(*bytes.Buffer).String()
			if goterr != tt.outerr {
				t.Errorf("goterr = %v, want %v", goterr, tt.outerr)
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
		{"version", args{false, true}, exitCodeOK, fmt.Sprintf("%s version %s.%s\n", Name, Version, revision), ""},
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

	dir, err := ioutil.TempDir("test-data", "furit-test")
	if err != nil {
		t.Errorf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)
	mdF, err := ioutil.TempFile(dir, "test*.md")
	defer mdF.Close()
	imgF, err := ioutil.TempFile(dir, "sample*.png")
	defer imgF.Close()

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
		{"delete forcely", args{[]string{dir}, true, true}, exitCodeFoundUnreferencedImages, fmt.Sprintf("%v\n", imgF.Name()), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			flagSet.Parse(tt.args.root)
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
