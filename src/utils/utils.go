package utils

import (
	"fmt"
	"strings"

	"github.com/lalatina11/markita.git/src/config"
)

type CommonUtility struct {
	SupabaseConfig *config.SupabaseConfig
}

func NewCommonUtiliity() *CommonUtility {
	Config := config.NewSupabaseConfig()
	return &CommonUtility{Config}
}

func (this *CommonUtility) GenerateMediaURL(mediaURL string) string {
	if strings.HasPrefix(mediaURL, "http") {
		return mediaURL
	}
	return fmt.Sprintf("%s/object/public/%s", this.SupabaseConfig.StorageURL, mediaURL)
}
