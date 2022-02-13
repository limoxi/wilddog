package middleware

import (
	"github.com/limoxi/ghost"
	ghost_middleware "github.com/limoxi/ghost/middleware"
)

func init(){
	ghost.RegisterMiddleware(&ghost_middleware.EntryMiddleware{})
}