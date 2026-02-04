package config

// Config holds application configuration
type Config struct {
	Port string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}
