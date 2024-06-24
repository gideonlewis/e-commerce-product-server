package config

import (
	"github.com/spf13/viper"
)

var (
	Api      = new(APIConfig)
	Server   = new(ServerConfig)
	Smtp     = new(SmtpConfig)
	Postgres = new(PostgresConfig)
	Redis    = new(RedisConfig)
)

type Config struct {
	Server   ServerConfig
	Api      APIConfig
	Smtp     SmtpConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

func LoadConfig(filename, path string) error {
	var config Config
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Set default values for the configurations
	// viper.SetDefault("app.name", "DefaultApp")
	viper.SetDefault("app.version", "0.0.1")

	err := viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	// Set default values for global
	Api = &config.Api
	Server = &config.Server
	Postgres = &config.Postgres
	Smtp = &config.Smtp
	Redis = &config.Redis

	return nil
}

type APIConfig struct {
	JWTSecret string
	APIKey    string
	StripeKey string
}

type ServerConfig struct {
	Port    int
	Host    string
	CmsPort int
	CmsHost string
}

type SmtpConfig struct {
	Username      string
	Pass          string
	Endpoint      string
	SenderAddress string
}

type PostgresConfig struct {
	Name      string
	Host      string
	Port      int
	User      string
	Pass      string
	EnableSLL bool
}

type RedisConfig struct {
	Address  string
	Password string
}
