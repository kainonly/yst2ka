package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1032BindPhone(t *testing.T) {
	ctx := context.TODO()
	num0 := Num(`X`, cfg.PersonCode, `0`)
	dto0 := yst2ka.NewTm1030Dto(num0, cfg.PersonCode, cfg.Phone).
		SetNotifyUrl(v.Notify(`/tm1030/callback`))

	r0, err := client.Tm1030(ctx, dto0)
	assert.NoError(t, err)

	num := Num(`X`, cfg.PersonCode, `0`)

	verifyCode := `111111`
	dto := yst2ka.NewTm1032Dto(num, cfg.PersonCode, r0.RespTraceNum, cfg.Phone, verifyCode)
	r, err := client.Tm1032(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, cfg.Phone)
}

func TestYst2Ka_Tm1032Unbind(t *testing.T) {
	ctx := context.TODO()
	num0 := Num(`X`, cfg.PersonCode, `0`)

	dto0 := yst2ka.NewTm1031Dto(num0, cfg.PersonCode, cfg.Phone)
	r0, err := client.Tm1031(ctx, dto0)
	assert.NoError(t, err)

	num := Num(`X`, cfg.PersonCode, `0`)

	verifyCode := `111111`
	dto := yst2ka.NewTm1032Dto(num, cfg.PersonCode, r0.RespTraceNum, cfg.Phone, verifyCode)
	r, err := client.Tm1032(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.Phone)
}
