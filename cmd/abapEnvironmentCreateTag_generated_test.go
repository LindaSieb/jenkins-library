//go:build unit
// +build unit

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbapEnvironmentCreateTagCommand(t *testing.T) {
	t.Parallel()

	testCmd := AbapEnvironmentCreateTagCommand()

	// only high level testing performed - details are tested in step generation procedure
	assert.Equal(t, "abapEnvironmentCreateTag", testCmd.Use, "command name incorrect")

}
