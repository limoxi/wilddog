package store

import (
	"github.com/limoxi/ghost"
)

type EncodedStoredData struct {
	Id        int       `json:"id"`
	Biz       string    `json:"biz"`
	Data      ghost.Map `json:"data"`
	UpdatedAt string    `json:"updated_at"`
}
