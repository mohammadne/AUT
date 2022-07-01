package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/mohammadne/university/cryptography/internal"
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

func LoadEnv(k *koanf.Koanf) error {
	var prefix = strings.ToUpper(internal.Subsystem) + seperator

	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, prefix))
		return strings.ReplaceAll(base, seperator, delimeter)
	}

	// load environment variables
	if err := k.Load(env.Provider(prefix, delimeter, callback), nil); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}
