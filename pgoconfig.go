package main

import (
	"os"
	"path/filepath"
)

// PgoConfig pgo application config
type PgoConfig struct {
	SourcePath     string
	ControllerPath string
}

func defaultPgoConfig() *PgoConfig {
	pwd, _ := os.Getwd()
	conf := &PgoConfig{
		SourcePath:     filepath.Join(pwd, "pkg"),
		ControllerPath: filepath.Join(pwd, "pkg", "controller"),
	}
	return conf
}
