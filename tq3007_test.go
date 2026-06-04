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

	dto := yst2ka.NewTq3007Dto(batchNo)
	r, err := client.Tq3007(ctx, dto)
	assert.NoError(t, err)

	t.Log(`batchNo:`, r.BatchNo)
	t.Log(`status:`, r.Status)
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`resultList:`, r.ResultList)
}
