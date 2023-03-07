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

/*
func (m Message) Name() string {
	return m.name
}

func (m Message) Text() string {
	return m.text
}
*/
