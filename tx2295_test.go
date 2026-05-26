package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2295(t *testing.T) {
	ctx := context.TODO()

	orgRespTraceNum := ``
	if orgRespTraceNum == `` {
		t.Skip("请先准备有效的原通联订单号后再执行真实请求测试")
	}

	dto := yst2ka.NewTx2295Dto(orgRespTraceNum).
		SetCloseReason(`测试关单`)

	r, err := client.Tx2295(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`closeResult:`, r.CloseResult)
		t.Log(`closeFinishTime:`, r.CloseFinishTime)
		t.Log(`result:`, r.Result)
	}
}
