package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Metrics  MetricsConfig
	Logger   LoggerConfig
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
	SSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

type ServerConfig struct {
	GRPCPort     int    `envconfig:"GRPC_PORT" default:"50051"`
	HTTPPort     int    `envconfig:"HTTP_PORT" default:"8080"`
	GRPCAddress  string `envconfig:"GRPC_ADDRESS" required:"true"`
	HTTPPAddress string `envconfig:"HTTP_ADDRESS" required:"true"`
}

type MetricsConfig struct {
	PrometheusPort int `envconfig:"PROMETHEUS_PORT" default:"9090"`
}

type LoggerConfig struct {
	LogLevel      string `envconfig:"LOG_LEVEL" default:"debug"`
	LogJsonFormat bool   `envconfig:"LOG_JSON_FORMAT" default:"false"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки конфигурации: %w", err)
	}
	return &cfg, nil
}
