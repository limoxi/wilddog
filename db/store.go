package db

import (
	"github.com/limoxi/ghost"
)

type Store struct {
	ghost.BaseDBModel
	UUID      string     `gorm:"size:32"`
	Biz       string     `gorm:"size:128"`
	Data      ghost.Json `gorm:"type:jsonb"`
	IsDeleted bool
}

func (Store) TableName() string {
	return "store_store"
}

func init() {
	ghost.RegisterDBModel(&Store{})
}
