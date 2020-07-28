package crawler

import (
	"path/filepath"

	"github.com/gronka/tg"
)

type Dir struct {
	Path      string
	ImageName string
	Folders   []string
	FolderIdx int
	Images    []string
	ImageIdx  int
}

// incrementFolderPath should be called with skipChildren false on first call
func incrementFolderPath(path string, skipChildren bool) (newPath string) {
	clean := filepath.Clean(path)
	children := getChildren(clean)

	if !skipChildren && len(children) > 0 {
		newPath = filepath.Join(clean, children[0])
	} else {
		// this else clause will return the next directory down, but returns
		// this directory if it is very last branch in the directory tree
		siblings, currentIdx := getSiblingsAndIdx(clean)
		nextIdx := currentIdx + 1
		parentPath, _ := cleanSplit(clean)
		if nextIdx >= len(siblings) {
			newPath = incrementFolderPath(parentPath, true)
		} else {
			newPath = filepath.Join(parentPath, siblings[nextIdx])
		}
	}

	return newPath
}

func decrementFolderPath(path string) (newPath string) {
	clean := filepath.Clean(path)

	siblings, currentIdx := getSiblingsAndIdx(clean)
	prevIdx := currentIdx - 1
	if prevIdx >= 0 {
		parentPath, _ := cleanSplit(clean)
		pickSibling := filepath.Join(parentPath, siblings[prevIdx])
		newPath = findLastChild(pickSibling)
	} else {
		parentPath, _ := cleanSplit(clean)
		newPath = decrementFolderPath(parentPath)
	}
	return newPath
}

func findLastChild(path string) (lastChild string) {
	tg.Info(path)
	children := getChildren(path)
	if len(children) > 0 {
		nextBranch := filepath.Join(path, children[len(children)-1])
		lastChild = findLastChild(nextBranch)
	} else {
		lastChild = path
	}
	return lastChild
}

func (craw *Crawler) incrementDir() {
	newPath := incrementFolderPath(craw.CurrentDir.Path, false)
	tg.Info("incrementDir: " + newPath)
	craw.JumpToPath(newPath)
}

func (craw *Crawler) decrementDir() {
	newPath := decrementFolderPath(craw.CurrentDir.Path)
	tg.Info("decrementDir: " + newPath)
	craw.JumpToPath(newPath)
}
