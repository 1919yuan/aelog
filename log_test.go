package aelog_test

import (
	"testing"

	"github.com/1919yuan/aelog"
	"gotest.tools/assert"
)

func TestLogger(t *testing.T) {
	aelog.Debug("Testing: Debug")
	aelog.Info("Testing: Info")
	aelog.Warning("Testing: Warning")
	aelog.Error("Testing: Error")
	// Fatal will call log.Fatalln, so test will fail.
	// logger.Fatal("Testing: Fatal")
	assert.Assert(t, true)
}
