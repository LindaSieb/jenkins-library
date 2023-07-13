// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/gcs"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/SAP/jenkins-library/pkg/validation"
	"github.com/bmatcuk/doublestar"
	"github.com/spf13/cobra"
)

type detectExecuteScanOptions struct {
	Token                       string   `json:"token,omitempty"`
	CodeLocation                string   `json:"codeLocation,omitempty"`
	ProjectName                 string   `json:"projectName,omitempty"`
	Scanners                    []string `json:"scanners,omitempty" validate:"possible-values=signature source"`
	ScanPaths                   []string `json:"scanPaths,omitempty"`
	DependencyPath              string   `json:"dependencyPath,omitempty"`
	Unmap                       bool     `json:"unmap,omitempty"`
	ScanProperties              []string `json:"scanProperties,omitempty"`
	ServerURL                   string   `json:"serverUrl,omitempty"`
	Groups                      []string `json:"groups,omitempty"`
	FailOn                      []string `json:"failOn,omitempty" validate:"possible-values=ALL BLOCKER CRITICAL MAJOR MINOR NONE"`
	VersioningModel             string   `json:"versioningModel,omitempty" validate:"possible-values=major major-minor semantic full"`
	Version                     string   `json:"version,omitempty"`
	CustomScanVersion           string   `json:"customScanVersion,omitempty"`
	ProjectSettingsFile         string   `json:"projectSettingsFile,omitempty"`
	GlobalSettingsFile          string   `json:"globalSettingsFile,omitempty"`
	M2Path                      string   `json:"m2Path,omitempty"`
	InstallArtifacts            bool     `json:"installArtifacts,omitempty"`
	IncludedPackageManagers     []string `json:"includedPackageManagers,omitempty"`
	ExcludedPackageManagers     []string `json:"excludedPackageManagers,omitempty"`
	MavenExcludedScopes         []string `json:"mavenExcludedScopes,omitempty"`
	DetectTools                 []string `json:"detectTools,omitempty"`
	ScanOnChanges               bool     `json:"scanOnChanges,omitempty"`
	CustomEnvironmentVariables  []string `json:"customEnvironmentVariables,omitempty"`
	MinScanInterval             int      `json:"minScanInterval,omitempty"`
	GithubToken                 string   `json:"githubToken,omitempty"`
	CreateResultIssue           bool     `json:"createResultIssue,omitempty"`
	GithubAPIURL                string   `json:"githubApiUrl,omitempty"`
	Owner                       string   `json:"owner,omitempty"`
	Repository                  string   `json:"repository,omitempty"`
	Assignees                   []string `json:"assignees,omitempty"`
	CustomTLSCertificateLinks   []string `json:"customTlsCertificateLinks,omitempty"`
	FailOnSevereVulnerabilities bool     `json:"failOnSevereVulnerabilities,omitempty"`
	BuildTool                   string   `json:"buildTool,omitempty"`
	ExcludedDirectories         []string `json:"excludedDirectories,omitempty"`
	NpmDependencyTypesExcluded  []string `json:"npmDependencyTypesExcluded,omitempty" validate:"possible-values=NONE DEV PEER"`
	NpmArguments                []string `json:"npmArguments,omitempty"`
}

type detectExecuteScanInflux struct {
	step_data struct {
		fields struct {
			detect bool
		}
		tags struct {
		}
	}
	detect_data struct {
		fields struct {
			vulnerabilities       int
			major_vulnerabilities int
			minor_vulnerabilities int
			components            int
			policy_violations     int
		}
		tags struct {
		}
	}
}

func (i *detectExecuteScanInflux) persist(path, resourceName string) {
	measurementContent := []struct {
		measurement string
		valType     string
		name        string
		value       interface{}
	}{
		{valType: config.InfluxField, measurement: "step_data", name: "detect", value: i.step_data.fields.detect},
		{valType: config.InfluxField, measurement: "detect_data", name: "vulnerabilities", value: i.detect_data.fields.vulnerabilities},
		{valType: config.InfluxField, measurement: "detect_data", name: "major_vulnerabilities", value: i.detect_data.fields.major_vulnerabilities},
		{valType: config.InfluxField, measurement: "detect_data", name: "minor_vulnerabilities", value: i.detect_data.fields.minor_vulnerabilities},
		{valType: config.InfluxField, measurement: "detect_data", name: "components", value: i.detect_data.fields.components},
		{valType: config.InfluxField, measurement: "detect_data", name: "policy_violations", value: i.detect_data.fields.policy_violations},
	}

	errCount := 0
	for _, metric := range measurementContent {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(metric.measurement, fmt.Sprintf("%vs", metric.valType), metric.name), metric.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting influx environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Error("failed to persist Influx environment")
	}
}

