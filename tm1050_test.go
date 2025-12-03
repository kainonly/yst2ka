package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1050(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	dto := yst2ka.NewTm1050Dto(num, cfg.PersonCode, `李一四`, `1`,
		v.Notify(`/tm1052/callback`),
	)
	r, err := client.Tm1050(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signAgreementUrl:`, r.SignAgreementUrl)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
