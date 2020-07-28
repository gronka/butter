package crawler

import (
	"path/filepath"
)

func cleanSplit(path string) (string, string) {
	clean := filepath.Clean(path)
	parent, folderName := filepath.Split(clean)
	return filepath.Clean(parent), folderName
}
