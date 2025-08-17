package event

import (
	"context"
	"github.com/go-study-lab/go-mall/event/events"
	"github.com/kevinyan815/event"
)

// DispatchUserCreated 分发用户创建事件
func DispatchUserCreated(ctx context.Context, userId int64, userName string) {
	event.Dispatcher().Dispatch(
		event.NewEventWithContext(ctx, &events.UserCreated{
			UserId:   userId,
			UserName: userName,
		}),
	)
}

// DispatchUserLoggedIn 分发用户登录事件
func DispatchUserLoggedIn(ctx context.Context, userId int64, userName, platform, loginAt string) {
	event.Dispatcher().Dispatch(
		event.NewEventWithContext(ctx, &events.UserLoggedIn{
			UserId:   userId,
			UserName: userName,
			Platform: platform,
			LoginAt:  loginAt,
		}),
	)
}
