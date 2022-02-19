package store

import (
	"context"
	"github.com/limoxi/ghost"
	"wilddog/business/service"
	common_util "wilddog/common/util"
	m_db "wilddog/db"
)

type StoredDataRepository struct {
	ghost.DomainService
}

func (this *StoredDataRepository) GetByFilters(filters ghost.Map, paginator *ghost.Paginator, args ...[]string) []*StoredData {
	if filters == nil {
		filters = ghost.Map{}
	}
	filters["uuid"] = service.GetUUIDFromContext(this.GetCtx())
	db := this.GetDB().Model(&m_db.Store{}).Where(filters)
	orderBy := "id desc"
	switch len(args) {
	case 1:
		orderBy = common_util.ParseOrderFields(args[0])
	}
	var dbModels []*m_db.Store
	if paginator != nil {
		db = paginator.Paginate(db)
	}
	result := db.Order(orderBy).Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}
	ctx := this.GetCtx()
	datas := make([]*StoredData, 0, len(dbModels))
	for _, dbModel := range dbModels {
		datas = append(datas, NewStoredDataFromDbModel(ctx, dbModel))
	}
	return datas
}

func (this *StoredDataRepository) GetPagedDatasByBiz(biz string, paginator *ghost.Paginator) []*StoredData {
	return this.GetByFilters(ghost.Map{
		"biz": biz,
	}, paginator)
}

func NewStoredDataRepository(ctx context.Context) *StoredDataRepository {
	inst := new(StoredDataRepository)
	inst.SetCtx(ctx)
	return inst
}
