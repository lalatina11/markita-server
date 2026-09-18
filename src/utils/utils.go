package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
)

type CommonUtility struct {
	SupabaseConfig *config.SupabaseConfig
}

func NewCommonUtiliity() *CommonUtility {
	SupabaseConfig := config.NewSupabaseConfig()
	return &CommonUtility{SupabaseConfig}
}

func (this *CommonUtility) GenerateMediaURL(mediaURL string) string {
	if strings.HasPrefix(mediaURL, "http") {
		return mediaURL
	}
	return fmt.Sprintf("%s/object/public/%s", this.SupabaseConfig.StorageURL, mediaURL)
}

func (this *CommonUtility) GenerateFileName() string {
	now := time.Now().UTC().String()
	name := uuid.NewString()

	return strings.ReplaceAll(fmt.Sprintf("%s-%s", now, name), " ", "")
}
