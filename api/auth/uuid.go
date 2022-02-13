package auth

import (
	"github.com/limoxi/ghost"
	"wilddog/business/service"
)

type Uuid struct {
	ghost.ApiTemplate
}

func (this *Uuid) Resource() string {
	return "auth.uuid"
}

func (this *Uuid) Put() ghost.Response {
	id := service.NewAuthService(this.GetCtx()).MakeUUID()
	return ghost.NewJsonResponse(id)
}

func init() {
	ghost.RegisterApi(&Uuid{})
}
