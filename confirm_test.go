package main

import (
	"bytes"
	"io"
	"testing"
)

func Test_askForConfirmation(t *testing.T) {
	type args struct {
		s     string
		in    io.Reader
		retry int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantOut string
		wantErr bool
	}{
		{"yes", args{"Are you sure to delete it?", bytes.NewBufferString("y\n"), 2}, true, "Are you sure to delete it? [y/n]: ", false},
		{"no", args{"Are you sure to delete it?", bytes.NewBufferString("n\n"), 2}, false, "Are you sure to delete it? [y/n]: ", false},
		{"no EOF", args{"Are you sure to delete it?", bytes.NewBufferString("hoge"), 2}, false, "Are you sure to delete it? [y/n]: ", true},
		{"invalid input", args{"Are you sure to delete it?", bytes.NewBufferString("foo\nbar\n"), 2}, false, "Are you sure to delete it? [y/n]: Are you sure to delete it? [y/n]: ", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := askForConfirmation(tt.args.s, tt.args.in, out, tt.args.retry)
			if (err != nil) != tt.wantErr {
				t.Errorf("askForConfirmation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("askForConfirmation() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("askForConfirmation() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
