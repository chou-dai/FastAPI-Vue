package model

type SessionInfo struct {
	UserId    int    `json:"userId"`
	Name      string `json:"name"`
	SessionId string `json:"sessionId"`
}
