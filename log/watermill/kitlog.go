package watermill

import (
	"github.com/ThreeDotsLabs/watermill"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type kitlogAdapter struct {
	errorLogger  kitlog.Logger
	infoLogger   kitlog.Logger
	debugLogger  kitlog.Logger
	traceLogger  kitlog.Logger
	fields       watermill.LogFields
	traceEnabled bool
}

// WithEnableTrace enables trace logging.
func WithTraceEnabled(enabled bool) Option {
	return func(l *kitlogAdapter) {
		l.traceEnabled = enabled
	}
}

type Option func(*kitlogAdapter)

func (l kitlogAdapter) Error(msg string, err error, fields watermill.LogFields) {
	l.log(l.errorLogger, level.Error, fields.Add(watermill.LogFields{"msg": msg, "err": err}))
}

func (l kitlogAdapter) Info(msg string, fields watermill.LogFields) {
	l.log(l.infoLogger, level.Info, fields.Add(watermill.LogFields{"msg": msg}))
}

func (l kitlogAdapter) Debug(msg string, fields watermill.LogFields) {
	l.log(l.debugLogger, level.Debug, fields.Add(watermill.LogFields{"msg": msg}))
}

func (l kitlogAdapter) Trace(msg string, fields watermill.LogFields) {
	l.log(l.traceLogger, level.Debug, fields.Add(watermill.LogFields{"msg": msg}))
}

func (l kitlogAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &kitlogAdapter{
		errorLogger: l.errorLogger,
		infoLogger:  l.infoLogger,
		debugLogger: l.debugLogger,
		traceLogger: l.traceLogger,
		fields:      l.fields.Add(fields),
	}
}

func (l kitlogAdapter) log(logger kitlog.Logger, logFunc func(kitlog.Logger) kitlog.Logger, fields watermill.LogFields) {
	if logger == nil {
		return
	}

	logFunc(logger).Log(l.keyvalsFromFields(fields)...)
}

func (l kitlogAdapter) keyvalsFromFields(fields watermill.LogFields) []any {
	allFields := l.fields.Add(fields)
	kvs := make([]any, 0, len(allFields)*2)
	for k, v := range allFields {
		kvs = append(kvs, k, v)
	}
	return kvs
}

// NewKitLogger creates a new watermill.LoggerAdapter that wraps a kitlog.Logger.
func NewKitLogger(logger kitlog.Logger, options ...Option) watermill.LoggerAdapter {
	l := &kitlogAdapter{
		errorLogger: logger,
		infoLogger:  logger,
		debugLogger: logger,
	}
	for _, option := range options {
		option(l)
	}

	if l.traceEnabled {
		l.traceLogger = logger
	}

	return l
}
