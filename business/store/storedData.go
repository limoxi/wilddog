package store

import (
	"context"
	"github.com/limoxi/ghost"
	"time"
	m_db "wilddog/db"
)

type StoredData struct {
	ghost.DomainModel

	Id        int        `json:"id"`
	UUID      string     `json:"uuid"`
	Biz       string     `json:"biz"`
	Data      ghost.Json `json:"data"`
	IsDeleted bool       `json:"is_deleted"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

func NewStoredDataFromDbModel(ctx context.Context, dbModel *m_db.Store) *StoredData {
	inst := new(StoredData)
	inst.SetCtx(ctx)
	inst.Set("dbModel", dbModel)

	inst.Id = dbModel.Id
	inst.UUID = dbModel.UUID
	inst.Biz = dbModel.Biz
	inst.Data = dbModel.Data
	inst.IsDeleted = dbModel.IsDeleted
	inst.UpdatedAt = dbModel.UpdatedAt
	inst.CreatedAt = dbModel.CreatedAt

	return inst
}
