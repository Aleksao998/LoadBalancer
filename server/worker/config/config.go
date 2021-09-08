package config

type GlobalConfig struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host string
	Port string
	Db   string
	User string
	Pwd  string
}
