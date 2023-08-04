//go:build unit
// +build unit

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPipelineEnvCommand(t *testing.T) {
	t.Parallel()

	testCmd := ReadPipelineEnvCommand()

	// only high level testing performed - details are tested in step generation procedure
	assert.Equal(t, "readPipelineEnv", testCmd.Use, "command name incorrect")

}
