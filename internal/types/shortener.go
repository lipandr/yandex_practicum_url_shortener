package types

import "errors"

const UserIDSessionKey SessionKey = "userID"

var (
	ErrKeyExists  = errors.New("already shorten")
	ErrKeyDeleted = errors.New("shorten url was deleted")
)

type SessionKey string

type Session struct {
	UserID string
}

type ShortenRecord struct {
	UserID string
	Key    string
	Value  string
}
