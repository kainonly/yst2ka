package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm2299(t *testing.T) {
	ctx := context.TODO()
	cusID := ``
	if cusID == `` {
		t.Skip("请先准备有效的收银宝商户号和调拨场景后再执行真实请求测试")
	}

	dto := yst2ka.NewTm2299Dto(Num(`X`, cfg.EnterpriseCode, `0`), `1020`, 100, cusID)
	r, err := client.Tm2299(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`result:`, r.Result)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
	}
}
