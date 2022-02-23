package types

import "errors"

type Session struct {
	UserID string
}

type SessionKey string

var ErrKeyExists = errors.New("already shorten")
var ErrKeyDeleted = errors.New("shorten url was deleted")

const UserIDSessionKey SessionKey = "userID"

type ShortenRecord struct {
	UserID string
	Key    string
	Value  string
}
