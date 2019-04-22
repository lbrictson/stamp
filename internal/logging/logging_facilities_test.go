package logging

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	// This test will panic if the logger fails to setup
	Init()
	Logger.Info("test")
}
