package cg

import (
	"fmt"
)

type Player struct {
	Name string "name"
	Level int "level"
	Exp int "exp"
	Room int "room"
	mq chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func (p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player)

	return player
}