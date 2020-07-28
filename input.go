package main

import (
	"os"

	"butter/qs"
	"butter/types"
	"butter/window"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/gronka/tg"
)

func handleInput(key glfw.Key, window *window.Window) {
	switch key {
	case glfw.KeyF:
		window.ActionQ <- qs.Listener{
			Command: types.NEXT_IMAGE,
		}
		break

	case glfw.KeyS:
		window.ActionQ <- qs.Listener{
			Command: types.PREV_IMAGE,
		}
		break

	case glfw.KeyE:
		window.ActionQ <- qs.Listener{
			Command: types.PREV_DIR,
		}
		break

	case glfw.KeyD:
		window.ActionQ <- qs.Listener{
			Command: types.NEXT_DIR,
		}
		break

	case glfw.KeyEscape:
		tg.Info("escaped")
		os.Exit(0)
		break

	default:
		tg.Error("Input: key not captured")
	}
}
