package window

import (
	"fmt"
	"runtime"

	"butter/qs"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/gronka/tg"
)

type Window struct {
	GlWindow    *glfw.Window
	Canvas      *Canvas
	HandleInput func(glfw.Key, *Window)
	ActionQ     chan qs.Listener
	PrevKey     glfw.Key
	KeyTime     int64
}

func init() {
	runtime.LockOSThread()
}

func RunUntilDeath(window *Window) {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	err := StartGlWindow(window)
	if err != nil {
		panic(err)
	}

	fmt.Println("about to hit while")

	for {
		if window.GlWindow.ShouldClose() {
			window.GlWindow.Destroy()
			break
		}

		// TODO: this draw seems to be linked to fps
		window.DrawImage()
		glfw.PollEvents()

	}
}

func StartGlWindow(window *Window) error {
	w := 1920
	h := 1080

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glWindow, err := glfw.CreateWindow(w, h, "buttergallery", nil, nil)
	if err != nil {
		return err
	}

	window.GlWindow = glWindow
	window.GlWindow.MakeContextCurrent()
	glfw.SwapInterval(1)
	//window.GlWindow.SetRefreshCallback(window.onRefresh)

	setupInput(window)
	window.Canvas.Texture = NewTexture()

	return nil
}

func setupInput(window *Window) {
	window.GlWindow.SetKeyCallback(
		func(_ *glfw.Window,
			key glfw.Key,
			scancode int,
			action glfw.Action,
			mods glfw.ModifierKey,
		) {
			if !window.IgnoreInput(key) {
				window.HandleInput(key, window)
			}
		})
}

func (window *Window) onRefresh(glfwWindow *glfw.Window) {
	window.DrawImage()
}

func (window *Window) DrawImage() {
	// TODO: smart rendering - don't rerender if no changes?

	window.Canvas.Texture.SetImage(window.Canvas.Image)

	iw := float32(window.Canvas.Image.Bounds().Size().X)
	ih := float32(window.Canvas.Image.Bounds().Size().Y)
	wRead, hRead := window.GlWindow.GetFramebufferSize()
	w := float32(wRead)
	h := float32(hRead)

	imageAspectRatio := iw / ih
	windowAspectRatio := w / h // will give you 4:3 or 16:9 type value

	var x, y float32
	if imageAspectRatio > windowAspectRatio {
		x = 1
		y = 1 / windowAspectRatio
	} else {
		x = 1 / windowAspectRatio
		y = 1
	}

	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.TEXTURE_2D)
	window.Canvas.Texture.Bind()
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 1)
	gl.Vertex2f(-x, -y)
	gl.TexCoord2f(1, 1)
	gl.Vertex2f(x, -y)
	gl.TexCoord2f(1, 0)
	gl.Vertex2f(x, y)
	gl.TexCoord2f(0, 0)
	gl.Vertex2f(-x, y)
	gl.End()
	window.GlWindow.SwapBuffers()
}

func (window *Window) IgnoreInput(key glfw.Key) bool {
	if window.PrevKey != key {
		window.PrevKey = key
		window.KeyTime = tg.TimeNowMilli()
		return false
	}

	if window.GlWindow.GetKey(key) == glfw.Release {
		window.PrevKey = glfw.KeyUnknown
	}

	if (tg.TimeNowMilli() - window.KeyTime) > 150 {
		window.KeyTime = tg.TimeNowMilli()
		return false
	}

	return true
}
