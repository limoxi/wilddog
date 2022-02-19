package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limoxi/ghost"
	"strings"
	m_db "wilddog/db"
)

type AuthService struct {
	ghost.DomainService
}

func (this *AuthService) MakeUUID() string {
	str := strings.ReplaceAll(uuid.New().String(), "-", "")

	db := this.GetDB()
	var count int64
	db.Model(&m_db.Auth{}).Where(ghost.Map{
		"uuid": str,
	}).Count(&count)
	if count > 0 {
		return this.MakeUUID()
	}

	db.Create(&m_db.Auth{
		UUID: str,
	})
	return str
}

func NewAuthService(ctx context.Context) *AuthService {
	inst := new(AuthService)
	inst.SetCtx(ctx)
	return inst
}

func GetUUIDFromContext(ctx context.Context) string {
	uuid, ok := ctx.(*gin.Context).Get("uuid")
	if !ok || uuid == "" {
		panic(ghost.NewBusinessError("非法请求"))
	}
	return uuid.(string)
}
