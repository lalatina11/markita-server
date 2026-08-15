package config

import "fmt"

type SupabaseConfig struct {
	DatabaseURL    string
	ProjectURL     string
	AuthURL        string
	StorageURL     string
	PublishableKey string
	SecretKey      string
	AnonKey        string
}

func NewSupabaseConfig() *SupabaseConfig {
	DatabaseURL := GetEnv("DATABASE_URL")
	ProjectURL := GetEnv("SUPABASE_PROJECT_URL")
	AuthURL := fmt.Sprintf("%s/auth/v1", ProjectURL)
	StorageURL := fmt.Sprintf("%s/storage/v1", ProjectURL)
	PublishableKey := GetEnv("SUPABASE_PUBLISHABLE_KEY")
	SecretKey := GetEnv("SUPABASE_SECRET_KEY")
	AnonKey := GetEnv("SUPABASE_ANON_KEY")

	return &SupabaseConfig{DatabaseURL, ProjectURL, AuthURL, StorageURL, PublishableKey, SecretKey, AnonKey}
}
