package gorec

import (
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

type Request struct {
	Method string
	URL    string
}

func ParseFile(path string) (Request, error) {
	lState := lua.NewState()
	defer lState.Close()

	var r Request

	lState.SetGlobal("get", lState.NewFunction(get(&r)))
	if err := lState.DoFile(path); err != nil {
		return r, err
	}

	return r, nil
}

func Do(r Request) (*http.Response, error) {
	// Convert the Request to an http.Request
	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return nil, err
	}

	// Do the request
	return http.DefaultClient.Do(req)
}

func get(r *Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		r.Method = http.MethodGet

		// Get the first argument as a table
		// TODO: Error handling
		t := L.ToTable(1)

		// Get the URl
		url := t.RawGetString("url")
		r.URL = url.String()

		return 0
	}
}
