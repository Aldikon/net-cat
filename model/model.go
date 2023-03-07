package model

type Message struct {
	Name string
	Text string
}

func NewMessage(name string, text string) Message {
	return Message{
		Name: name,
		Text: text,
	}
}
