package cfg

type Config struct {
	Database DB `yaml:"db"`
}

type DB struct {
	Url               string `yaml:"url"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	MigrationsEnabled string `yaml:"migrationsenabled"`
}
