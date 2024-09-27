package goreq

import (
	"encoding/json"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type GQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// TODO: Error handling
func gql(r *Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		// First argument is the query
		query := L.ToString(1)

		// Second argument is the variables
		var variables map[string]interface{}
		if L.Get(2) != lua.LNil {
			variables = make(map[string]interface{})
			L.CheckTypes(2, lua.LTTable)
			t := L.ToTable(2)
			t.ForEach(func(k lua.LValue, v lua.LValue) {
				variables[k.String()] = v.String()
			})
		}

		// create the request
		gqlRequest := GQLRequest{
			Query:     query,
			Variables: variables,
		}

		jsonBody, err := json.Marshal(gqlRequest)
		if err != nil {
			fmt.Println("Error marshalling request body:", err)
			return 0
		}

		r.Headers["Content-Type"] = "application/json"
		r.Body = jsonBody
		return 0
	}
}
