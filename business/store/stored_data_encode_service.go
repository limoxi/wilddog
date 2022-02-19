package store

import (
	"context"
	"fmt"
	"github.com/limoxi/ghost"
	ghost_utils "github.com/limoxi/ghost/utils"
)

type StoredDataEncodeService struct {
	ghost.DomainService
}

func (this *StoredDataEncodeService) Encode(storedData *StoredData) *EncodedStoredData {
	var jsonData ghost.Map
	err := ghost_utils.Decode(string(storedData.Data), &jsonData)
	if err != nil {
		panic(ghost.NewSystemError(fmt.Sprintf("解析数据出错:id(%d)", storedData.Id)))
	}
	return &EncodedStoredData{
		Id:        storedData.Id,
		Biz:       storedData.Biz,
		Data:      jsonData,
		UpdatedAt: ghost_utils.FormatDatetime(storedData.UpdatedAt),
	}
}

func (this *StoredDataEncodeService) EncodeMany(datas []*StoredData) []*EncodedStoredData {
	encodedRecords := make([]*EncodedStoredData, 0, len(datas))
	for _, storedData := range datas {
		encodedRecords = append(encodedRecords, this.Encode(storedData))
	}
	return encodedRecords
}

func NewStoredDataEncodeService(ctx context.Context) *StoredDataEncodeService {
	inst := new(StoredDataEncodeService)
	inst.SetCtx(ctx)
	return inst
}
