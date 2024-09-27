package goreq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
}

func ParseFile(path string, environment map[string]string) (Request, error) {
	lState := lua.NewState()
	defer lState.Close()

	var r Request
	r.Headers = make(map[string]string)

	lState.SetGlobal("get", lState.NewFunction(get(&r)))
	lState.SetGlobal("post", lState.NewFunction(post(&r)))

	lState.SetGlobal("headers", lState.NewFunction(headers(&r)))
	lState.SetGlobal("body", lState.NewFunction(body(&r)))
	lState.SetGlobal("env", lState.NewFunction(env(environment)))

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

	// Add the headers to the request
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// Add the body to the request
	if len(r.Body) > 0 {
		req.Body = io.NopCloser(bytes.NewReader(r.Body))
	}

	// Do the request
	return http.DefaultClient.Do(req)
}

func env(environment map[string]string) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		if environment == nil {
			return 0
		}

		// Get the first argument as a string
		key := L.ToString(1)

		// Look up the environemnt value
		value, ok := environment[key]
		if !ok {
			return 0
		}

		v := lua.LString(value)

		L.Push(v)

		return 1
	}
}

func headers(r *Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		// Get the first argument as a table
		// TODO: Error handling
		t := L.ToTable(1)

		t.ForEach(func(k lua.LValue, v lua.LValue) {
			r.Headers[k.String()] = v.String()
		})

		return 0
	}
}

func body(r *Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.CheckTypes(1, lua.LTString, lua.LTTable)

		switch L.Get(1).Type() {
		case lua.LTString:
			r.Body = []byte(L.ToString(1))
			return 0
		case lua.LTTable:
			var err error

			// encode the table as JSON
			t := L.ToTable(1)

			data := make(map[string]interface{})
			t.ForEach(func(k lua.LValue, v lua.LValue) {
				data[k.String()] = v.String()
			})

			r.Body, err = json.Marshal(data)
			if err != nil {
				return 1
			}

			return 0
		}

		return 0
	}
}

func post(r *Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		r.Method = http.MethodPost

		// Get the first argument as a table
		// TODO: Error handling
		t := L.ToTable(1)

		// Get the URl
		// TODO: Check if the url is LNil
		url := t.RawGetString("url")
		r.URL = url.String()

		return 0
	}
}

func get(r *Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		r.Method = http.MethodGet

		// Get the first argument as a table
		// TODO: Error handling
		t := L.ToTable(1)

		// Get the URl
		// TODO: Check if the url is LNil
		url := t.RawGetString("url")
		r.URL = url.String()

		return 0
	}
}
