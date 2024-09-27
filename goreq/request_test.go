package goreq

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	env := map[string]string{
		"API_KEY": "abc123",
	}

	r, err := ParseFile("testdata/env.lua", env)
	assert.NoError(t, err)

	headers := r.Headers
	assert.Equal(t, "Bearer abc123", headers["Authorization"])
}

func TestExecute(t *testing.T) {
	r, err := ParseFile("testdata/get.lua", nil)
	assert.NoError(t, err)

	assert.Equal(t, http.MethodGet, r.Method)
	assert.Equal(t, "https://example.com/", r.URL)

	headers := r.Headers
	assert.Equal(t, "application/json", headers["Content-Type"])
}

func TestDoPostJson(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		data := make(map[string]interface{})
		err = json.Unmarshal(body, &data)
		assert.NoError(t, err)

		expected := map[string]interface{}{
			"hello": "world",
		}

		assert.Equal(t, expected, data)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	// Update the URL in the test data to use the test server's URL
	r, err := ParseFile("testdata/post_json.lua", nil)
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

func TestDoPost(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		assert.Equal(t, "hello world", string(body))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	// Update the URL in the test data to use the test server's URL
	r, err := ParseFile("testdata/post_raw_body.lua", nil)
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

func TestDo(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	// Update the URL in the test data to use the test server's URL
	r, err := ParseFile("testdata/get.lua", nil)
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
