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
// TODO: maybe ignore folders with no images
func (craw *Crawler) ShiftFolderIdx(shift int, path string) {
	subFolders := getFolderInfo(path)

	if shift > 0 {
		if len(subFolders) > 0 {
			newPath := filepath.Join(path, subFolders[0])
			craw.CurrentDir = newDir(newPath)
			craw.setCurrentImage(craw.CurrentImagePath())
			return
		} else {

		}
	}

	/*
	* Below here, we need to step up the folder tree
	 */
	path = filepath.Clean(path)
	idx, folders := getFolderIdxInParent(path)

	idx = idx + shift
	if idx < len(folders) && idx >= 0 {
		parentPath, _ := filepath.Split(path)
		folderPath := filepath.Join(parentPath, folders[idx])
		craw.CurrentDir = newDir(folderPath)
		craw.setCurrentImage(craw.CurrentImagePath())
		return
	}

	if idx >= len(folders) {
		upOneLevel, _ := filepath.Split(path)
		tg.Debug("index too high")
		craw.ShiftFolderIdx(shift, upOneLevel)
		return
	}

	if idx < 0 {
		upOneLevel, _ := filepath.Split(path)
		tg.Debug("index too low")
		craw.ShiftFolderIdx(shift, upOneLevel)
		return
	}
}

func getFolderInfo(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var folders []string
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders

}

func getFolderIdxInParent(path string) (int, []string) {
	parentPath, folderName := filepath.Split(path)
	tg.Info(parentPath)
	tg.Info(parentPath)
	tg.Info(parentPath)
	files, err := ioutil.ReadDir(parentPath)
	if err != nil {
		panic(err)
	}

	var folders []string
	folderIdx := -1
	for idx, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
			if folderName == file.Name() {
				folderIdx = idx
				folderIdx = len(folders) - 1
			}
		}
	}
	return folderIdx, folders
}

func newDir(path string) Dir {
	var dir Dir
	dir.Path = path

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	dir.Images = nil
	dir.Folders = nil
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			dir.Folders = append(dir.Folders, fileName)

		} else {
			extension := filepath.Ext(fileName)
			if types.IsImage[extension] {
				dir.Images = append(dir.Images, fileName)
			}

			//if fileName == dir.ImageName {
			//dir.ImageIdx = len(dir.Images) - 1
			//}
		}
		//tg.Info(file.Name())
	}
	dir.ImageIdx = 0
	return dir
}

func (craw *Crawler) setCurrentImage(path string) {
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
