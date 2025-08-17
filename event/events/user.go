package events

import (
	"fmt"
)

// UserCreated 用户创建事件
type UserCreated struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

func (*UserCreated) EventName() string {
	return "UserCreated"
}

func (e *UserCreated) EntityID() string {
	return fmt.Sprintf("%d", e.UserId)
}

func (e *UserCreated) EntityType() string {
	return "User"
}

// UserLoggedIn 用户登录事件
type UserLoggedIn struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Platform string `json:"platform"`
	LoginAt  string `json:"login_at"`
}

func (*UserLoggedIn) EventName() string {
	return "UserLoggedIn"
}

func (e *UserLoggedIn) EntityID() string {
	return fmt.Sprintf("%d", e.UserId)
}

func (e *UserLoggedIn) EntityType() string {
	return "User"
}
