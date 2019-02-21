// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"

	"github.com/sorintlab/agola/internal/services/gateway/api"
	"github.com/spf13/cobra"
)

var cmdProjectList = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		if err := projectList(cmd, args); err != nil {
			log.Fatalf("err: %v", err)
		}
	},
	Short: "list",
}

type projectListOptions struct {
	limit int
	start string
}

var projectListOpts projectListOptions

func init() {
	flags := cmdProjectList.PersistentFlags()

	flags.IntVar(&projectListOpts.limit, "limit", 10, "max number of runs to show")
	flags.StringVar(&projectListOpts.start, "start", "", "starting project name (excluded) to fetch")

	cmdProject.AddCommand(cmdProjectList)
}

func printProjects(projectsResponse *api.GetProjectsResponse) {
	for _, project := range projectsResponse.Projects {
		fmt.Printf("%s: Name: %s\n", project.ID, project.Name)
	}
}

func projectList(cmd *cobra.Command, args []string) error {
	gwclient := api.NewClient(gatewayURL, token)

	projectsResponse, _, err := gwclient.GetProjects(context.TODO(), projectListOpts.start, projectListOpts.limit, false)
	if err != nil {
		return err
	}

	printProjects(projectsResponse)

	return nil
}
