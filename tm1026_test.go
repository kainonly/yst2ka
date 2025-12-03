package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1026(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	dto := yst2ka.NewTm1026Dto(num)

	r, err := client.Tm1026(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`totalAmt:`, r.TotalAmt)
	t.Log(`cusId:`, r.CusId)
	t.Log(`bankCardNo:`, r.BankCardNo)
	t.Log(`bankTotalAmt:`, r.BankTotalAmt)
	t.Log(`yesterdayBalance:`, r.YesterdayBalance)
}
