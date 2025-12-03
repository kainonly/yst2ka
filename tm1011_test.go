package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1011(t *testing.T) {
	// 以下两种需要确认
	// 6-通联通协议支付签约
	// 7-收银宝快捷支付签约

	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	dto := yst2ka.NewTm1011Dto(num, cfg.PersonCode,
		`20251203151610101000526390`,
		cfg.Phone,
		`277733`,
	)
	r, err := client.Tm1011(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`agreementNo:`, r.AgreementNo)
}
