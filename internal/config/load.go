package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
)

const (
	delimeter = "."
	seperator = "__"

	prefix = "ZAR" + seperator
)

func Load() *Config {
	k := koanf.New(delimeter)

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	if err := loadEnv(k); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	config := Config{}
	if err := k.Unmarshal("", &config); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	if config.Print {
		var (
			upTemplate     = "================ Loaded Configuration ================"
			bottomTemplate = "======================================================"
		)

		// pretty print loaded configuration using provided template
		log.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config
}

// load from environment variables
func loadEnv(k *koanf.Koanf) error {
	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, prefix))
		return strings.ReplaceAll(base, seperator, delimeter)
	}

	if err := k.Load(env.Provider(prefix, delimeter, callback), nil); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}
