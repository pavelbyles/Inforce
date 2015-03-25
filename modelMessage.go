package inforce

import (
	"time"
)

type Message struct {
	FromUser  UserInfo
	Text      string
	Timestamp time.Time
}
