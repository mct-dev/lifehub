package config

import "github.com/spf13/viper"

// EnvPrefix the prefix for environment variables
const EnvPrefix = "LH"

// Config contains config for backend
type Config struct {
	Port int    // Default: 8080
	Env  string // Default: production
}

// InitConfig from config file / env / default values
func InitConfig() *Config {
	var config Config

	viper.AutomaticEnv()
	viper.SetEnvPrefix(EnvPrefix)

	viper.SetDefault("Port", 8182)
	viper.SetDefault("Env", "production")

	_ = viper.Unmarshal(&config)

	return &config
}
