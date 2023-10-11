package logger_test

import (
	"testing"

	log "github.com/tyghr/logger"
)

func TestStdLogger(t *testing.T) {
	log.Debug("debug")
	log.Debugf("debugf %s %s", "a", "b")
	log.Debugw("debugw",
		"a", "1",
		"b", "2")

	log.Info("info")
	log.Infof("infof %s %s", "a", "b")
	log.Infow("infow",
		"a", "1",
		"b", "2")

	log.Warn("warn")
	log.Warnf("warnf %s %s", "a", "b")
	log.Warnw("warnw",
		"a", "1",
		"b", "2")

	log.Error("error")
	log.Errorf("errorf %s %s", "a", "b")
	log.Errorw("errorw",
		"a", "1",
		"b", "2")

	// log.Fatal("fatal")
	// log.Fatalf("fatalf %s %s", "a", "b")
	// log.Fatalw("fatalw",
	// 	"a", "1",
	// 	"b", "2")
}
