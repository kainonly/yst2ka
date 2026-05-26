package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tq3007(t *testing.T) {
	ctx := context.TODO()
	batchNo := ``
	if batchNo == `` {
		t.Skip("请先准备有效的批次号后再执行真实请求测试")
	}

	dto := yst2ka.NewTq3007Dto(batchNo)
	r, err := client.Tq3007(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`batchNo:`, r.BatchNo)
		t.Log(`status:`, r.Status)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`resultList:`, r.ResultList)
	}
}
