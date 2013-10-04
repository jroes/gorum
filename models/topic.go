package models

type Topic struct {
	Title string
	Creator User
        Messages []Message
}
