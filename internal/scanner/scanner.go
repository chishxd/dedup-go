package scanner

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// WalkFiles walks the root dir recursively
// and returns a slice of all file paths.
// It skips hidden directories (starting with '.').
func WalkFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden dirs
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return fs.SkipDir
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
