package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

// 提供logger包的门面函数供项目使用, 区别于门面方法 logger.New(ctx).Info(msg, "name", "golang")的方式
// 门面函数使用方式为 logger.Info(ctx, msg, "name", "golang")

var (
	f    *facade
	once sync.Once
)

type facade struct {
	_logger *zap.Logger
}

func Debug(ctx context.Context, msg string, kv ...interface{}) {
	logFacade().log(ctx, zapcore.DebugLevel, msg, kv...)
}

func Info(ctx context.Context, msg string, kv ...interface{}) {
	logFacade().log(ctx, zapcore.InfoLevel, msg, kv...)
}

func Warn(ctx context.Context, msg string, kv ...interface{}) {
	logFacade().log(ctx, zapcore.WarnLevel, msg, kv...)
}

func Error(ctx context.Context, msg string, kv ...interface{}) {
	logFacade().log(ctx, zapcore.ErrorLevel, msg, kv...)
}

func (f *facade) log(ctx context.Context, lvl zapcore.Level, msg string, kv ...interface{}) {
	fields := makeLogFields(ctx, kv...)
	ce := f._logger.Check(lvl, msg)
	ce.Write(fields...)
}

func logFacade() *facade {
	// 使用sync.Once确保初始化代码只执行一次
	once.Do(func() {
		f = &facade{
			_logger: _logger,
		}
	})
	return f
}
