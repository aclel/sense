package middleware

import (
	"net/http"

	"github.com/aclel/sense/store"

	"golang.org/x/net/context"
)

type Context struct {
	db *store.DB
}

func (c *Context) Handler(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		context.WithValue(ctx, "db", c.db)
	}
	return http.HandlerFunc(f)
}
