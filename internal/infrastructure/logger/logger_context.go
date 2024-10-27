package logger

import (
	"context"
	"log/slog"
)

type LogCtx struct {
	log slogLogger
	ctx context.Context
}

func (l LogCtx) WithError(err error) LogCtx {
	if err == nil {
		return l
	}

	return LogCtx{
		log: l.log.With(slog.Any(errKey, err)),
		ctx: l.ctx,
	}
}

func (l LogCtx) WithFields(fields map[string]any) LogCtx {
	var attrs []any

	for k, f := range fields {
		attrs = append(attrs, slog.Any(k, f))
	}

	return LogCtx{
		log: l.log.With(attrs...),
		ctx: l.ctx,
	}
}

func (l LogCtx) Debug(msg string) {
	l.log.DebugContext(l.ctx, msg)
}

func (l LogCtx) Info(msg string) {
	l.log.InfoContext(l.ctx, msg)
}

func (l LogCtx) Warning(msg string) {
	l.log.WarnContext(l.ctx, msg)
}

func (l LogCtx) Error(msg string) {
	l.log.ErrorContext(l.ctx, msg)
}
