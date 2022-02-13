package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
)

type AuthMiddleware struct{}

func (this *AuthMiddleware) Init() {
	ghost.Info("AuthMiddleware loaded")
}

func (this *AuthMiddleware) PreRequest(ctx *gin.Context) {
}

func (this *AuthMiddleware) AfterResponse(ctx *gin.Context) {

}

func init() {
	ghost.RegisterMiddleware(&AuthMiddleware{})
}
