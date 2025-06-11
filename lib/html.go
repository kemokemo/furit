package furit

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type html struct{}

// Find converts the path of images referenced by text to paths relative to the program and returns it.
func (m *html) Find(root string) ([]string, error) {
	var links []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return fmt.Errorf("failed to walk for html: %v", err)
		}
		// Exclude directories with names beginning with a dot. ex) .git, .node_modules etc..
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		ext := strings.ToLower(filepath.Ext(info.Name()))
		if info.IsDir() || (ext != ".htm" && ext != ".html") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer shortClose(f)

		doc, err := goquery.NewDocumentFromReader(f)
		if err != nil {
			return fmt.Errorf("failed to read html file, %v", err)
		}

		var err2 error
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			srcValue, exists := s.Attr("src")
			if !exists {
				return
			}
			srcLink, e := url.PathUnescape(srcValue)
			if e != nil {
				srcLink = srcValue
			}
			srcLink = strings.ReplaceAll(srcLink, "\\", "/")
			srcLink = strings.Split(srcLink, "?")[0]
			srcFilePath := filepath.Clean(filepath.Join(filepath.Dir(path), srcLink))
			if !slices.Contains(links, srcFilePath) {
				links = append(links, srcFilePath)
			}
		})
		return err2
	})

	return links, err
}