type detectExecuteScanReports struct {
}

func (p *detectExecuteScanReports) persist(stepConfig detectExecuteScanOptions, gcpJsonKeyFilePath string, gcsBucketId string, gcsFolderPath string, gcsSubFolder string) {
	if gcsBucketId == "" {
		log.Entry().Info("persisting reports to GCS is disabled, because gcsBucketId is empty")
		return
	}
	log.Entry().Info("Uploading reports to Google Cloud Storage...")
	content := []gcs.ReportOutputParam{
		{FilePattern: "**/*BlackDuck_RiskReport.pdf", ParamRef: "", StepResultType: "blackduck-ip"},
		{FilePattern: "**/blackduck-ip.json", ParamRef: "", StepResultType: "blackduck-ip"},
		{FilePattern: "**/toolrun_detectExecute_*.json", ParamRef: "", StepResultType: "blackduck-ip"},
		{FilePattern: "**/piper_detect_policy_violation_report.html", ParamRef: "", StepResultType: "blackduck-ip"},
		{FilePattern: "**/*BlackDuck_RiskReport.pdf", ParamRef: "", StepResultType: "blackduck-security"},
		{FilePattern: "**/detectExecuteScan_policy_*.json", ParamRef: "", StepResultType: "blackduck-security"},
		{FilePattern: "**/piper_detect_vulnerability_report.html", ParamRef: "", StepResultType: "blackduck-security"},
		{FilePattern: "**/toolrun_detectExecute_*.json", ParamRef: "", StepResultType: "blackduck-security"},
		{FilePattern: "**/piper_detect_vulnerability.sarif", ParamRef: "", StepResultType: "blackduck-security"},
		{FilePattern: "**/piper_hub_detect_sbom.xml", ParamRef: "", StepResultType: "blackduck-security"},
	}
	envVars := []gcs.EnvVar{
		{Name: "GOOGLE_APPLICATION_CREDENTIALS", Value: gcpJsonKeyFilePath, Modified: false},
	}
	gcsClient, err := gcs.NewClient(gcs.WithEnvVars(envVars))
	if err != nil {
		log.Entry().Errorf("creation of GCS client failed: %v", err)
		return
	}
	defer gcsClient.Close()
	structVal := reflect.ValueOf(&stepConfig).Elem()
	inputParameters := map[string]string{}
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Type().Field(i)
		if field.Type.String() == "string" {
			paramName := strings.Split(field.Tag.Get("json"), ",")
			paramValue, _ := structVal.Field(i).Interface().(string)
			inputParameters[paramName[0]] = paramValue
		}
	}
	if err := gcs.PersistReportsToGCS(gcsClient, content, inputParameters, gcsFolderPath, gcsBucketId, gcsSubFolder, doublestar.Glob, os.Stat); err != nil {
		log.Entry().Errorf("failed to persist reports: %v", err)
	}
}

