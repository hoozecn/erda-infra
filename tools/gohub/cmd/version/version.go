// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"fmt"

	"github.com/erda-project/erda-infra/tools/gohub/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(versionCmd)
}

// Version .
const Version = "1.0.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of ` + cmd.Name,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
