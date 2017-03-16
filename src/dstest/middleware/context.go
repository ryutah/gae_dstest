package middleware

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

var ctx *Context

func NewContext() *Context {
	ctx = new(Context)
	return ctx
}

type Context struct {
	ctx context.Context
}

func (c *Context) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	cx := appengine.NewContext(r)
	c.ctx = cx
	next(w, r)
}

func Ctx() context.Context {
	return ctx.ctx
}
