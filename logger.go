package logger

import (
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ColoredConsoleLogger LoggerType = iota
	ConsoleLogger
	ServiceLogger
)

type LoggerType int

//SugarLogger ...
type CustomLogger struct {
	zap.SugaredLogger
}

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugw(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Infow(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnw(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorw(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalw(string, ...interface{})
	FailOnError(error, string) bool
	Sync() error
}

// NewLogger
// Defaults opts are:
//   ConsoleLogger type,
//   DebugLevel,
//   output to "stdout"
func NewLogger(opts ...Option) Logger {
	optSet := OptionSet{
		loggerType: ConsoleLogger,
		level:      DebugLevel,
		outputPath: "stdout",
	}
	for _, opt := range opts {
		opt(&optSet)
	}

	var (
		levelEnc  zapcore.LevelEncoder
		timeEnc   zapcore.TimeEncoder
		callerEnc zapcore.CallerEncoder
	)
	callerTrimmedPathFunc := func(caller zapcore.EntryCaller) string {
		idx := strings.LastIndexByte(caller.Function, '/')
		idx += strings.IndexByte(caller.Function[idx+1:], '.')
		return caller.TrimmedPath() + ":" + caller.Function[idx+2:]
	}

	switch optSet.loggerType {
	case ColoredConsoleLogger:
		levelEnc = zapcore.CapitalColorLevelEncoder
		timeEnc = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(Colored("l_black", t.Format("2006.01.02  15:04:05 .000")))
		}
		callerEnc = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(Colored("l_black", callerTrimmedPathFunc(caller)))
		}
	case ConsoleLogger:
		levelEnc = zapcore.CapitalLevelEncoder
		timeEnc = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006.01.02  15:04:05 .000"))
		}
		callerEnc = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(callerTrimmedPathFunc(caller))
		}
	case ServiceLogger:
		levelEnc = zapcore.CapitalLevelEncoder
		callerEnc = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(callerTrimmedPathFunc(caller))
		}
	}

	encConfig := zapcore.EncoderConfig{
		LevelKey:     "level",
		EncodeLevel:  levelEnc,
		CallerKey:    "caller",
		EncodeCaller: callerEnc,
		MessageKey:   "message",
		//EncodeDuration: zapcore.NanosDurationEncoder,
		ConsoleSeparator: "  ",
	}
	if optSet.loggerType != ServiceLogger {
		encConfig.TimeKey = "timestamp"
		encConfig.EncodeTime = timeEnc
	}

	lgr, err := zap.Config{
		Encoding:      "console",
		Level:         zap.NewAtomicLevelAt(zapcore.Level(optSet.level)),
		OutputPaths:   []string{optSet.outputPath},
		EncoderConfig: encConfig,
	}.Build()
	if err != nil {
		return nil
	}

	defer func() {
		_ = lgr.Sync()
	}()

	return &CustomLogger{*lgr.Sugar()}
}

func (Logger *CustomLogger) FailOnError(err error, msg string) bool {
	if err != nil {
		Logger.Errorf("%s: %s", msg, err)
		return true
	}
	return false
}
