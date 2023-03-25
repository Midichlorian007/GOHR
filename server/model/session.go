package model

import "time"

const (
	SessionKey = "session"
	SessionExpire = time.Hour * 24 // 1 day (in days)
	SessionMaxAge = 24 * 60 * 60   // 1 day (in seconds)
)

type Session struct {
	ID     string    `json:"id"`
	User   string    `json:"user"`
	Expire time.Time `json:"expire"`
}
