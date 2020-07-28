package crawler

import (
	"path/filepath"

	"butter/qs"
	"butter/types"
	"butter/window"

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
	tg.Info(craw.CurrentDir.Path)
	if len(craw.CurrentDir.Images) == 0 {
		path = filepath.Clean("./assets/not_found.jpg")
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
			craw.ShiftFolderIdx(-1, craw.CurrentDir.Path)
			break

		case types.NEXT_DIR:
			craw.ShiftFolderIdx(1, craw.CurrentDir.Path)
			break

		default:
			tg.Info("Crawler: ActionQ case not found")
		}
	}
}

func (craw *Crawler) JumpToImage(path string) {
	craw.setCurrentImage(path)
	err := craw.Canvas.LoadImage(craw.CurrentImagePath())
	tg.Check(err, "failed to load image")
}

func (craw *Crawler) ShiftImageIdx(shift int) {
	idx := craw.CurrentDir.ImageIdx + shift
	if idx >= len(craw.CurrentDir.Images) {
		idx = 0
	} else if idx < 0 {
		idx = len(craw.CurrentDir.Images) - 1
	}
	craw.CurrentDir.ImageIdx = idx
	err := craw.Canvas.LoadImage(craw.CurrentImagePath())
	tg.Check(err, "failed to load image")
}
