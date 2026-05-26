package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2086(t *testing.T) {
	ctx := context.TODO()
	receiverSignNum := ``
	if receiverSignNum == `` {
		t.Skip("请先准备可用的垫资收款会员号和垫资发放场景后再执行真实请求测试")
	}

	dto := yst2ka.NewTx2086Dto(receiverSignNum, Num(`X`, cfg.EnterpriseCode, `0`), 100)
	r, err := client.Tx2086(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`result:`, r.Result)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`orderAmount:`, r.OrderAmount)
	}
}
