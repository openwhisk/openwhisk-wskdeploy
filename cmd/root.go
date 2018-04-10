/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/apache/incubator-openwhisk-wskdeploy/deployers"
	"github.com/apache/incubator-openwhisk-wskdeploy/utils"
	"github.com/apache/incubator-openwhisk-wskdeploy/wskderrors"
	"github.com/apache/incubator-openwhisk-wskdeploy/wski18n"
	"github.com/apache/incubator-openwhisk-wskdeploy/wskprint"
	"github.com/spf13/cobra"
)

var stderr = ""
var stdout = ""

// Whisk Deploy has root command: wskdeploy
// wskdeploy is being created using Cobra Library
var RootCmd = &cobra.Command{
	Use:           "wskdeploy",
	SilenceErrors: true,
	SilenceUsage:  true,
	Short:         wski18n.T(wski18n.ID_CMD_DESC_SHORT_ROOT),
	Long:          wski18n.T(wski18n.ID_CMD_DESC_LONG_ROOT),
	RunE:          RootCmdImp,
}

func RootCmdImp(cmd *cobra.Command, args []string) error {
	return Deploy()
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		wskprint.PrintOpenWhiskFromError(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Defining Persistent Flags of Whisk Deploy Root command (wskdeploy)
	// Persistent flags are global in terms of its availability and acceptable
	// with any other Whisk Deploy command e.g. undeploy, export, etc.
	// TODO() Publish command, not completed
	// TODO() Report command, not completed
	// TODO() have a single function that conditionally (i.e., Trace=true) prints ALL Flags
	RootCmd.PersistentFlags().StringVar(&utils.Flags.CfgFile, "config", "", wski18n.T(wski18n.ID_CMD_FLAG_CONFIG))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.ProjectPath, "project", "p", ".", wski18n.T(wski18n.ID_CMD_FLAG_PROJECT))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.ManifestPath, "manifest", "m", "", wski18n.T(wski18n.ID_CMD_FLAG_MANIFEST))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.DeploymentPath, "deployment", "d", "", wski18n.T(wski18n.ID_CMD_FLAG_DEPLOYMENT))
	RootCmd.PersistentFlags().BoolVarP(&utils.Flags.Strict, "strict", "s", false, wski18n.T(wski18n.ID_CMD_FLAG_STRICT))
	RootCmd.PersistentFlags().BoolVarP(&utils.Flags.Preview, "preview", "", false, wski18n.T(wski18n.ID_CMD_FLAG_PREVIEW))
	RootCmd.PersistentFlags().BoolVarP(&utils.Flags.Verbose, "verbose", "v", false, wski18n.T(wski18n.ID_CMD_FLAG_VERBOSE))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.ApiHost, "apihost", "", "", wski18n.T(wski18n.ID_CMD_FLAG_API_HOST))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.Namespace, "namespace", "n", "", wski18n.T(wski18n.ID_CMD_FLAG_NAMESPACE))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.Auth, "auth", "u", "", wski18n.T(wski18n.ID_CMD_FLAG_AUTH_KEY))
	RootCmd.PersistentFlags().StringVar(&utils.Flags.ApiVersion, "apiversion", "", wski18n.T(wski18n.ID_CMD_FLAG_API_VERSION))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.Key, "key", "k", "", wski18n.T(wski18n.ID_CMD_FLAG_KEY_FILE))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.Cert, "cert", "c", "", wski18n.T(wski18n.ID_CMD_FLAG_CERT_FILE))
	RootCmd.PersistentFlags().BoolVarP(&utils.Flags.Managed, "managed", "", false, wski18n.T(wski18n.ID_CMD_FLAG_MANAGED))
	RootCmd.PersistentFlags().StringVarP(&utils.Flags.ProjectName, "projectname", "", "", wski18n.T(wski18n.ID_CMD_FLAG_PROJECTNAME))
	RootCmd.PersistentFlags().BoolVarP(&utils.Flags.Trace, "trace", "t", false, wski18n.T(wski18n.ID_CMD_FLAG_TRACE))
	RootCmd.PersistentFlags().MarkHidden("trace")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	userHome := utils.GetHomeDirectory()
	defaultPath := path.Join(userHome, whisk.DEFAULT_LOCAL_CONFIG)
	if utils.Flags.CfgFile != "" {

		// Read the file as a wskprops file, to check if it is valid.
		_, err := whisk.ReadProps(utils.Flags.CfgFile)
		if err != nil {
			utils.Flags.CfgFile = defaultPath
			warn := wski18n.T(wski18n.ID_WARN_CONFIG_INVALID_X_path_X,
				map[string]interface{}{
					wski18n.KEY_PATH: utils.Flags.CfgFile})
			wskprint.PrintOpenWhiskWarning(warn)
		}

	} else {
		utils.Flags.CfgFile = defaultPath
	}
}

