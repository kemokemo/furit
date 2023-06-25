package furit

import (
	"path/filepath"
	"reflect"
	"testing"
)

func Test_html_Find(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "test-data", args: args{root: filepath.Join("test-data", "html")},
			want: []string{
				filepath.Join("test-data", "html", "assets", "sample.png"),
				filepath.Join("test-data", "html", "assets", "test.jpg"),
				filepath.Join("test-data", "html", "assets", "真夏の秋葉原.png"),
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &html{}
			got, err := m.Find(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("html.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("html.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
