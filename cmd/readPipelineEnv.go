package cmd

import (
	"encoding/json"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/SAP/jenkins-library/pkg/validation"
)

type readPipelineEnvOptions1 struct {
	GithubToken string `json:"githubToken,omitempty"`
}

func addReadPipelineEnvFlags1(cmd *cobra.Command, stepConfig *readPipelineEnvOptions) {
	cmd.Flags().StringVar(&stepConfig.GithubToken, "githubToken", os.Getenv("PIPER_githubToken"), "GitHub personal access token as per https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
}

// ReadPipelineEnv reads the commonPipelineEnvironment from disk and outputs it as JSON
// func ReadPipelineEnv() *cobra.Command {

// 	const STEP_NAME = "readPipelineEnv"

// 	metadata := readPipelineEnvMetadata()
// 	var stepConfig readPipelineEnvOptions

// 	return &cobra.Command{
// 		Use:   "readPipelineEnv",
// 		Short: "Reads the commonPipelineEnvironment from disk and outputs it as JSON",
// 		PreRun: func(cmd *cobra.Command, args []string) {
// 			path, _ := os.Getwd()
// 			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
// 			log.RegisterHook(fatalHook)
// 		},

// 		Run: func(cmd *cobra.Command, args []string) {
// 			err := runReadPipelineEnv()
// 			if err != nil {
// 				log.Entry().Fatalf("error when writing reading Pipeline environment: %v", err)
// 			}
// 		},
// 	}
// }

func readPipelineEnv(config readPipelineEnvOptions, _ *telemetry.CustomData) error {
	cpe := piperenv.CPEMap{}

	err := cpe.LoadFromDisk(path.Join(GeneralConfig.EnvRootPath, "commonPipelineEnvironment"))
	if err != nil {
		return err
	}

	// encoder := json.NewEncoder(os.Stdout)
	b, err := json.MarshalIndent(&cpe, "", "\t")
	if err != nil {
		return err
	}
	// fmt.Printf("===cpe: %s\n", string(b))
	// encoder.SetIndent("", "\t")
	// if err := encoder.Encode(cpe); err != nil {
	// return err
	// }

	_, err = os.Stdout.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// ShellExecuteCommand Step executes defined script
func ReadPipelineEnv1() *cobra.Command {
	const STEP_NAME = "readPipelineEnv"

	metadata := readPipelineEnvMetadata()
	var stepConfig readPipelineEnvOptions

	var createReadPipelineEnvCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Step executes defined script",
		Long:  `Step executes defined script provided in the 'sources' parameter`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			GeneralConfig.GitHubAccessTokens = ResolveAccessTokens(GeneralConfig.GitHubTokens)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.GithubToken)

			validation, err := validation.New(validation.WithJSONNamesForStructFields(), validation.WithPredefinedErrorMessages())
			if err != nil {
				return err
			}
			if err = validation.ValidateStruct(stepConfig); err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			handler := func() {
				config.RemoveVaultSecretFiles()
			}
			log.DeferExitHandler(handler)
			defer handler()
			// err := runReadPipelineEnv(stepConfig)
			// if err != nil {
			// 	log.Entry().Fatalf("error when writing reading Pipeline environment: %v", err)
			// }
			log.Entry().Info("SUCCESS")
		},
	}

	addReadPipelineEnvFlags(createReadPipelineEnvCmd, &stepConfig)
	return createReadPipelineEnvCmd
}

// retrieve step metadata
func readPipelineEnvMetadata1() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "readPipelineEnv",
			Aliases:     []config.Alias{},
			Description: "description",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "githubTokenCredentialsId", Description: "Jenkins credentials ID containing the github token.", Type: "jenkins"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "githubToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "githubTokenCredentialsId",
								Type: "secret",
							},

							{
								Name:    "githubVaultSecretName",
								Type:    "vaultSecret",
								Default: "github",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "access_token"}},
						Default:   os.Getenv("PIPER_githubToken"),
					},
				},
			},
		},
	}
	return theMetaData
}
