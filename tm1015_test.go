package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1015(t *testing.T) {
	ctx := context.TODO()
	agreementNo := ``
	agreeMerchant := ``

	cerNum, err := v.Encrypt(`310101199001010011`)
	assert.NoError(t, err)
	acctNum, err := v.Encrypt(`6222021234567890123`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1015Dto(
		Num(`X`, cfg.PersonCode, `0`),
		cfg.PersonCode,
		`张三`,
		`1`,
		cerNum,
		acctNum,
		cfg.Phone,
		`6`,
		agreementNo,
		agreeMerchant,
	)

	r, err := client.Tm1015(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
}
