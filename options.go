package logger

type OptionSet struct {
	loggerType LoggerType
	level      int
	outputPath string
}

type Option func(*OptionSet)

func WithOutputPath(path string) Option {
	return func(s *OptionSet) {
		s.outputPath = path
	}
}

func WithLevel(level int) Option {
	return func(s *OptionSet) {
		s.level = level
	}
}

func WithType(t LoggerType) Option {
	return func(s *OptionSet) {
		s.loggerType = t
	}
}
