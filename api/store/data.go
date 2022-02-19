package store

import (
	"github.com/limoxi/ghost"
	"wilddog/business/service"
)

type Data struct {
	ghost.ApiTemplate

	PutParams *struct {
		Biz  string     `json:"biz"`
		Data ghost.Json `json:"data"`
	}

	PostParams *struct {
		Biz  string     `json:"biz"`
		Id   int        `json:"id"`
		Data ghost.Json `json:"data"`
	}

	DeleteParams *struct {
		Biz string `json:"biz"`
		Id  int    `json:"id"`
	}
}

func (this *Data) Resource() string {
	return "store.data"
}

func (this *Data) Put() ghost.Response {
	ctx := this.GetCtx()
	service.NewStoreService(ctx).Add(
		this.PutParams.Biz,
		this.PutParams.Data,
	)
	return ghost.NewJsonResponse(nil)
}

func (this *Data) Post() ghost.Response {
	ctx := this.GetCtx()
	params := this.PostParams
	service.NewStoreService(ctx).Modify(
		params.Biz,
		params.Id,
		params.Data,
	)
	return ghost.NewJsonResponse(nil)
}

func (this *Data) Delete() ghost.Response {
	ctx := this.GetCtx()
	service.NewStoreService(ctx).Delete(
		this.DeleteParams.Biz,
		this.DeleteParams.Id,
	)
	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&Data{})
}
