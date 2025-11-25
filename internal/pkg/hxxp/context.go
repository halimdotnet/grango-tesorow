package hxxp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/constants"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Ctx     context.Context
}

func (c *Context) Bind(v interface{}) error {
	//return json.NewDecoder(c.Request.Body).Decode(v)

	reader := http.MaxBytesReader(c.Writer, c.Request.Body, constants.MaxRequestBodyBytes)
	defer reader.Close()

	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	if err := dec.Decode(v); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			return fmt.Errorf("request payload exceeds allowed size")
		}

		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return fmt.Errorf("invalid JSON position %d", syntaxErr.Offset)
		}

		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil

}

func (c *Context) Response(statusCode int, resp Response) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(statusCode)

	if err := json.NewEncoder(c.Writer).Encode(resp); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	return chi.URLParam(c.Request, key)
}

func (c *Context) Header(key string) string {
	return c.Request.Header.Get(key)
}
