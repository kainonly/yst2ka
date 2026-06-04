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

	dto := yst2ka.NewTx4013Dto(orderNo, code, `https://example.com/callback`)
	r, err := client.Tx4013(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`payAmount:`, r.PayAmount)
	t.Log(`chnlFrontParamInfo:`, r.ChnlFrontParamInfo)
	t.Log(`orderNo:`, r.OrderNo)
	t.Log(`trxReserve:`, r.TrxReserve)
}
