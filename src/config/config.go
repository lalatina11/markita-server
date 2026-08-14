package config

type AppConfig struct {
	Env    string
	Server *ServerConfig
}

func NewAppConfig() *AppConfig {
	Env := GetEnv("ENV")
	Server := NewServerConfig()
	return &AppConfig{Env, Server}
}
