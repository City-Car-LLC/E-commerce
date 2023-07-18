package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const DefaultPath = "config.json"

type Config struct {
	AppHost     string `json:"app_host"`
	AppPort     string `json:"app_port"`
	PostgresDSN string `json:"postgres_dsn"`
}

func (c *Config) DefaultIfNotSet() {
	c.AppHost = set(c.AppHost, "localhost")
	c.AppPort = set(c.AppPort, "5432")
	c.PostgresDSN = set(c.PostgresDSN,
		"user=postgres"+
			"password=mspz3jic"+
			"host=localhost "+
			"port=5432 "+
			"dbname=commerce "+
			"sslmode=disable "+
			"search_path=public "+
			"TimeZone=Asia/Dushanbe")
}

func set(s, v string) string {
	if s == "" {
		return v
	}
	return s
}

func NewConfig(filepath string) (*Config, error) {

	folder := "./files/"
	if _, err := os.Stat(folder); err != nil {
		os.Mkdir(folder, 0755)
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filepath, err)
	}
	var c *Config
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}
	return c, nil
}
