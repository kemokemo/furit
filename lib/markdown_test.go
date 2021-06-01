package furit

import (
	"path/filepath"
	"reflect"
	"testing"
)

func Test_markdown_Find(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "test-data", args: args{root: filepath.Join("test-data", "markdown")},
			want: []string{
				filepath.Join("test-data", "markdown", "posts", "assets", "gopher.png"),
				filepath.Join("test-data", "markdown", "assets", "sample1.png"),
				filepath.Join("test-data", "markdown", "assets", "サンプル.png"),
				filepath.Join("test-data", "markdown", "logo.jpg"),
				filepath.Join("test-data", "markdown", "テスト.png"),
				filepath.Join("test-data", "markdown", "assets", "refereed_with_query.png"),
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &markdown{}
			got, err := m.Find(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("markdown.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("markdown.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
