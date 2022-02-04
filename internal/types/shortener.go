package types

type Session struct {
	UserID string
}

type SessionKey string

const UserIdSessionKey SessionKey = "userID"
