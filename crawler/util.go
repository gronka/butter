package crawler

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gronka/butter/types"
)

func cleanSplit(path string) (string, string) {
	clean := filepath.Clean(path)
	parent, folderName := filepath.Split(clean)
	return filepath.Clean(parent), folderName
}

func getChildren(path string) []string {
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
