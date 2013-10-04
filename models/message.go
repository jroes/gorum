package models

type Message struct {
	Creator User
	Content string
	Topic   Topic
}
