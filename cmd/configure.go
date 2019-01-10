// Copyright © 2018 Peter Souter p.morsou@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	input "github.com/tcnksm/go-input"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your TFE credentials",
	Long: `Prompts for your TFE API credentials, then writes them to
	a configuration file (defaults to ~/.tgc.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchedTfeURL, fetchedTfeAPIToken, err := GetConfigValuesFromPrompts(os.Stdin, os.Stdout)
		if err != nil {
			log.Fatal("Failed to get value:", err)
			os.Exit(1)
		}
		CreateConfigFileFromValues(fetchedTfeURL, fetchedTfeAPIToken)
	},
}

// GetConfigValuesFromPrompts prompts the user for input then returns them as strings
func GetConfigValuesFromPrompts(stdin io.Reader, stdout io.Writer) (string, string, error) {
	ui := &input.UI{
		Writer: stdout,
		Reader: stdin,
	}

	tfeURL, err := ui.Ask("TFE URL:", &input.Options{
		Default:  "https://app.terraform.io",
		Required: true,
		Loop:     true,
	})

	if err != nil {
		return "", "", err
	}

	tfeAPIToken, err := ui.Ask(fmt.Sprintf("TFE API Token (Create one at %s/app/settings/tokens)", tfeURL), &input.Options{
		Default:     "",
		Required:    true,
		Loop:        true,
		// Mask:        true,
		// MaskDefault: true,
	})

	if err != nil {
		return "", "", err
	}

	return tfeURL, tfeAPIToken, nil
}

// CreateConfigFileFromValues creates a config file with viper using given values
func CreateConfigFileFromValues(url string, token string) {
	viper.Set("tfe_url", url)
	viper.Set("tfe_api_token", token)
	configPath := ConfigPath()
	viper.SetConfigFile(configPath)
	err := viper.WriteConfig()

	if err != nil {
		log.Fatal("Failed to write to: ", configPath, " Error was: ", err)
		os.Exit(1)
	}

	fmt.Println("Saved to", configPath)
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
