package cfg

import "github.com/spf13/viper"

//TODO: decided to use env from the dockerfile

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	ServerAddress string
}

type DatabaseConfig struct {
	Host              string
	DatabaseName      string
	User              string
	Name              string
	Port              string
	Password          string
	MigrationsEnabled string
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "username")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "entain-task")

	config := &Config{
		Server: ServerConfig{
			ServerAddress: viper.GetString("SERVER_ADDRESS"),
		},

		Database: DatabaseConfig{
			Host:         viper.GetString("DB_HOST"),
			Port:         viper.GetString("DB_PORT"),
			User:         viper.GetString("DB_USER"),
			DatabaseName: viper.GetString("DB_NAME"),
			Password:     viper.GetString("DB_PASSWORD"),
			Name:         viper.GetString("DB_NAME"),
		},
	}

	return config, nil
}
