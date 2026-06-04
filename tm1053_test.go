package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1053(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	dto := yst2ka.NewTm1053Dto(num, cfg.PersonCode, `李一四`, `1`)
	r, err := client.Tm1053(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signAgreementUrl:`, r.SignAgreementUrl)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