// DetectExecuteScanCommand Executes Synopsys Detect scan
func DetectExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "detectExecuteScan"

	metadata := detectExecuteScanMetadata()
	var stepConfig detectExecuteScanOptions
	var startTime time.Time
	var influx detectExecuteScanInflux
	var reports detectExecuteScanReports
	var logCollector *log.CollectorHook
	var splunkClient *splunk.Splunk
	telemetryClient := &telemetry.Telemetry{}

	var createDetectExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Executes Synopsys Detect scan",
		Long: `This step executes [Synopsys Detect](https://community.synopsys.com/s/document-item?bundleId=integrations-detect&topicId=introduction.html&_LANG=enus) scans.
Synopsys Detect command line utlity can be used to run various scans including BlackDuck and Polaris scans. This step allows users to run BlackDuck scans by default.
Please configure your BlackDuck server Url using the serverUrl parameter and the API token of your user using the apiToken parameter for this step.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
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
			log.RegisterSecret(stepConfig.Token)
			log.RegisterSecret(stepConfig.GithubToken)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient = &splunk.Splunk{}
				logCollector = &log.CollectorHook{CorrelationID: GeneralConfig.CorrelationID}
				log.RegisterHook(logCollector)
			}

			if err = log.RegisterANSHookIfConfigured(GeneralConfig.CorrelationID); err != nil {
				log.Entry().WithError(err).Warn("failed to set up SAP Alert Notification Service log hook")
			}

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
			stepTelemetryData := telemetry.CustomData{}
			stepTelemetryData.ErrorCode = "1"
			handler := func() {
				influx.persist(GeneralConfig.EnvRootPath, "influx")
				reports.persist(stepConfig, GeneralConfig.GCPJsonKeyFilePath, GeneralConfig.GCSBucketId, GeneralConfig.GCSFolderPath, GeneralConfig.GCSSubFolder)
				config.RemoveVaultSecretFiles()
				stepTelemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				stepTelemetryData.ErrorCategory = log.GetErrorCategory().String()
				stepTelemetryData.PiperCommitHash = GitCommit
				telemetryClient.SetData(&stepTelemetryData)
				telemetryClient.Send()
				if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
					splunkClient.Initialize(GeneralConfig.CorrelationID,
						GeneralConfig.HookConfig.SplunkConfig.Dsn,
						GeneralConfig.HookConfig.SplunkConfig.Token,
						GeneralConfig.HookConfig.SplunkConfig.Index,
						GeneralConfig.HookConfig.SplunkConfig.SendLogs)
					splunkClient.Send(telemetryClient.GetData(), logCollector)
				}
				if len(GeneralConfig.HookConfig.SplunkConfig.ProdCriblEndpoint) > 0 {
					splunkClient.Initialize(GeneralConfig.CorrelationID,
						GeneralConfig.HookConfig.SplunkConfig.ProdCriblEndpoint,
						GeneralConfig.HookConfig.SplunkConfig.ProdCriblToken,
						GeneralConfig.HookConfig.SplunkConfig.ProdCriblIndex,
						GeneralConfig.HookConfig.SplunkConfig.SendLogs)
					splunkClient.Send(telemetryClient.GetData(), logCollector)
				}
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetryClient.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			detectExecuteScan(stepConfig, &stepTelemetryData, &influx)
			stepTelemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addDetectExecuteScanFlags(createDetectExecuteScanCmd, &stepConfig)
	return createDetectExecuteScanCmd
}

func addDetectExecuteScanFlags(cmd *cobra.Command, stepConfig *detectExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.Token, "token", os.Getenv("PIPER_token"), "Api token to be used for connectivity with Synopsis Detect server.")
	cmd.Flags().StringVar(&stepConfig.CodeLocation, "codeLocation", os.Getenv("PIPER_codeLocation"), "An override for the name Detect will use for the scan file it creates.")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", os.Getenv("PIPER_projectName"), "Name of the Synopsis Detect (formerly BlackDuck) project.")
	cmd.Flags().StringSliceVar(&stepConfig.Scanners, "scanners", []string{`signature`}, "List of scanners to be used for Synopsis Detect (formerly BlackDuck) scan.")
	cmd.Flags().StringSliceVar(&stepConfig.ScanPaths, "scanPaths", []string{`.`}, "List of paths which should be scanned by the Synopsis Detect (formerly BlackDuck) scan.")
	cmd.Flags().StringVar(&stepConfig.DependencyPath, "dependencyPath", `.`, "Absolute Path of the dependency management file of the project. This path represents the folder which contains the pom file, package.json etc. If the project contains multiple pom files, provide the path to the parent pom file or the base folder of the project")
	cmd.Flags().BoolVar(&stepConfig.Unmap, "unmap", false, "Unmap flag will unmap all previous code locations and keep only the current scan results in the specified project version. Set this parameter to true, when the project version needs to store only the latest scan results.")
	cmd.Flags().StringSliceVar(&stepConfig.ScanProperties, "scanProperties", []string{`--blackduck.signature.scanner.memory=4096`, `--detect.timeout=6000`, `--blackduck.trust.cert=true`, `--logging.level.com.synopsys.integration=DEBUG`, `--detect.maven.excluded.scopes=test`}, "Properties passed to the Synopsis Detect (formerly BlackDuck) scan. You can find details in the [Synopsis Detect documentation](https://community.synopsys.com/s/document-item?bundleId=integrations-detect&topicId=properties%2Fall-properties.html&_LANG=enus)")
	cmd.Flags().StringVar(&stepConfig.ServerURL, "serverUrl", os.Getenv("PIPER_serverUrl"), "Server URL to the Synopsis Detect (formerly BlackDuck) Server.")
	cmd.Flags().StringSliceVar(&stepConfig.Groups, "groups", []string{}, "Users groups to be assigned for the Project")
	cmd.Flags().StringSliceVar(&stepConfig.FailOn, "failOn", []string{`BLOCKER`}, "Mark the current build as fail based on the policy categories applied.")
	cmd.Flags().StringVar(&stepConfig.VersioningModel, "versioningModel", `major`, "The versioning model used for result reporting (based on the artifact version). Example 1.2.3 using `major` will result in version 1")
	cmd.Flags().StringVar(&stepConfig.Version, "version", os.Getenv("PIPER_version"), "Defines the version number of the artifact being build in the pipeline. It is used as source for the Detect version.")
	cmd.Flags().StringVar(&stepConfig.CustomScanVersion, "customScanVersion", os.Getenv("PIPER_customScanVersion"), "A custom version used along with the uploaded scan results.")
	cmd.Flags().StringVar(&stepConfig.ProjectSettingsFile, "projectSettingsFile", os.Getenv("PIPER_projectSettingsFile"), "Path or url to the mvn settings file that should be used as project settings file.")
	cmd.Flags().StringVar(&stepConfig.GlobalSettingsFile, "globalSettingsFile", os.Getenv("PIPER_globalSettingsFile"), "Path or url to the mvn settings file that should be used as global settings file")
	cmd.Flags().StringVar(&stepConfig.M2Path, "m2Path", os.Getenv("PIPER_m2Path"), "Path to the location of the local repository that should be used.")
	cmd.Flags().BoolVar(&stepConfig.InstallArtifacts, "installArtifacts", false, "If enabled, it will install all artifacts to the local maven repository to make them available before running detect. This is required if any maven module has dependencies to other modules in the repository and they were not installed before.")
	cmd.Flags().StringSliceVar(&stepConfig.IncludedPackageManagers, "includedPackageManagers", []string{}, "The package managers that need to be included for this scan. Providing the package manager names with this parameter will ensure that the build descriptor file of that package manager will be searched in the scan folder For the complete list of possible values for this parameter, please refer [Synopsys detect documentation](https://community.synopsys.com/s/document-item?bundleId=integrations-detect&topicId=properties%2Fconfiguration%2Fdetector.html&_LANG=enus&anchor=detector-types-included-advanced)")
	cmd.Flags().StringSliceVar(&stepConfig.ExcludedPackageManagers, "excludedPackageManagers", []string{}, "The package managers that need to be excluded for this scan. Providing the package manager names with this parameter will ensure that the build descriptor file of that package manager will be ignored in the scan folder For the complete list of possible values for this parameter, please refer [Synopsys detect documentation](https://community.synopsys.com/s/document-item?bundleId=integrations-detect&topicId=properties%2Fconfiguration%2Fdetector.html&_LANG=enus&anchor=detector-types-excluded-advanced)")
	cmd.Flags().StringSliceVar(&stepConfig.MavenExcludedScopes, "mavenExcludedScopes", []string{}, "The maven scopes that need to be excluded from the scan. For example, setting the value 'test' will exclude all components which are defined with a test scope in maven")
	cmd.Flags().StringSliceVar(&stepConfig.DetectTools, "detectTools", []string{}, "The type of BlackDuck scanners to include while running the BlackDuck scan. By default All scanners are included. For the complete list of possible values, Please refer [Synopsys detect documentation](https://community.synopsys.com/s/document-item?bundleId=integrations-detect&topicId=properties%2Fconfiguration%2Fpaths.html&_LANG=enus&anchor=detect-tools-included)")
	cmd.Flags().BoolVar(&stepConfig.ScanOnChanges, "scanOnChanges", false, "This flag determines if the scan is submitted to the server. If set to true, then the scan request is submitted to the server only when changes are detected in the Open Source Bill of Materials If the flag is set to false, then the scan request is submitted to server regardless of any changes. For more details please refer to the [documentation](https://github.com/blackducksoftware/detect_rescan/blob/master/README.md)")
	cmd.Flags().StringSliceVar(&stepConfig.CustomEnvironmentVariables, "customEnvironmentVariables", []string{}, "A list of environment variables which can be set to prepare the environment to run a BlackDuck scan. This includes a list of environment variables defined by Synopsys. The full list can be found [here](https://community.synopsys.com/s/document-item?bundleId=integrations-detect&topicId=configuring%2Fenvvars.html&_LANG=enus) This list affects the detect script downloaded while running the scan. Right now only detect7.sh is available for downloading")
	cmd.Flags().IntVar(&stepConfig.MinScanInterval, "minScanInterval", 0, "This parameter controls the frequency (in number of hours) at which the signature scan is re-submitted for scan. When set to a value greater than 0, the signature scans are skipped until the specified number of hours has elapsed since the last signature scan.")
	cmd.Flags().StringVar(&stepConfig.GithubToken, "githubToken", os.Getenv("PIPER_githubToken"), "GitHub personal access token as per https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
	cmd.Flags().BoolVar(&stepConfig.CreateResultIssue, "createResultIssue", false, "Activate creation of a result issue in GitHub.")
	cmd.Flags().StringVar(&stepConfig.GithubAPIURL, "githubApiUrl", `https://api.github.com`, "Set the GitHub API URL.")
	cmd.Flags().StringVar(&stepConfig.Owner, "owner", os.Getenv("PIPER_owner"), "Set the GitHub organization.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Set the GitHub repository.")
	cmd.Flags().StringSliceVar(&stepConfig.Assignees, "assignees", []string{``}, "Defines the assignees for the Github Issue created/updated with the results of the scan as a list of login names.")
	cmd.Flags().StringSliceVar(&stepConfig.CustomTLSCertificateLinks, "customTlsCertificateLinks", []string{}, "List of download links to custom TLS certificates. This is required to ensure trusted connections to instances with repositories (like nexus) when publish flag is set to true.")
	cmd.Flags().BoolVar(&stepConfig.FailOnSevereVulnerabilities, "failOnSevereVulnerabilities", true, "Whether to fail the step on severe vulnerabilties or not")
	cmd.Flags().StringVar(&stepConfig.BuildTool, "buildTool", os.Getenv("PIPER_buildTool"), "Defines the tool which is used for building the artifact.")
	cmd.Flags().StringSliceVar(&stepConfig.ExcludedDirectories, "excludedDirectories", []string{}, "List of directories which should be excluded from the scan.")
	cmd.Flags().StringSliceVar(&stepConfig.NpmDependencyTypesExcluded, "npmDependencyTypesExcluded", []string{}, "List of npm dependency types which Detect should exclude from the BOM.")
	cmd.Flags().StringSliceVar(&stepConfig.NpmArguments, "npmArguments", []string{}, "List of additional arguments that Detect will add at then end of the npm ls command line when Detect executes the NPM CLI Detector on an NPM project.")

	cmd.MarkFlagRequired("token")
	cmd.MarkFlagRequired("projectName")
	cmd.MarkFlagRequired("serverUrl")
}

