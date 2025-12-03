package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1010(t *testing.T) {
	ctx := context.TODO()
	code := `PS2001`
	num := Num(`X`, code, `0`)

	// 证件号码
	cerNum, err := v.Encrypt(`110102200305048508`)
	assert.NoError(t, err)

	// 银行卡号
	acctNum, err := v.Encrypt(`6210262695475575477`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1010Dto(num, code, `李一四`, `1`, cerNum).
		SetMemberRole(`门店`).
		SetPhone(cfg.Phone).
		SetBindType(`8`).
		SetAcctNum(acctNum)

	r, err := client.Tm1010(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251203141134101000525663
	t.Log(`signNum:`, r.SignNum)
}
