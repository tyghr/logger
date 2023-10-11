package logger

var (
	std = NewLogger()

	Debug       = std.Debug
	Debugf      = std.Debugf
	Debugw      = std.Debugw
	Info        = std.Info
	Infof       = std.Infof
	Infow       = std.Infow
	Warn        = std.Warn
	Warnf       = std.Warnf
	Warnw       = std.Warnw
	Error       = std.Error
	Errorf      = std.Errorf
	Errorw      = std.Errorw
	Fatal       = std.Fatal
	Fatalf      = std.Fatalf
	Fatalw      = std.Fatalw
	FailOnError = std.FailOnError
	Sync        = std.Sync
)
