package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm4001(t *testing.T) {
	ctx := context.TODO()
	cusID := ``
	termNo := ``
	if cusID == `` || termNo == `` {
		t.Skip("请先准备有效的收银宝商户号和终端号后再执行真实请求测试")
	}

	dto := yst2ka.NewTm4001Dto(cusID, termNo, `03`)
	r, err := client.Tm4001(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`retCode:`, r.RetCode)
		t.Log(`retMsg:`, r.RetMsg)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`wxState:`, r.WxState)
		t.Log(`alState:`, r.AlState)
		t.Log(`unState:`, r.UnState)
	}
}
