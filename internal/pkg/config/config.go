package config

import (
	"fmt"
	"os"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/constants"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Logger      *logger.Config   `mapstructure:"logger"`
	Server      *hxxp.Config     `mapstructure:"server"`
	Postgres    *postgres.Config `mapstructure:"postgres"`
}

func BindAllConfig() (*Config, error) {
	env, err := LoadEnvironment()
	if err != nil {
		return nil, err
	}

	cfg, err := BindKey[*Config]("", env)
	if err != nil {
		return nil, fmt.Errorf("failed to bind config: %w", err)
	}

	cfg.Environment = env

	return cfg, nil
}

func BindLoggerConfig(env string) (*logger.Config, error) {
	return BindKey[*logger.Config]("logger", env)
}

func BindServerConfig(env string) (*hxxp.Config, error) {
	return BindKey[*hxxp.Config]("server", env)
}

func LoadEnvironment() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("could not load .env file: %w", err)
	}

	return os.Getenv(constants.AppEnv), nil
}

func BindKey[T any](key string, env string) (T, error) {
	if env == "" {
		env = constants.EnvDevelopment
	}

	var result T

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(constants.ConfigPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return result, fmt.Errorf("failed to read config: %w", err)
	}

	if key == "" {
		if err := viper.Unmarshal(&result); err != nil {
			return result, fmt.Errorf("failed to unmarshal config: %w", err)
		}
	} else {
		if err := viper.UnmarshalKey(key, &result); err != nil {
			return result, fmt.Errorf("failed to unmarshal config by key: %w", err)
		}
	}

	return result, nil
}
