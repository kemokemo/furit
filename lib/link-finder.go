package furit

// ImageLinkFinder will search files and find image links.
type ImageLinkFinder interface {
	// Find converts the path of images referenced by text to paths relative to the program and returns it.
	Find(root string) ([]string, error)
}

// Markdown is the finder of image link from markdown files.
var Markdown ImageLinkFinder = (*markdown)(nil)

// HTML is the finder of image link from html files.
var HTML ImageLinkFinder = (*html)(nil)
