package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx4013(t *testing.T) {
	ctx := context.TODO()
	orderNo := ``
	code := ``
	if orderNo == `` || code == `` {
		t.Skip("请先准备有效的当面付通联订单号和二维码编号后再执行真实请求测试")
	}

	dto := yst2ka.NewTx4013Dto(orderNo, code, `https://example.com/callback`)
	r, err := client.Tx4013(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`payAmount:`, r.PayAmount)
		t.Log(`chnlFrontParamInfo:`, r.ChnlFrontParamInfo)
		t.Log(`orderNo:`, r.OrderNo)
		t.Log(`trxReserve:`, r.TrxReserve)
	}
}
