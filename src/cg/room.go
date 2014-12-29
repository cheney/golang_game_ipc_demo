package cg

type Room struct {
	Name string "name"
	Num int "num"
	MaxUser int "max" //
	mq chan *Message
}

func NewRoom() *Room {
	m := make(chan *Message, 1024)
	room := &Room{"", 0, 0, m}
	return room
}