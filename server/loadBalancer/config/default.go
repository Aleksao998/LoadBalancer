package config

var Config = &GlobalConfig{
	Database: DatabaseConfig{
		Host: "localhost",
		Port: "5432",
		Db:   "loadBalancer",
		User: "postgres",
		Pwd:  "opacicaleksa32",
	},
	Web: WebConfig{
		Port: "5000",
	},
	JWT: JWTConfig{
		Secret: "someSecret",
	},
}