// retrieve step metadata
func detectExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "detectExecuteScan",
			Aliases:     []config.Alias{},
			Description: "Executes Synopsys Detect scan",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "detectTokenCredentialsId", Description: "Jenkins 'Secret text' credentials ID containing the API token used to authenticate with the Synopsis Detect (formerly BlackDuck) Server.", Type: "jenkins", Aliases: []config.Alias{{Name: "apiTokenCredentialsId", Deprecated: false}}},
					{Name: "githubTokenCredentialsId", Description: "Jenkins 'Secret text' credentials ID containing token to authenticate to GitHub.", Type: "jenkins"},
				},
				Resources: []config.StepResources{
					{Name: "buildDescriptor", Type: "stash"},
					{Name: "checkmarx", Type: "stash"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "token",
						ResourceRef: []config.ResourceReference{
							{
								Name: "detectTokenCredentialsId",
								Type: "secret",
							},

							{
								Name:    "detectVaultSecretName",
								Type:    "vaultSecret",
								Default: "detect",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "blackduckToken"}, {Name: "detectToken"}, {Name: "apiToken", Deprecated: true}, {Name: "detect/apiToken", Deprecated: true}},
						Default:   os.Getenv("PIPER_token"),
					},
					{
						Name:        "codeLocation",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_codeLocation"),
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "detect/projectName"}},
						Default:     os.Getenv("PIPER_projectName"),
					},
					{
						Name:        "scanners",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/scanners"}},
						Default:     []string{`signature`},
					},
					{
						Name:        "scanPaths",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/scanPaths"}},
						Default:     []string{`.`},
					},
					{
						Name:        "dependencyPath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/dependencyPath"}},
						Default:     `.`,
					},
					{
						Name:        "unmap",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/unmap"}},
						Default:     false,
					},
					{
						Name:        "scanProperties",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/scanProperties"}},
						Default:     []string{`--blackduck.signature.scanner.memory=4096`, `--detect.timeout=6000`, `--blackduck.trust.cert=true`, `--logging.level.com.synopsys.integration=DEBUG`, `--detect.maven.excluded.scopes=test`},
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "detect/serverUrl"}},
						Default:     os.Getenv("PIPER_serverUrl"),
					},
					{
						Name:        "groups",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/groups"}},
						Default:     []string{},
					},
					{
						Name:        "failOn",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/failOn"}},
						Default:     []string{`BLOCKER`},
					},
					{
						Name:        "versioningModel",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "GENERAL", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `major`,
					},
					{
						Name: "version",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "artifactVersion",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "projectVersion"}, {Name: "detect/projectVersion"}},
						Default:   os.Getenv("PIPER_version"),
					},
					{
						Name:        "customScanVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STAGES", "STEPS", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_customScanVersion"),
					},
					{
						Name:        "projectSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/projectSettingsFile"}},
						Default:     os.Getenv("PIPER_projectSettingsFile"),
					},
					{
						Name:        "globalSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/globalSettingsFile"}},
						Default:     os.Getenv("PIPER_globalSettingsFile"),
					},
					{
						Name:        "m2Path",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/m2Path"}},
						Default:     os.Getenv("PIPER_m2Path"),
					},
					{
						Name:        "installArtifacts",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     false,
					},
					{
						Name:        "includedPackageManagers",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/includedPackageManagers"}},
						Default:     []string{},
					},
					{
						Name:        "excludedPackageManagers",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/excludedPackageManagers"}},
						Default:     []string{},
					},
					{
						Name:        "mavenExcludedScopes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/mavenExcludedScopes"}},
						Default:     []string{},
					},
					{
						Name:        "detectTools",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/detectTools"}},
						Default:     []string{},
					},
					{
						Name:        "scanOnChanges",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     false,
					},
					{
						Name:        "customEnvironmentVariables",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     []string{},
					},
					{
						Name:        "minScanInterval",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     0,
					},
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
					{
						Name: "createResultIssue",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "custom/isOptimizedAndScheduled",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:        "githubApiUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `https://api.github.com`,
					},
					{
						Name: "owner",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "github/owner",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "githubOrg"}},
						Default:   os.Getenv("PIPER_owner"),
					},
					{
						Name: "repository",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "github/repository",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "githubRepo"}},
						Default:   os.Getenv("PIPER_repository"),
					},
					{
						Name:        "assignees",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     []string{``},
					},
					{
						Name:        "customTlsCertificateLinks",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     []string{},
					},
					{
						Name:        "failOnSevereVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name: "buildTool",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "buildTool",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_buildTool"),
					},
					{
						Name:        "excludedDirectories",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/excludedDirectories"}},
						Default:     []string{},
					},
					{
						Name:        "npmDependencyTypesExcluded",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/npmDependencyTypesExcluded"}},
						Default:     []string{},
					},
					{
						Name:        "npmArguments",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/npmArguments"}},
						Default:     []string{},
					},
				},
			},
			Containers: []config.Container{
				{Name: "openjdk", Image: "openjdk:11", WorkingDir: "/root", Options: []config.Option{{Name: "-u", Value: "0"}}},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "influx",
						Type: "influx",
						Parameters: []map[string]interface{}{
							{"name": "step_data", "fields": []map[string]string{{"name": "detect"}}},
							{"name": "detect_data", "fields": []map[string]string{{"name": "vulnerabilities"}, {"name": "major_vulnerabilities"}, {"name": "minor_vulnerabilities"}, {"name": "components"}, {"name": "policy_violations"}}},
						},
					},
					{
						Name: "reports",
						Type: "reports",
						Parameters: []map[string]interface{}{
							{"filePattern": "**/*BlackDuck_RiskReport.pdf", "type": "blackduck-ip"},
							{"filePattern": "**/blackduck-ip.json", "type": "blackduck-ip"},
							{"filePattern": "**/toolrun_detectExecute_*.json", "type": "blackduck-ip"},
							{"filePattern": "**/piper_detect_policy_violation_report.html", "type": "blackduck-ip"},
							{"filePattern": "**/*BlackDuck_RiskReport.pdf", "type": "blackduck-security"},
							{"filePattern": "**/detectExecuteScan_policy_*.json", "type": "blackduck-security"},
							{"filePattern": "**/piper_detect_vulnerability_report.html", "type": "blackduck-security"},
							{"filePattern": "**/toolrun_detectExecute_*.json", "type": "blackduck-security"},
							{"filePattern": "**/piper_detect_vulnerability.sarif", "type": "blackduck-security"},
							{"filePattern": "**/piper_hub_detect_sbom.xml", "type": "blackduck-security"},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
