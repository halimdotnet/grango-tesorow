package hxxp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error    bool                   `json:"error"`
	Message  string                 `json:"message"`
	Data     interface{}            `json:"data,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Details  map[string]interface{} `json:"details,omitempty"`
}

func (c *Context) SendJSON(statusCode int, resp Response) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.status = statusCode
	c.Writer.WriteHeader(statusCode)

	if err := json.NewEncoder(c.Writer).Encode(resp); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
