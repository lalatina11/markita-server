package config

type SupabaseConfig struct {
	DatabaseURL    string
	ProjectURL     string
	PublishableKey string
	SecretKey      string
	AnonKey        string
}

func NewSupabaseConfig() *SupabaseConfig {
	DatabaseURL := GetEnv("DATABASE_URL")
	ProjectURL := GetEnv("SUPABASE_PROJECT_URL")
	PublishableKey := GetEnv("SUPABASE_PUBLISHABLE_KEY")
	SecretKey := GetEnv("SUPABASE_SECRET_KEY")
	AnonKey := GetEnv("SUPABASE_ANON_KEY")

	return &SupabaseConfig{DatabaseURL, ProjectURL, PublishableKey, SecretKey, AnonKey}
}
