package valueobject

import "testing"

func TestCreateSlugFromText(t *testing.T) {
	t.Run("should be able to create a slug from a text", func(t *testing.T) {
		text := "Example of a text-"

		slug := NewSlug("").CreateSlugFromText(text)

		expected := "example-of-a-text"

		if slug.Value() != expected {
			t.Errorf("Expected slug to be %s, but got %s", expected, slug.Value())
		}
	})
}
