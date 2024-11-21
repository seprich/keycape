package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"strings"

	"github.com/samber/lo"
)

var Config *koanf.Koanf = loadConfig()

func loadConfig() *koanf.Koanf {
	config := koanf.New(".")
	lo.Must0(config.Load(file.Provider("config.yaml"), yaml.Parser()))
	lo.Must0(config.Load(env.Provider("", ".", envVarNameConverter), nil))
	return config
}

func envVarNameConverter(s string) string {
	return strings.Replace(s, "__", ".", -1)
}
