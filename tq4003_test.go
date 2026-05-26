package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tq4003(t *testing.T) {
	ctx := context.TODO()
	respTraceNum := ``
	if respTraceNum == `` {
		t.Skip("请先准备有效的通联订单号或批次号后再执行真实请求测试")
	}

	dto := yst2ka.NewTq4003Dto().SetRespTraceNum(respTraceNum)
	r, err := client.Tq4003(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`fileUrl:`, r.FileURL)
	}
}
