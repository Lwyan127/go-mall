package listeners

import (
	"fmt"

	"github.com/go-study-lab/go-mall/common/errcode"
	"github.com/go-study-lab/go-mall/common/logger"
	"github.com/go-study-lab/go-mall/event/events"
	"github.com/kevinyan815/event"
)

// UserCreatedListener 用户创建事件监听器
type UserCreatedListener struct{}

func (*UserCreatedListener) EventHandler() event.Handler {
	return func(e *event.Event) error {
		userData, ok := e.EventData().(*events.UserCreated)
		if !ok {
			return errcode.ErrCoverData.WithCause(fmt.Errorf("cannot convert event data to UserCreated"))
		}

		logger.New(e.Context()).Info("User created event processed",
			"user_id", userData.UserId,
			"user_name", userData.UserName,
			"event_id", e.ID,
			"event_name", e.Name,
			"event_data", e.EventData(),
			"occurred_at", e.OccurredOn(),
		)

		// 这里可以添加具体的业务逻辑
		// 比如：发送欢迎邮件、初始化用户数据等

		return nil
	}
}

func (*UserCreatedListener) AsyncProcess() bool {
	return true // 异步处理
}

// UserLoggedInListener 用户登录事件监听器
type UserLoggedInListener struct{}

func (*UserLoggedInListener) EventHandler() event.Handler {
	return func(e *event.Event) error {
		userData, ok := e.EventData().(*events.UserLoggedIn)
		if !ok {
			return errcode.ErrCoverData.WithCause(fmt.Errorf("cannot convert event data to UserLoggedIn"))
		}

		logger.New(e.Context()).Info("User logged in event processed",
			"user_id", userData.UserId,
			"user_name", userData.UserName,
			"platform", userData.Platform,
			"login_at", userData.LoginAt)

		// 这里可以添加具体的业务逻辑
		// 比如：记录登录日志、更新最后登录时间等

		return nil
	}
}

func (*UserLoggedInListener) AsyncProcess() bool {
	return false // 同步处理
}
