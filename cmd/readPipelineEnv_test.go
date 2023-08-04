package cmd

import (
	"github.com/SAP/jenkins-library/pkg/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type readPipelineEnvMockUtils struct {
	*mock.ExecMockRunner
	*mock.FilesMock
}

func newReadPipelineEnvTestsUtils() readPipelineEnvMockUtils {
	utils := readPipelineEnvMockUtils{
		ExecMockRunner: &mock.ExecMockRunner{},
		FilesMock:      &mock.FilesMock{},
	}
	return utils
}

func TestRunReadPipelineEnv(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		// init
		config := readPipelineEnvOptions{}

		utils := newReadPipelineEnvTestsUtils()
		utils.AddFile("file.txt", []byte("dummy content"))

		// test
		err := runReadPipelineEnv(&config, nil, utils)

		// assert
		assert.NoError(t, err)
	})

	t.Run("error path", func(t *testing.T) {
		t.Parallel()
		// init
		config := readPipelineEnvOptions{}

		utils := newReadPipelineEnvTestsUtils()

		// test
		err := runReadPipelineEnv(&config, nil, utils)

		// assert
		assert.EqualError(t, err, "cannot run without important file")
	})
}
