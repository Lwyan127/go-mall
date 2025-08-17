package event

import (
	"context"
	"github.com/go-study-lab/go-mall/common/logger"
	"github.com/kevinyan815/event"
)

type EventLogger struct {
}

func NewEventLogger() *EventLogger {
	return &EventLogger{}
}

func (l *EventLogger) Debug(ctx context.Context, msg string, params ...event.LogParam) {
	logger.Debug(ctx, msg, convertParams(params)...)
}

func (l *EventLogger) Info(ctx context.Context, msg string, params ...event.LogParam) {
	logger.Info(ctx, msg, convertParams(params)...)
}

func (l *EventLogger) Warn(ctx context.Context, msg string, params ...event.LogParam) {
	logger.Warn(ctx, msg, convertParams(params)...)
}

func (l *EventLogger) Error(ctx context.Context, msg string, params ...event.LogParam) {
	logger.Error(ctx, msg, convertParams(params)...)
}

func convertParams(params []event.LogParam) []interface{} {
	result := make([]interface{}, 0, len(params)*2)
	for _, p := range params {
		result = append(result, p.Key, p.Value)
	}
	return result
}
