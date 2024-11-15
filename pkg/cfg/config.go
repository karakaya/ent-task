package cfg

import (
	"github.com/spf13/viper"
)

//TODO: decided to use env from the dockerfile

type Config struct {
	Server          ServerConfig
	Database        DatabaseConfig
	RunningInDocker bool `mapstructure:"RUNNING_IN_DOCKER"`
}

type ServerConfig struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type DatabaseConfig struct {
	Host              string `mapstructure:"DB_HOST"`
	DatabaseName      string `mapstructure:"DB_NAME"`
	User              string `mapstructure:"DB_USER"`
	Port              string `mapstructure:"DB_PORT"`
	Password          string `mapstructure:"DB_PASSWORD"`
	MigrationsEnabled bool   `mapstructure:"MIGRATIONS_ENABLED"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "username")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "entain-task")
	viper.SetDefault("MIGRATIONS_ENABLED", true)
	viper.SetDefault("RUNNING_IN_DOCKER", false)

	config := &Config{
		Server: ServerConfig{
			ServerAddress: viper.GetString("SERVER_ADDRESS"),
		},
		Database: DatabaseConfig{
			Host:              viper.GetString("DB_HOST"),
			Port:              viper.GetString("DB_PORT"),
			User:              viper.GetString("DB_USER"),
			DatabaseName:      viper.GetString("DB_NAME"),
			Password:          viper.GetString("DB_PASSWORD"),
			MigrationsEnabled: viper.GetBool("MIGRATIONS_ENABLED"),
		},
		RunningInDocker: viper.GetBool("RUNNING_IN_DOCKER"),
	}

	return config, nil
}
