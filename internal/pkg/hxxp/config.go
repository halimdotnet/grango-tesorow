package hxxp

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	BasePath string `mapstructure:"base-path"`
}
