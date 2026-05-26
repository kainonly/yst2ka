package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2094(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	receiverSignNum := ``
	if receiverSignNum == `` {
		t.Skip("请先准备有效的储值卡核销收款会员，并确认当前环境已开通该业务后再执行真实请求测试")
	}

	dto := yst2ka.NewTx2094Dto(num,
		[]yst2ka.Tx2094ReceiverList{*yst2ka.NewTx2094ReceiverList(receiverSignNum, 100)},
	).
		SetRespUrl(v.Notify(`/tx2094/callback`))

	r, err := client.Tx2094(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`result:`, r.Result)
	}
}
