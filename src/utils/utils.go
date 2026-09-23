package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

func (_ *CommonUtility) SlugifyUnicode(s string) string {
	// Normalize accents (é → e)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}))
	s, _, _ = transform.String(t, s)

	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace spaces with hyphens
	s = strings.ReplaceAll(s, " ", "-")

	// Remove non-alphanumeric characters (except hyphens)
	reg := regexp.MustCompile("[^a-z0-9-]+")
	s = reg.ReplaceAllString(s, "")

	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile("-+")
	s = reg.ReplaceAllString(s, "-")

	// Trim hyphens
	s = strings.Trim(s, "-")

	return s
}
