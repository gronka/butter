package crawler

import (
	"path/filepath"

	"github.com/gronka/butter/qs"
	"github.com/gronka/butter/types"
	"github.com/gronka/butter/window"
	"github.com/gronka/tg"
)

type Crawler struct {
	DirTree     map[string]Dir
	PrevDirPath string
	CurrentDir  Dir
	ParentDir   Dir
	ActionQ     chan qs.Listener
	Listener    chan qs.Listener
	Canvas      *window.Canvas
}

func (craw *Crawler) CurrentImagePath() string {
	path := ""
	if len(craw.CurrentDir.Images) == 0 {
		path = filepath.Join(
			craw.CurrentDir.Path,
			"",
		)
	} else {
		path = filepath.Join(
			craw.CurrentDir.Path,
			craw.CurrentDir.Images[craw.CurrentDir.ImageIdx],
		)
	}
	return path
}

func (craw *Crawler) Start() {
	for {
		action := <-craw.Listener
		switch action.Command {
		case types.NEXT_IMAGE:
			craw.ShiftImageIdx(1)
			break

		case types.PREV_IMAGE:
			craw.ShiftImageIdx(-1)
			break

		case types.PREV_DIR:
			craw.decrementDir()
			break

		case types.NEXT_DIR:
			craw.incrementDir()
			break

		default:
			tg.Info("Crawler: ActionQ case not found")
		}
	}
}

func (craw *Crawler) JumpToPath(path string) {
	craw.CurrentDir.Path = path
	craw.CurrentDir.Folders, craw.CurrentDir.Images = getFoldersAndImages(path)
	craw.CurrentDir.ImageName = ""
	craw.CurrentDir.ImageIdx = 0
	craw.LoadImage()
}

func (craw *Crawler) JumpToImage(path string) {
	craw.setDirFromImagePath(path)
	craw.LoadImage()
}

func (craw *Crawler) ShiftImageIdx(shift int) {
	if len(craw.CurrentDir.Images) == 0 {
		return
	}

	idx := craw.CurrentDir.ImageIdx + shift
	if idx >= len(craw.CurrentDir.Images) {
		idx = 0
	} else if idx < 0 {
		idx = len(craw.CurrentDir.Images) - 1
	}
	craw.CurrentDir.ImageIdx = idx
	craw.LoadImage()
}

func (craw *Crawler) LoadImage() {
	path := ""
	if len(craw.CurrentDir.Images) == 0 {
		path = filepath.Clean("./assets/not_found.jpg")
	} else {
		path = filepath.Join(
			craw.CurrentDir.Path,
			craw.CurrentDir.Images[craw.CurrentDir.ImageIdx],
		)
	}
	err := craw.Canvas.LoadImage(path)
	tg.Check(err, "failed to load image")
}
