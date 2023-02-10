package config

// Config ...
type Config struct {
	BindAddr      string `toml:"bind_addr"`
	LogLevel      string `toml:"log_level"`
	DatabaseURL   string `toml:"database_url"`
	NatsClusterId string `toml:"nats_cluster_id"`
	NatsHostname  string `toml:"nats_hostname"`
	DbDriverName  string `toml:"db_driver_name"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
