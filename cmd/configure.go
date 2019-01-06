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
	"log"
	"os"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	input "github.com/tcnksm/go-input"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your TFE credentials",
	Long:  `Prompts for your TFE API credentials, then writes them to
	a configuration file (defaults to ~/.tgc.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateConfigFileFromPrompts(os.Stdin, os.Stdout)
	},
}

func CreateConfigFileFromPrompts(stdin io.Reader, stdout io.Writer) {
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
		log.Fatal(err)
	}
	viper.Set("tfe_url", tfeURL)

	tfeAPIToken, err := ui.Ask(fmt.Sprintf("TFE API Token (Create one at %s/app/settings/tokens)", tfeURL), &input.Options{
		Default:  	 "",
		Required: 	 true,
		Loop:     	 true,
		Mask:        true,
		MaskDefault: true,
		})

	if err != nil {
		log.Fatal(err)
	}
	viper.Set("tfe_api_token", tfeAPIToken)

	configPath := ConfigPath()
	viper.SetConfigFile(configPath)

	err = viper.WriteConfig()

	if err != nil {
		log.Fatal("Failed to write to: ", configPath, " Error was: ", err)
	}

	fmt.Println("Saved to", configPath)
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
