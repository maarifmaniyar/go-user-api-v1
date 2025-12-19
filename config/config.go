package config

type Config struct {
	DBUrl string
	Port  string
}

func Load() Config {
	return Config{
		DBUrl: "postgres://postgres:66gxmjrxbe@localhost:5432/postgres?sslmode=disable",
		Port:  "8080",
	}
}
