package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2084(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	signNum := ``
	inSignNum := ``
	if signNum == `` || inSignNum == `` {
		t.Skip("请先准备有效的转出方和转入方会员编号后再执行真实请求测试")
	}

	dto := yst2ka.NewTx2084Dto(num, signNum, inSignNum, 100).
		SetRespUrl(v.Notify(`/tx2084/callback`))

	r, err := client.Tx2084(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`result:`, r.Result)
		t.Log(`authWay:`, r.AuthWay)
	}
}
