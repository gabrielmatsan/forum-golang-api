package valueobject

import (
	"regexp"
	"strings"
)

type Slug struct {
	value string
}

func NewSlug(value string) *Slug {
	return &Slug{
		value: value,
	}
}

func (s *Slug) CreateSlugFromText(text string) *Slug {
	normalizedText := strings.ToLower(text)
	normalizedText = strings.ReplaceAll(normalizedText, " ", "-")

	reg := regexp.MustCompile(`[^\w-]+`)
	normalizedText = reg.ReplaceAllString(normalizedText, "")

	normalizedText = strings.ReplaceAll(normalizedText, "--", "-")

	normalizedText = strings.TrimSuffix(normalizedText, "-")

	return NewSlug(normalizedText)
}

func (s *Slug) Value() string {
	return s.value
}
