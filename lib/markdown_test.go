package furit

import (
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
		{name: "test-data", args: args{root: "test-data/markdown"},
			want: []string{
				"test-data/markdown/posts/assets/gopher.png",
				"test-data/markdown/assets/sample1.png",
				"test-data/markdown/assets/サンプル.png",
				"test-data/markdown/logo.jpg",
				"test-data/markdown/テスト.png",
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
