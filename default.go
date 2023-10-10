package logger

var std = NewLogger()

func Debug(args ...interface{})              { std.Debug(args...) }
func Debugf(fmt string, args ...interface{}) { std.Debugf(fmt, args...) }
func Debugw(fmt string, args ...interface{}) { std.Debugw(fmt, args...) }
func Info(args ...interface{})               { std.Info(args...) }
func Infof(fmt string, args ...interface{})  { std.Infof(fmt, args...) }
func Infow(fmt string, args ...interface{})  { std.Infow(fmt, args...) }
func Warn(args ...interface{})               { std.Warn(args...) }
func Warnf(fmt string, args ...interface{})  { std.Warnf(fmt, args...) }
func Warnw(fmt string, args ...interface{})  { std.Warnw(fmt, args...) }
func Error(args ...interface{})              { std.Error(args...) }
func Errorf(fmt string, args ...interface{}) { std.Errorf(fmt, args...) }
func Errorw(fmt string, args ...interface{}) { std.Errorw(fmt, args...) }
func Fatal(args ...interface{})              { std.Fatal(args...) }
func Fatalf(fmt string, args ...interface{}) { std.Fatalf(fmt, args...) }
func Fatalw(fmt string, args ...interface{}) { std.Fatalw(fmt, args...) }
func FailOnError(err error, msg string) bool { return std.FailOnError(err, msg) }
func Sync() error                            { return std.Sync() }
