package main

import (
	"os"
)

func main() {
	// TODO: use flags, maybe to set initial image
	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "/home/supper/Downloads/memes/cc/monster.jpg"
	}

	var engine Engine
	params := EngineParams{InitialImagePath: path}
	engine.init(params)
	engine.start()
}
