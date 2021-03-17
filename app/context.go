package app

import (
	"context"
	"time"
)

type Context struct {
	Context         context.Context
	App             App
	HandlerName     string
	ResponseHeaders map[string]string
}

func NewContext(app App) *Context {
	return &Context{
		App:             app,
		Context:         context.Background(),
		ResponseHeaders: make(map[string]string),
	}
}

func (c *Context) AddHandlerName(handlerName string) error {
	c.HandlerName = handlerName
	return nil
}

func (c *Context) AddValue(key interface{}, value interface{}) error {
	c.Context = context.WithValue(c.Context, key, value)
	return nil
}

func (c *Context) GetValue(key interface{}) interface{} {
	return c.Value(key)
}

func (c *Context) GetHandlerName() string {
	return c.HandlerName
}

func (c *Context) AddRespHeader(key, value string) {
	c.ResponseHeaders[key] = value
}

//////////////////////////// Context Interface ////////////////////////////////////
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *Context) Err() error {
	return c.Context.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}

////////////////////////////////////////////////////////////////////////////////
