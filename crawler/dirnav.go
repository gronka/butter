package crawler

import (
	"io/ioutil"
	"path/filepath"

	"butter/types"

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

// TODO: remember position image idx for each folder in the tree
func getChildren(path string) []string {
	//parentPath, _ := cleanSplit(path)
	folders, _ := getFoldersAndImages(path)
	return folders
}

func getSiblingsAndIdx(path string) ([]string, int) {
	parentPath, currentFolderName := cleanSplit(path)
	files, err := ioutil.ReadDir(parentPath)
	if err != nil {
		panic(err)
	}

	var siblings []string
	currentIdx := -1
	for _, file := range files {
		if file.IsDir() {
			siblings = append(siblings, file.Name())
			if currentFolderName == file.Name() {
				currentIdx = len(siblings) - 1
			}
		}
	}
	return siblings, currentIdx
}

func getFoldersAndImages(path string) ([]string, []string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var folders []string
	var images []string
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			folders = append(folders, fileName)

		} else {
			extension := filepath.Ext(fileName)
			if types.IsImage[extension] {
				images = append(images, fileName)
			}

		}
	}
	return folders, images
}

func incrementFolderPath(path string, skipChildren bool) (newPath string) {
	tg.Info("INC=============================")
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

	tg.Info("INC=============================")
	return newPath
}

func (craw *Crawler) incrementDir() {
	newPath := incrementFolderPath(craw.CurrentDir.Path, false)
	tg.Info(newPath)
	craw.JumpToPath(newPath)
}
