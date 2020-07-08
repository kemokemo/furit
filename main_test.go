package main

import (
	"bytes"
	"flag"
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		out    string
		outerr string
	}{
		{"not found", args{[]string{"lib/test-data/markdown0"}}, exitCodeOK, "", ""},
		{"found files", args{[]string{"lib/test-data/markdown"}}, exitCodeFoundUnreferencedImages, "lib/test-data/markdown/assets/画面.bmp\nlib/test-data/markdown/テスト.gif\n", ""},
		{"empty args", args{[]string{}}, exitCodeInvalidArgs, "", "path is empty. please set it.\n\nUsage: furit [<option>...] <1st path> <2nd path>...\n you can set mutiple paths to search invalid images.\n"},
		{"empty string", args{[]string{""}}, exitCodeInvalidArgs, "", "path is invalid: lstat : no such file or directory\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bout := new(bytes.Buffer)
			out = bout
			bouterr := new(bytes.Buffer)
			outerr = bouterr

			if got := run(tt.args.args); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}

			stdout := bout.String()
			if stdout != tt.out {
				t.Errorf("stdout = %v, want %v", stdout, tt.out)
			}

			stderr := bouterr.String()
			if stderr != tt.outerr {
				t.Errorf("stderr = %v, want %v", stderr, tt.outerr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		exitCode int
		out      string
		outerr   string
	}{
		{"not found", []string{"lib/test-data/markdown0"}, exitCodeOK, "", ""},
		{"found files", []string{"lib/test-data/markdown"}, exitCodeFoundUnreferencedImages, "lib/test-data/markdown/assets/画面.bmp\nlib/test-data/markdown/テスト.gif\n", ""},
		{"empty args", []string{}, exitCodeInvalidArgs, "", "path is empty. please set it.\n\nUsage: furit [<option>...] <1st path> <2nd path>...\n you can set mutiple paths to search invalid images.\n"},
		{"empty string", []string{""}, exitCodeInvalidArgs, "", "path is invalid: lstat : no such file or directory\n"},
		{"multiple paths", []string{"lib/test-data/markdown", "lib/test-data/image-files"}, exitCodeFoundUnreferencedImages, "lib/test-data/markdown/assets/画面.bmp\nlib/test-data/markdown/テスト.gif\nlib/test-data/image-files/blank.jpg\nlib/test-data/image-files/sample.png\nlib/test-data/image-files/画像/テスト.gif\nlib/test-data/image-files/画面.bmp\n", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			flagSet.Parse(tt.args)
			cmdArgs = flagSet.Args()

			bout := new(bytes.Buffer)
			out = bout
			bouterr := new(bytes.Buffer)
			outerr = bouterr

			var exitCode int
			exit = func(n int) {
				exitCode = n
			}

			main()

			if exitCode != tt.exitCode {
				t.Errorf("exit code = %v, want %v", exitCode, tt.exitCode)
			}

			stdout := bout.String()
			if stdout != tt.out {
				t.Errorf("stdout = %v, want %v", stdout, tt.out)
			}

			stderr := bouterr.String()
			if stderr != tt.outerr {
				t.Errorf("stderr = %v, want %v", stderr, tt.outerr)
			}
		})
	}
}

func Test_main_flags(t *testing.T) {
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
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			flagSet.Parse([]string{""})
			cmdArgs = flagSet.Args()

			bout := new(bytes.Buffer)
			out = bout
			bouterr := new(bytes.Buffer)
			outerr = bouterr

			var exitCode int
			exit = func(n int) {
				exitCode = n
			}

			help = tt.args.help
			ver = tt.args.ver
			main()

			if exitCode != tt.exitCode {
				t.Errorf("exit code = %v, want %v", exitCode, tt.exitCode)
			}

			stdout := bout.String()
			if stdout != tt.out {
				t.Errorf("stdout = %v, want %v", stdout, tt.out)
			}

			stderr := bouterr.String()
			if stderr != tt.outerr {
				t.Errorf("stderr = %v, want %v", stderr, tt.outerr)
			}
		})
	}
}
