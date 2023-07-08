package config

type Config struct {
	AppHost     string `json:"app_host"`
	AppPort     string `json:"app_port"`
	PostgresDSN string `json:"postgres_dsn"`
}
