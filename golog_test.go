package golog

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.Debugln("debug hi")
	logger.Infoln("info hi")
}