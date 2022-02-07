package types

import "errors"

type Session struct {
	UserID string
}

type SessionKey string

var ErrKeyExists = errors.New("already shorten")

const UserIDSessionKey SessionKey = "userID"

type ShortenRecord struct {
	UserID string
	Key    string
	Value  string
}
