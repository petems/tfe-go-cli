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

  "github.com/blang/semver"
  "github.com/spf13/cobra"
)

const (
  progMajor        = 0
  progMinor        = 1
  progPatch        = 0
  progReleaseLevel = "alpha"
  progReleaseNum   = 1
)

var (
  gitCommit string

  progVersion = semver.Version{
    Major: progMajor,
    Minor: progMinor,
    Patch: progPatch,
    Pre: []semver.PRVersion{{
      VersionStr: progReleaseLevel,
    }, {
      VersionNum: progReleaseNum,
      IsNum:      true,
    }},
  }
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Return tfe-go-cli version",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(progVersion)
  },
}

func init() {
  rootCmd.AddCommand(versionCmd)
}
