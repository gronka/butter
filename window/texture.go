package window

import (
	"image"
	"image/draw"

	"github.com/go-gl/gl/v2.1/gl"
)

type Texture struct {
	Handle uint32
}

func NewTexture() *Texture {
	var handle uint32
	gl.GenTextures(1, &handle)
	t := &Texture{handle}
	t.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	return t
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.Handle)
}

func (t *Texture) SetImage(im image.Image) {
	rgba := image.NewRGBA(im.Bounds())
	draw.Draw(rgba, rgba.Rect, im, image.ZP, draw.Src)

	size := rgba.Rect.Size()
	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA, int32(size.X), int32(size.Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
}
