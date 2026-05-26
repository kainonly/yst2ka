package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2286(t *testing.T) {
	ctx := context.TODO()
	orgRespTraceNum := ``
	if orgRespTraceNum == `` {
		t.Skip("请先准备有效的原垫资发放通联订单号后再执行真实请求测试")
	}

	dto := yst2ka.NewTx2286Dto(Num(`X`, cfg.EnterpriseCode, `0`), orgRespTraceNum, 100)
	r, err := client.Tx2286(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`result:`, r.Result)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`extendParams:`, r.ExtendParams)
	}
}
