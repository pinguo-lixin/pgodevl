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
