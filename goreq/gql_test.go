package goreq

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGQL(t *testing.T) {
	r, err := ParseFile("testdata/gql.lua", nil)
	assert.NoError(t, err)

	assert.Equal(t, http.MethodPost, r.Method)
	assert.Equal(t, "https://example.com/graphql", r.URL)

	headers := r.Headers
	assert.Equal(t, "application/json", headers["Content-Type"])

	body := r.Body

	var req GQLRequest
	err = json.Unmarshal(body, &req)
	require.NoError(t, err)
}
