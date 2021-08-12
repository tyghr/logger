package logger

type NilLogger struct{}

func (nl *NilLogger) Debug(...interface{}) {}

func (nl *NilLogger) Debugf(string, ...interface{}) {}

func (nl *NilLogger) Debugw(string, ...interface{}) {}

func (nl *NilLogger) Info(...interface{}) {}

func (nl *NilLogger) Infof(string, ...interface{}) {}

func (nl *NilLogger) Infow(string, ...interface{}) {}

func (nl *NilLogger) Warn(...interface{}) {}

func (nl *NilLogger) Warnf(string, ...interface{}) {}

func (nl *NilLogger) Warnw(string, ...interface{}) {}

func (nl *NilLogger) Error(...interface{}) {}

func (nl *NilLogger) Errorf(string, ...interface{}) {}

func (nl *NilLogger) Errorw(string, ...interface{}) {}

func (nl *NilLogger) Fatal(...interface{}) {}

func (nl *NilLogger) Fatalf(string, ...interface{}) {}

func (nl *NilLogger) Fatalw(string, ...interface{}) {}

func (nl *NilLogger) FailOnError(err error, msg string) bool {
	return err != nil
}

func (nl *NilLogger) Sync() error {
	return nil
}
