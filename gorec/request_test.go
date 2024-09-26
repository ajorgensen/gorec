package gorec

import (
	"io"
	"net/http"
	"net/http/httptest"
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

func TestExecute(t *testing.T) {
	r, err := ParseFile("testdata/get.lua")
	assert.NoError(t, err)

	assert.Equal(t, http.MethodGet, r.Method)
	assert.Equal(t, "https://example.com/", r.URL)
}

func TestDo(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	// Update the URL in the test data to use the test server's URL
	r, err := ParseFile("testdata/get.lua")
	assert.NoError(t, err)

	r.URL = ts.URL // Use the test server's URL

	// Execute the request
	resp, err := Do(r)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", string(body))
}
