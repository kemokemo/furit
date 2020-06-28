package furit

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type markdown struct{}

const mdImgLinkFormat = `\!\[.*\]\((.+\.(png|PNG|jpg|JPG|jpeg|JPEG|bmp|BMP|gif|GIF|tiff|TIFF))\)`

var mdLinkReg = regexp.MustCompile(mdImgLinkFormat)

// Find converts the path of images referenced by text to paths relative to the program and returns it.
func (m *markdown) Find(root string) ([]string, error) {
	var links []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return fmt.Errorf("failed to walk for markdown: %v", err)
		}
		// Exclude directories with names beginning with a dot. ex) .git, .node_modules etc..
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		ext := filepath.Ext(info.Name())
		if info.IsDir() || (ext != ".md" && ext != ".markdown") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			group := mdLinkReg.FindSubmatch([]byte(s.Text()))
			if len(group) > 1 {
				links = append(links, filepath.Join(filepath.Dir(path), string(group[1])))
			}
		}
		return nil
	})

	return links, err
}
