package gorec

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func loadExample(t *testing.T, name string) Request {
	d, err := os.ReadFile("testdata/" + name)
	assert.NoError(t, err)

	r := Request{}
	err = yaml.Unmarshal(d, &r)
	assert.NoError(t, err)

	return r
}

func TestRequest(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		r := loadExample(t, "simple.yaml")
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "https://example.com/", r.URL)
	})

	t.Run("headers", func(t *testing.T) {
		r := loadExample(t, "headers.yaml")
		assert.Equal(t, map[string]string{
			"Authorization": "Basic 123",
			"Content-Type":  "application/json",
		}, r.Headers)
	})

	t.Run("body", func(t *testing.T) {
		r := loadExample(t, "body.yaml")
		assert.Equal(t, "{\n  \"foo\": \"bar\"\n}\n", r.Body)
	})
}
