package qs

import (
	"fmt"

	"butter/types"
	//"github.com/gronka/tg"
)

type Listener struct {
	Command int
	Message string
}

type ActionQ struct {
	Actions         chan Listener
	CrawlerListener chan Listener
	WindowListener  chan Listener
}

func (aq *ActionQ) Start() {
	var action Listener
	for {

		action = <-aq.Actions
		switch action.Command {

		case types.NEXT_IMAGE:
			aq.CrawlerListener <- Listener{
				Command: types.NEXT_IMAGE,
			}
			break

		case types.PREV_IMAGE:
			aq.CrawlerListener <- Listener{
				Command: types.PREV_IMAGE,
			}
			break

		case types.PREV_DIR:
			aq.CrawlerListener <- Listener{
				Command: types.PREV_DIR,
			}
			break

		case types.NEXT_DIR:
			aq.CrawlerListener <- Listener{
				Command: types.NEXT_DIR,
			}
			break

		default:
			fmt.Println("ActionQ: ActionQ case not found")
		}
	}
}
