package event

import (
	"sync"
	"time"

	"github.com/go-study-lab/go-mall/event/events"
	"github.com/go-study-lab/go-mall/event/listeners"
	"github.com/kevinyan815/event"
)

var registerOnce sync.Once

// 初始化事件分发器
func init() {
	// 设置事件日志Logger
	event.SetLogger(NewEventLogger())
	// 注册用户相关事件监听器
	RegisterEventListeners()
}

func RegisterEventListeners() {
	registerOnce.Do(func() {
		// 注册用户相关事件监听器
		dispatcher := event.Dispatcher()
		dispatcher.Subscribe(
			event.NewEvent(&events.UserCreated{}),
			&listeners.UserCreatedListener{},
		)
		// 注册用户登录事件监听器
		dispatcher.Subscribe(
			event.NewEvent(&events.UserLoggedIn{}),
			&listeners.UserLoggedInListener{},
		)
	})
}

func ShutdownEventDispatcher(waitTimeout time.Duration) error {
	if waitTimeout == 0 || waitTimeout > 30*time.Second {
		waitTimeout = 30 * time.Second
	}
	return event.Dispatcher().Shutdown(waitTimeout)
}
