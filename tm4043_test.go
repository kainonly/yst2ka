package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm4043(t *testing.T) {
	ctx := context.TODO()
	cusID := ``
	authCode := ``
	if cusID == `` || authCode == `` {
		t.Skip("请先准备有效的收银宝商户号和授权码后再执行真实请求测试")
	}

	dto := yst2ka.NewTm4043Dto(Num(`X`, cfg.PersonCode, `0`), cusID, authCode, `01`)
	r, err := client.Tm4043(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`cusId:`, r.CusID)
		t.Log(`acct:`, r.Acct)
	}
}
