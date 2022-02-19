package service

import (
	"context"
	"github.com/limoxi/ghost"
	"gorm.io/gorm"
	"time"
	m_db "wilddog/db"
)

type StoreService struct {
	ghost.DomainService
}

func (this *StoreService) AddMany(biz string, datas []ghost.Json, replace bool) {
	uuid := GetUUIDFromContext(this.GetCtx())
	db := this.GetDB()
	if replace {
		this.DeleteByBiz(biz)
	}

	creatingList := make([]*m_db.Store, 0, len(datas))
	for _, data := range datas {
		creatingList = append(creatingList, &m_db.Store{
			UUID: uuid,
			Biz:  biz,
			Data: data,
		})
	}
	result := db.CreateInBatches(creatingList, len(creatingList))
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError(err.Error(), "存储失败"))
	}
}

func (this *StoreService) Add(biz string, data ghost.Json) {
	uuid := GetUUIDFromContext(this.GetCtx())
	db := this.GetDB()
	qs := db.Model(&m_db.Store{}).Where(ghost.Map{
		"uuid": uuid,
		"biz":  biz,
	})
	var count int64
	qs.Count(&count)
	var result *gorm.DB
	if count > 0 {
		result = qs.Updates(ghost.Map{
			"data": data,
		})
	} else {
		result = db.Create(&m_db.Store{
			UUID: uuid,
			Biz:  biz,
			Data: data,
		})
	}
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError(err.Error(), "存储失败"))
	}
}

func (this *StoreService) Modify(biz string, dataId int, data ghost.Json) {
	uuid := GetUUIDFromContext(this.GetCtx())
	result := this.GetDB().Model(&m_db.Store{}).Where(ghost.Map{
		"uuid": uuid,
		"biz":  biz,
		"id":   dataId,
	}).Updates(ghost.Map{
		"data":       data,
		"updated_at": time.Now(),
	})
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError("修改数据失败"))
	}
}

func (this *StoreService) Delete(biz string, dataId int) {
	uuid := GetUUIDFromContext(this.GetCtx())
	result := this.GetDB().Where(ghost.Map{
		"uuid": uuid,
		"biz":  biz,
		"id":   dataId,
	}).Updates(ghost.Map{
		"is_deleted": true,
		"updated_at": time.Now(),
	})
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError("删除数据失败"))
	}
}

func (this *StoreService) DeleteByBiz(biz string) {
	uuid := GetUUIDFromContext(this.GetCtx())
	result := this.GetDB().Where(ghost.Map{
		"uuid": uuid,
		"biz":  biz,
	}).Updates(ghost.Map{
		"is_deleted": true,
		"updated_at": time.Now(),
	})
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError("删除数据失败"))
	}
}

func NewStoreService(ctx context.Context) *StoreService {
	inst := new(StoreService)
	inst.SetCtx(ctx)
	return inst
}
