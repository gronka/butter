package main

import (
	"github.com/gronka/tg"
	"os"
)

func main() {
	// TODO: use flags, maybe to set initial image
	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	conf := GenerateConfig(path)

	tg.Info(conf)

	var engine Engine
	params := EngineParams{InitialImagePath: conf.InitialPath}
	engine.init(params)
	engine.start()
}
