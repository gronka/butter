package main

import (
	"github.com/gronka/butter/crawler"
	"github.com/gronka/butter/qs"
	"github.com/gronka/butter/window"

	"github.com/gronka/tg"
)

type Engine struct {
	Window           *window.Window
	Crawler          crawler.Crawler
	ActionQ          qs.ActionQ
	InitialImagePath string
	// TODO: config
	// TODO: keybinds structure
}

type EngineParams struct {
	InitialImagePath string
}

func (e *Engine) init(params EngineParams) {
	tg.Info("e.init()")
	e.InitialImagePath = params.InitialImagePath

	e.ActionQ = qs.ActionQ{
		Actions:         make(chan qs.Listener),
		CrawlerListener: make(chan qs.Listener),
		WindowListener:  make(chan qs.Listener),
	}

	e.Window = &window.Window{
		Canvas:      &window.Canvas{},
		HandleInput: handleInput,
		ActionQ:     e.ActionQ.Actions,
	}

	e.Crawler = crawler.Crawler{
		ActionQ:  e.ActionQ.Actions,
		Listener: e.ActionQ.CrawlerListener,
		Canvas:   e.Window.Canvas,
	}
}

func (e *Engine) start() {
	tg.Info("e.start()")
	go e.ActionQ.Start()
	go e.Crawler.Start()
	e.Crawler.JumpToImage(e.InitialImagePath)
	window.RunUntilDeath(e.Window)
}
