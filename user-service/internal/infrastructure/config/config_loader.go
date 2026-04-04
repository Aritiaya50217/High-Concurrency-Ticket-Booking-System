package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"app"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	Kafka struct {
		Brokers            []string `yaml:"brokers"`
		TopicBookingCreate string   `yaml:"topic_booking_created"`
	} `yaml:"kafka"`

	JWT struct {
		Secret      string
		ExpireHours int
	}
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) JWTExpireDuration() time.Duration {
	return time.Duration(c.JWT.ExpireHours) * time.Hour
}
