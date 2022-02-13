package db

import "github.com/limoxi/ghost"

type Auth struct {
	ghost.BaseDBModel
	UUID string `gorm:"size:32;unique"`
}

func (Auth) TableName() string {
	return "auth_uuid"
}

func init() {
	ghost.RegisterDBModel(&Auth{})
}
