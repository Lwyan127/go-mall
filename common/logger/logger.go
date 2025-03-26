package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

// 使用方式为 logger.New(ctx).Info(msg, "name", "golang", "val", 8)

type logger struct {
	ctx     context.Context
	traceId string
	spanId  string
	pSpanId string
	_logger *zap.Logger
}

func (l *logger) Debug(msg string, kv ...interface{}) {
	l.log(zapcore.DebugLevel, msg, kv...)
}

func (l *logger) Info(msg string, kv ...interface{}) {
	l.log(zapcore.InfoLevel, msg, kv...)
}

func (l *logger) Warn(msg string, kv ...interface{}) {
	l.log(zapcore.WarnLevel, msg, kv...)
}

func (l *logger) Error(msg string, kv ...interface{}) {
	l.log(zapcore.ErrorLevel, msg, kv...)
}

// kv 应该是成对的数据, 类似: name,张三,age,10,...
func (l *logger) log(lvl zapcore.Level, msg string, kv ...interface{}) {
	fields := makeLogFields(l.ctx, kv...)
	ce := l._logger.Check(lvl, msg)
	ce.Write(fields...)
}

// 组装日志行中的字段
func makeLogFields(ctx context.Context, kv ...interface{}) []zap.Field {
	// 保证日志信息以键值对的形式成对出现
	if len(kv)%2 != 0 {
		kv = append(kv, "unknown")
	}
	// 设置日志行信息中的追踪字段
	kv = append(kv, "traceid", ctx.Value("traceid"), "spanid", ctx.Value("spanid"), "pspanid", ctx.Value("pspanid"))
	// 增加日志调用者信息
	funcName, file, line := getLoggerCallerInfo()
	kv = append(kv, "func", funcName, "file", file, "line", line)
	fields := make([]zap.Field, 0, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		k := fmt.Sprintf("%v", kv[i])
		fields = append(fields, zap.Any(k, kv[i+1]))
	}

	return fields
}

// getLoggerCallerInfo 日志调用者信息 -- 方法名, 文件名, 行号
func getLoggerCallerInfo() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(4) // 回溯拿调用日志方法的业务函数的信息
	if !ok {
		return
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}

func New(ctx context.Context) *logger {
	var traceId, spanId, pSpanId string
	if ctx.Value("traceid") != nil {
		traceId = ctx.Value("traceid").(string)
	}
	if ctx.Value("spanid") != nil {
		spanId = ctx.Value("spanid").(string)
	}
	if ctx.Value("psapnid") != nil {
		pSpanId = ctx.Value("pspanid").(string)
	}

	return &logger{
		ctx:     ctx,
		traceId: traceId,
		spanId:  spanId,
		pSpanId: pSpanId,
		_logger: _logger,
	}
}
