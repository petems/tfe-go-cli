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
  "context"

  "net/http"
  "net/url"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"

  tfe "github.com/hashicorp/go-tfe"
)

var validateCmd = &cobra.Command{
  Use:   "validate",
  Short: "Validate your TFE credentials",
  Long:  `Validates your credentials from your
  configuration file (defaults to ~/.tgc.yaml`,
  Run: func(cmd *cobra.Command, args []string) {

    initConfig()

    tfeConfigURL := viper.GetString("tfe_url")
    tfeConfigToken := viper.GetString("tfe_api_token")

    fmt.Println("URL in config:", tfeConfigURL)

    url, err := url.ParseRequestURI(tfeConfigURL)
    if err == nil {
      fmt.Println("URL is valid:", url)
    } else {
      fmt.Println("URL isnt valid:", err)
    }

    response, err := CheckIf200FromURL(tfeConfigURL)

    if err == nil {
      fmt.Println("URL is reachable:", tfeConfigURL, response.StatusCode)
    } else {
      fmt.Println("Error reaching hostname:", err)
    }

    tfeConfig := &tfe.Config{
      Token: tfeConfigToken,
    }

    client, err := tfe.NewClient(tfeConfig)
    if err != nil {
      fmt.Println(err)
    }

    // Create a context
    ctx := context.Background()

    user, err := client.Users.ReadCurrent(ctx)
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println("API Token valid - User is:", user.Username)
    }

  },
}


func CheckIf200FromURL(url string) (*http.Response, error) {
  rsp, err := http.Get(url)
  if err != nil {
    return nil, fmt.Errorf("Failed at 'get' stage. Error was: %s", err)
  }

  if rsp.StatusCode != 200 {
    return rsp, fmt.Errorf("HTTP request error. Response code: %d", rsp.StatusCode)
  }

  return rsp, nil
}

func init() {
  rootCmd.AddCommand(validateCmd)
}
