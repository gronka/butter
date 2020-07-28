package crawler

import (
	"io/ioutil"
	"path/filepath"

	"butter/types"
	//"github.com/gronka/tg"
)

func (craw *Crawler) setDirFromImagePath(path string) {
	clean := filepath.Clean(path)
	craw.CurrentDir.Path = filepath.Dir(clean)
	craw.CurrentDir.ImageName = filepath.Base(clean)

	if path == craw.PrevDirPath {
		return
	}

	files, err := ioutil.ReadDir(craw.CurrentDir.Path)
	if err != nil {
		panic(err)
	}

	craw.CurrentDir.Images = nil
	craw.CurrentDir.Folders = nil
	craw.CurrentDir.ImageIdx = 0
	craw.CurrentDir.FolderIdx = 0
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			craw.CurrentDir.Folders = append(craw.CurrentDir.Folders, fileName)

		} else {
			extension := filepath.Ext(fileName)
			if types.IsImage[extension] {
				craw.CurrentDir.Images = append(craw.CurrentDir.Images, fileName)
			}

			if fileName == craw.CurrentDir.ImageName {
				craw.CurrentDir.ImageIdx = len(craw.CurrentDir.Images) - 1
			}
		}
	}
	craw.PrevDirPath = craw.CurrentDir.Path
}
