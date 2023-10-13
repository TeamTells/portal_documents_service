package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config is a config of service.
type Config struct {
	HTTPAddr   string
	HTTPPort   string
	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

// Read reads config from environment.
func Read() Config {
	var config Config
	_ = godotenv.Load(".env")
	httpAddr, existsAddr := os.LookupEnv("HTTP_ADDR")
	if existsAddr {
		config.HTTPAddr = httpAddr
	}
	httpPort, existsPort := os.LookupEnv("HTTP_PORT")
	if existsPort {
		config.HTTPPort = httpPort
	}
	dBName, existsDBName := os.LookupEnv("DB_NAME")
	if existsDBName {
		config.DBName = dBName
	}
	dBPassword, existsDBPassword := os.LookupEnv("DB_PASSWORD")
	if existsDBPassword {
		config.DBPassword = dBPassword
	}
	dBPort, existsDBPort := os.LookupEnv("DB_PORT")
	if existsDBPort {
		config.DBPort = dBPort
	}
	return config
}

func GetServiceHostPort(cfg Config) string {
	return fmt.Sprintf("%s:%s", cfg.HTTPAddr, cfg.HTTPPort)
}
