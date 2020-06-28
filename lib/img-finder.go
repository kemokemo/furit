package furit

// ImageFinder finds image files
type ImageFinder interface {
	// Find finds image files and returns paths to the program.
	Find(root string) ([]string, error)
}

// Image is the finder of image files.
var Image ImageFinder = (*imageFinder)(nil)
