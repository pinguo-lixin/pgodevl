/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// PgoConfig pgo application config
type PgoConfig struct {
	SourcePath     string
	ControllerPath string
}

var (
	path      string
	pgoConfig *PgoConfig
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "pgodevl",
	Short: "Pgo develop tools",
	Long:  `PGO 开发工具包`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("root called")
	// },
}

func init() {
	path = *rootCmd.PersistentFlags().String("path", "", "specify pgo project path, default is current directory")

	pgoConfig = new(PgoConfig)
	if path == "" {
		pwd, _ := os.Getwd()
		pgoConfig.SourcePath = pwd
	} else {
		pgoConfig.SourcePath, _ = filepath.Abs(path)
	}
}

// Execute 命令入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
