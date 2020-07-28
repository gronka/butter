package window

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	//"github.com/gronka/tg"
)

// Canvas is the application's interface with what is drawn to the screen when
// we call Window.Draw()
type Canvas struct {
	RGBA    *image.RGBA
	Image   image.Image
	Texture *Texture
}

// LoadImage expects a cleaned path. Maybe we should test for thsi somehow
func (canvas *Canvas) LoadImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	im, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	canvas.Image = im

	return nil
}
