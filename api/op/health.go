package op

import (
	"github.com/limoxi/ghost"
)

type Health struct {
	ghost.ApiTemplate
}

func (this *Health) Resource() string {
	return "op.health"
}

func (this *Health) Get() ghost.Response {
	return ghost.NewJsonResponse(ghost.Map{
		"status": "healthy",
	})
}

func init() {
	ghost.RegisterApi(&Health{})
}
