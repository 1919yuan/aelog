package aelog_test

import (
	"testing"

	"github.com/1919yuan/aelog"
	"gotest.tools/assert"
)

func TestSysLogger(t *testing.T) {
	logger := aelog.NewSysLogger()
	defer logger.Close()

	logger.Debug("Testing SysLogger: Debug")
	logger.Info("Testing SysLogger: Info")
	logger.Warning("Testing SysLogger: Warning")
	logger.Error("Testing SysLogger: Error")
	// Fatal will call log.Fatalln, so test will fail.
	// logger.Fatal("Testing SysLogger: Fatal")
	assert.Assert(t, true)
}
