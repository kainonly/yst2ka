package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx4037(t *testing.T) {
	ctx := context.TODO()
	vspCusID := ``
	if vspCusID == `` {
		t.Skip("请先准备有效的收银宝商户号后再执行真实请求测试")
	}

	dto := yst2ka.NewTx4037Dto(vspCusID, 1000)
	r, err := client.Tx4037(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`amt3:`, r.Amt3)
		t.Log(`fee3:`, r.Fee3)
		t.Log(`amt6:`, r.Amt6)
		t.Log(`fee6:`, r.Fee6)
		t.Log(`amt12:`, r.Amt12)
		t.Log(`fee12:`, r.Fee12)
	}
}
