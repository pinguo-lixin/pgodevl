package main

// PgoConfig pgo application config
type PgoConfig struct {
	SourcePath     string
	ControllerPath string
}

func defaultPgoConfig() *PgoConfig {
	conf := &PgoConfig{
		// SourcePath: ,
	}
	return conf
}
