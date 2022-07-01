package config

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/structs"
)

const (
	delimeter = "."
	seperator = "__"

	tagName = "koanf"

	upTemplate     = "================ Loaded Configuration ================"
	bottomTemplate = "======================================================"
)

func Load() (*Config, error) {
	k := koanf.New(delimeter)

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load default configuration from environment variables
	if err := LoadEnv(k); err != nil {
		return nil, fmt.Errorf("error loading default values: %v", err)
	}

	config := Config{}
	var tag = koanf.UnmarshalConf{Tag: tagName}
	if err := k.UnmarshalWithConf("", &config, tag); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	// pretty print loaded configuration using provided template
	log.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	return &config, nil
}
