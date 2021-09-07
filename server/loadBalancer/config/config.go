package config

type GlobalConfig struct {
	Database DatabaseConfig
	Web      WebConfig
	JWT      JWTConfig
}

type WebConfig struct {
	Port string
}

type JWTConfig struct {
	Secret string
}

type DatabaseConfig struct {
	Host string
	Port string
	Db   string
	User string
	Pwd  string
}
