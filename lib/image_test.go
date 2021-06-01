package furit

import (
	"path/filepath"
	"reflect"
	"testing"
)

func Test_imageFinder_Find(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "test-data", args: args{root: filepath.Join("test-data", "image-files")},
			want: []string{
				filepath.Join("test-data", "image-files", "blank.jpg"),
				filepath.Join("test-data", "image-files", "sample.png"),
				filepath.Join("test-data", "image-files", "画像/テスト.gif"),
				filepath.Join("test-data", "image-files", "画面.bmp"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &imageFinder{}
			got, err := i.Find(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageFinder.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageFinder.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
