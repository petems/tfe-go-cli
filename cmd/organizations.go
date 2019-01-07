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
  "strings"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"

  tfe "github.com/hashicorp/go-tfe"
)

var organizationsCmd = &cobra.Command{
  Use:   "organizations",
  Short: "List your organizations",
  Long:  `Lists the organizations for your user`,
  Run: func(cmd *cobra.Command, args []string) {

    initConfig()

    tfeConfigURL := viper.GetString("tfe_url")
    tfeConfigToken := viper.GetString("tfe_api_token")

    orgList, err := ListOrganizations(tfeConfigURL, tfeConfigToken)

    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println(orgList)
    }
  },
}


func ListOrganizations(url string, token string) (string, error) {
  tfeConfig := &tfe.Config{
    Token: token,
    Address: url,
  }

  client, err := tfe.NewClient(tfeConfig)
  if err != nil {
    return "", err
  }

    // Create a context
  ctx := context.Background()

  orgl, err := client.Organizations.List(ctx, tfe.OrganizationListOptions{})

  if err != nil {
    return "", err
  }

  s := []string{}
  for _, org := range orgl.Items {
    s = append(s, fmt.Sprintf("Orgs are: %s", org.Name))
  }
  return strings.Join(s, "\n"), nil
}

func init() {
  rootCmd.AddCommand(organizationsCmd)
}
