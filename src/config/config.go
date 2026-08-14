package config

type AppConfig struct {
	Env      string
	Server   *ServerConfig
	Supabase *SupabaseConfig
}

func NewAppConfig() *AppConfig {
	Env := GetEnv("ENV")
	Server := NewServerConfig()
	Supabase := NewSupabaseConfig()
	return &AppConfig{Env, Server, Supabase}
}