// TODO() add Trace of runtimes found at apihost
func setSupportedRuntimes(apiHost string) {
	op, error := utils.ParseOpenWhisk(apiHost)
	if error == nil {
		utils.SupportedRunTimes = utils.ConvertToMap(op)
		utils.DefaultRunTimes = utils.DefaultRuntimes(op)
		utils.FileExtensionRuntimeKindMap = utils.FileExtensionRuntimes(op)
		utils.FileRuntimeExtensionsMap = utils.FileRuntimeExtensions(op)
	}
}

func displayCommandUsingFilenameMessage(command string, filetype string, path string) {
	msg := wski18n.T(wski18n.ID_MSG_COMMAND_USING_X_cmd_X_filetype_X_path_X,
		map[string]interface{}{
			wski18n.KEY_CMD:       command,
			wski18n.KEY_FILE_TYPE: filetype,
			wski18n.KEY_PATH:      path})
	wskprint.PrintlnOpenWhiskVerbose(utils.Flags.Verbose, msg)
}

func loadDefaultManifestFileFromProjectPath(command string, projectPath string) error {

	if _, err := os.Stat(path.Join(projectPath, utils.ManifestFileNameYaml)); err == nil {
		utils.Flags.ManifestPath = path.Join(projectPath, utils.ManifestFileNameYaml)
	} else if _, err := os.Stat(path.Join(projectPath, utils.ManifestFileNameYml)); err == nil {
		utils.Flags.ManifestPath = path.Join(projectPath, utils.ManifestFileNameYml)
	} else {
		stderr = wski18n.T(wski18n.ID_ERR_MANIFEST_FILE_NOT_FOUND_X_path_X,
			map[string]interface{}{wski18n.KEY_PATH: projectPath})
		return wskderrors.NewErrorManifestFileNotFound(projectPath, stderr)
	}
	displayCommandUsingFilenameMessage(command, wski18n.MANIFEST_FILE, utils.Flags.ManifestPath)
	return nil
}

func loadDefaultDeploymentFileFromProjectPath(command string, projectPath string) error {

	if _, err := os.Stat(path.Join(projectPath, utils.DeploymentFileNameYaml)); err == nil {
		utils.Flags.DeploymentPath = path.Join(projectPath, utils.DeploymentFileNameYaml)
	} else if _, err := os.Stat(path.Join(projectPath, utils.DeploymentFileNameYml)); err == nil {
		utils.Flags.DeploymentPath = path.Join(projectPath, utils.DeploymentFileNameYml)
	}
	displayCommandUsingFilenameMessage(command, wski18n.DEPLOYMENT_FILE, utils.Flags.DeploymentPath)
	return nil
}

func Deploy() error {

	// Convey flags for verbose and trace to Go client
	whisk.SetVerbose(utils.Flags.Verbose)
	whisk.SetDebug(utils.Flags.Trace)

	project_Path := strings.TrimSpace(utils.Flags.ProjectPath)
	if len(project_Path) == 0 {
		project_Path = utils.DEFAULT_PROJECT_PATH
	}
	projectPath, _ := filepath.Abs(project_Path)

	// If manifest filename is not provided, attempt to load default manifests from project path
	// default manifests are manifest.yaml and manifest.yml
	// return failure if none of the default manifest files were found
	if utils.Flags.ManifestPath == "" {
		if err := loadDefaultManifestFileFromProjectPath(wski18n.CMD_DEPLOY, projectPath); err != nil {
			return err
		}
	}

	// If deployment filename is not provided, attempt to load default deployment files from project path
	// default deployments are deployment.yaml and deployment.yml
	// continue processing manifest file, even if none of the default
	// deployment files were found as deployment files are optional
	if utils.Flags.DeploymentPath == "" {
		if err := loadDefaultDeploymentFileFromProjectPath(wski18n.CMD_DEPLOY, projectPath); err != nil {
			return err
		}
	}

	if utils.MayExists(utils.Flags.ManifestPath) {

		// Create an instance of ServiceDeployer
		var deployer = deployers.NewServiceDeployer()
		// Set Project Path, Manifest Path, and Deployment Path of ServiceDeployer
		deployer.ProjectPath = projectPath
		deployer.ManifestPath = utils.Flags.ManifestPath
		deployer.DeploymentPath = utils.Flags.DeploymentPath
		deployer.Preview = utils.Flags.Preview

		// master record of any dependency that has been downloaded
		deployer.DependencyMaster = make(map[string]utils.DependencyRecord)

		// Read credentials from Configuration file, manifest file or deployment file
		clientConfig, error := deployers.NewWhiskConfig(
			utils.Flags.CfgFile,
			utils.Flags.DeploymentPath,
			utils.Flags.ManifestPath)
		if error != nil {
			return error
		}

		whiskClient, error := deployers.CreateNewClient(clientConfig)
		if error != nil {
			return error
		}

		deployer.Client = whiskClient
		deployer.ClientConfig = clientConfig

		// The auth, apihost and namespace have been chosen, so that we can check the supported runtimes here.
		setSupportedRuntimes(clientConfig.Host)

		// Construct Deployment Plan
		err := deployer.ConstructDeploymentPlan()
		if err != nil {
			return err
		}

		// Deploy all OW entities
		err = deployer.Deploy()
		if err != nil {
			return err
		} else {
			return nil
		}

	} else {
		errString := wski18n.T(wski18n.ID_ERR_MANIFEST_FILE_NOT_FOUND_X_path_X,
			map[string]interface{}{wski18n.KEY_PATH: utils.Flags.ManifestPath})
		whisk.Debug(whisk.DbgError, errString)
		return wskderrors.NewErrorManifestFileNotFound(utils.Flags.ManifestPath, errString)
	}

}

