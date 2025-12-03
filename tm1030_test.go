package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1030(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.EnterpriseCode, `0`)
	dto := yst2ka.NewTm1030Dto(num, cfg.EnterpriseCode, cfg.Phone).
		SetPhoneType(`1`).
		SetNotifyUrl(v.Notify(`/tm1030/callback`))

	r, err := client.Tm1030(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251203111910103000184619
	t.Log(`signNum:`, r.SignNum)
	t.Log(`signUrl:`, r.SignUrl)
}
