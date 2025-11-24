package hxxp

import "encoding/json"

func (c *Context) BindJSON(v interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	return c.Param(key)
}

func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}
