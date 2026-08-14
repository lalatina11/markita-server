package config

type ServerConfig struct {
	Host string
	Port string
}

func NewServerConfig() *ServerConfig {
	Host := GetEnv("HOST")
	Port := GetEnv("PORT")
	return &ServerConfig{Host, Port}
}
