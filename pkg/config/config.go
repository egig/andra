package config

import (
	"github.com/spf13/viper"
	"os"
)

const KeyEnvPrefix = "dt"
const KeyPort = "port"
const KeyDir = "dir"

var cfg *viper.Viper
var Defaults map[string]string

func init() {
	cfg = viper.New()
	cfg.SetEnvPrefix(KeyEnvPrefix)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg.SetDefault(KeyDir, wd)
	cfg.SetConfigName("andra")
	contentDir := cfg.GetString(KeyDir)
	cfg.AddConfigPath(contentDir)

	err = viper.ReadInConfig()
	if err != nil {
		// Config is optional
		//panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Get(key string) string {
	return cfg.GetString(key)
}
