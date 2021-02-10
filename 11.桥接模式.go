package main



type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImlementer interface {
	Send(text, to string)
}
