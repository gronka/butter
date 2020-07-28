package crawler

import (
	"path/filepath"

	"github.com/gronka/tg"
)

func decrementFolderPath(path string, toBottom bool) (newPath string) {
	tg.Info("DEC=============================")
	clean := filepath.Clean(path)

	tg.Info(toBottom)

	if !toBottom {
		siblings, currentIdx := getSiblingsAndIdx(clean)
		prevIdx := currentIdx - 1
		tg.Info(prevIdx)
		if prevIdx >= 0 {
			tg.Info("prevIdx > 0")
			parentPath, _ := cleanSplit(clean)
			newPath = filepath.Join(parentPath, siblings[prevIdx])
		} else {
			tg.Info("prevIdx <= 0")
			parentPath, _ := cleanSplit(clean)
			newPath = decrementFolderPath(parentPath, true)
		}
	} else {
		// this else clause will return the very last branch of the directory
		// tree of the previous sibling folder
		children := getChildren(clean)
		tg.Info("children of toBottom: ")
		tg.Info(children)
		tg.Info(path)
		tg.Info("TOBOTTOM=============================")
		if len(children) == 0 {
			newPath = clean
		} else {
			lastChildPath := filepath.Join(clean, children[len(children)-1])
			newPath = decrementFolderPath(lastChildPath, true)
		}
	}

	tg.Info("DEC=============================")
	return newPath
}

func (craw *Crawler) decrementDir() {
	newPath := decrementFolderPath(craw.CurrentDir.Path, false)
	tg.Info(newPath)
	craw.JumpToPath(newPath)
}
