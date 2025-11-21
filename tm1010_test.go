package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/go/help"
	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1010(t *testing.T) {
	ctx := context.TODO()
	code := `SUP10001`
	num := Num(`X`, code, `0`)

	cerNum, err := help.SM4Encrypt(secretKey, `51370119380325580x`)
	assert.NoError(t, err)

	acctNum, err := help.SM4Encrypt(secretKey, `6210260123456789012`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1010Dto(num, code, `王三华`, `1`, cerNum).
		SetMemberRole(`门店`).
		SetPhone(`12312341234`).
		//SetBindType(`6`).
		SetAcctNum(acctNum)

	r, err := x.Tm1010(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
}
