package config

type AppConfig struct {
	Env           string
	AvatarBaseURL string
	Server        *ServerConfig
	Supabase      *SupabaseConfig
}

func NewAppConfig() *AppConfig {
	Env := GetEnv("ENV")
	AvatarGeneratorBaseURL := GetEnv("AVATAR_GENERATOR_BASE_URL")
	Server := NewServerConfig()
	Supabase := NewSupabaseConfig()
	return &AppConfig{Env, AvatarGeneratorBaseURL, Server, Supabase}
}
