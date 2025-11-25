package hxxp

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	BasePath string `mapstructure:"base-path"`
}

type Response struct {
	Error   bool                   `json:"error"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data,omitempty"`
	Metas   map[string]interface{} `json:"metadata,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}
