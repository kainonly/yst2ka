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
	code := `2705wx100002`
	num := Num(`X`, code, `0`)

	cerNum, err := help.SM4Encrypt(secretKey, `410725199907022818`)
	assert.NoError(t, err)

	acctNum, err := help.SM4Encrypt(secretKey, `6217858000141669850`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1010Dto(num, code, `苏大大`, `1`, cerNum).
		SetMemberRole(`分销方`).
		SetPhone(`15617906676`).
		SetBindType(`8`).
		SetAcctNum(acctNum)

	r, err := x.Tm1010(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251121135117101000409346
	t.Log(`signNum:`, r.SignNum)
}
