package store

import (
	"github.com/limoxi/ghost"
	"wilddog/business/service"
	b_store "wilddog/business/store"
)

type Datas struct {
	ghost.ApiTemplate

	GetParams *struct {
		Biz         string    `form:"biz"`
		Filters     ghost.Map `form:"filters"`
		OrderFields []string  `form:"order_fields"`
		CurPage     int       `form:"cur_page"`
		PageSize    int       `form:"page_size"`
	}

	PutParams *struct {
		Biz   string       `json:"biz"`
		Datas []ghost.Json `json:"datas"`
	}

	PostParams *struct {
		Biz   string       `json:"biz"`
		Datas []ghost.Json `json:"datas"`
	}

	DeleteParams *struct {
		Biz string `json:"biz"`
	}
}

func (this *Datas) Resource() string {
	return "store.datas"
}

func (this *Datas) Get() ghost.Response {
	ctx := this.GetCtx()
	params := this.GetParams
	paginator := ghost.NewPaginator(params.CurPage, params.PageSize)
	storedDatas := b_store.NewStoredDataRepository(ctx).GetPagedDatasByBiz(
		params.Biz,
		paginator,
	)
	return ghost.NewJsonResponse(ghost.Map{
		"datas":     b_store.NewStoredDataEncodeService(ctx).EncodeMany(storedDatas),
		"page_info": paginator.ToMap(),
	})
}

func (this *Datas) Put() ghost.Response {
	ctx := this.GetCtx()
	service.NewStoreService(ctx).AddMany(
		this.PutParams.Biz,
		this.PutParams.Datas,
		true,
	)
	return ghost.NewJsonResponse(nil)
}

func (this *Datas) Post() ghost.Response {
	ctx := this.GetCtx()
	service.NewStoreService(ctx).AddMany(
		this.PutParams.Biz,
		this.PutParams.Datas,
		false,
	)
	return ghost.NewJsonResponse(nil)
}

func (this *Datas) Delete() ghost.Response {
	ctx := this.GetCtx()
	service.NewStoreService(ctx).DeleteByBiz(
		this.PutParams.Biz,
	)
	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&Datas{})
}