func Undeploy() error {

	// Convey flags for verbose and trace to Go client
	whisk.SetVerbose(utils.Flags.Verbose)
	whisk.SetDebug(utils.Flags.Trace)

	if len(utils.Flags.ProjectName) != 0 {
		var deployer = deployers.NewServiceDeployer()
		deployer.Preview = utils.Flags.Preview

		clientConfig, error := deployers.NewWhiskConfig(utils.Flags.CfgFile, "", "")
		if error != nil {
			return error
		}

		whiskClient, error := deployers.CreateNewClient(clientConfig)
		if error != nil {
			return error
		}

		deployer.Client = whiskClient
		deployer.ClientConfig = clientConfig

		// The auth, apihost and namespace have been chosen, so that we can check the supported runtimes here.
		setSupportedRuntimes(clientConfig.Host)

		err := deployer.UnDeployProject()
		if err != nil {
			return err
		}

		return nil
	}

	project_Path := strings.TrimSpace(utils.Flags.ProjectPath)
	if len(project_Path) == 0 {
		project_Path = utils.DEFAULT_PROJECT_PATH
	}
	projectPath, _ := filepath.Abs(project_Path)

	// If manifest filename is not provided, attempt to load default manifests from project path
	if utils.Flags.ManifestPath == "" {
		if err := loadDefaultManifestFileFromProjectPath(wski18n.CMD_UNDEPLOY, projectPath); err != nil {
			return err
		}
	}

	if utils.Flags.DeploymentPath == "" {

		if err := loadDefaultDeploymentFileFromProjectPath(wski18n.CMD_UNDEPLOY, projectPath); err != nil {
			return err
		}
	}

	if utils.FileExists(utils.Flags.ManifestPath) {

		var deployer = deployers.NewServiceDeployer()
		deployer.ProjectPath = utils.Flags.ProjectPath
		deployer.ManifestPath = utils.Flags.ManifestPath
		deployer.DeploymentPath = utils.Flags.DeploymentPath
		deployer.Preview = utils.Flags.Preview

		clientConfig, error := deployers.NewWhiskConfig(utils.Flags.CfgFile, utils.Flags.DeploymentPath, utils.Flags.ManifestPath)
		if error != nil {
			return error
		}

		whiskClient, error := deployers.CreateNewClient(clientConfig)
		if error != nil {
			return error
		}

		deployer.Client = whiskClient
		deployer.ClientConfig = clientConfig

		// The auth, apihost and namespace have been chosen, so that we can check the supported runtimes here.
		setSupportedRuntimes(clientConfig.Host)

		verifiedPlan, err := deployer.ConstructUnDeploymentPlan()
		if err != nil {
			return err
		}

		err = deployer.UnDeploy(verifiedPlan)
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		errString := wski18n.T(wski18n.ID_ERR_MANIFEST_FILE_NOT_FOUND_X_path_X,
			map[string]interface{}{wski18n.KEY_PATH: utils.Flags.ManifestPath})
		return wskderrors.NewErrorManifestFileNotFound(utils.Flags.ManifestPath, errString)
	}
}
