/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/piyush1146115/Go-REST-API-with-CLI/api"
	"github.com/spf13/cobra"
)

var port string

// startServerCmd represents the startServer command
var startServerCmd = &cobra.Command{
	Use:   "startServer",
	Short: "This will start my server",
	Long:  `This will start an REST API server that can serve some CRUD request and take port number as a flag to run on a specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("startServer called")
		api.CreateServer(port)
		api.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
	startServerCmd.PersistentFlags().StringVarP(&port, "port", "p", "10000", "This flag will set the post")
}
