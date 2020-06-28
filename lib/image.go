package furit

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type imageFinder struct{}

const imgFileFormat = `\.(png|PNG|jpg|JPG|jpeg|JPEG|bmp|BMP|gif|GIF|tiff|TIFF)`

var imgFileReg = regexp.MustCompile(imgFileFormat)

// Find finds image files and returns paths to the program.
func (i *imageFinder) Find(root string) ([]string, error) {
	var imgPaths []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return fmt.Errorf("failed to walk for image: %v", err)
		}
		// Exclude directories with names beginning with a dot. ex) .git, .node_modules etc..
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if info.IsDir() || !imgFileReg.MatchString(info.Name()) {
			return nil
		}
		imgPaths = append(imgPaths, path)
		return nil
	})

	return imgPaths, err
}
