package types

type Session struct {
	UserID string
}

type SessionKey string

const UserIDSessionKey SessionKey = "userID"

type ShortenRecord struct {
	UserID string
	Key    string
	Value  string
}
