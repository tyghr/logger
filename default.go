package logger

var (
	std Logger

	Debug       func(...any)
	Debugf      func(string, ...any)
	Debugw      func(string, ...any)
	Info        func(...any)
	Infof       func(string, ...any)
	Infow       func(string, ...any)
	Warn        func(...any)
	Warnf       func(string, ...any)
	Warnw       func(string, ...any)
	Error       func(...any)
	Errorf      func(string, ...any)
	Errorw      func(string, ...any)
	Fatal       func(...any)
	Fatalf      func(string, ...any)
	Fatalw      func(string, ...any)
	FailOnError func(error, string) bool
	Sync        func() error
)

func init() {
	Init()
}

func Init(opts ...Option) {
	std = NewLogger(opts...)

	if std != nil {
		Debug = std.Debug
		Debugf = std.Debugf
		Debugw = std.Debugw
		Info = std.Info
		Infof = std.Infof
		Infow = std.Infow
		Warn = std.Warn
		Warnf = std.Warnf
		Warnw = std.Warnw
		Error = std.Error
		Errorf = std.Errorf
		Errorw = std.Errorw
		Fatal = std.Fatal
		Fatalf = std.Fatalf
		Fatalw = std.Fatalw
		FailOnError = std.FailOnError
		Sync = std.Sync
	}
}
