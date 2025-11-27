package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1032(t *testing.T) {
	ctx := context.TODO()
	code := `bf10006`
	num := Num(`X`, code, `0`)

	applyRespTraceNum := `20251127083616103000400262`
	verifyCode := `111111`
	dto := yst2ka.NewTm1032Dto(num, code, applyRespTraceNum, cfg.Phone, verifyCode)
	r, err := client.Tm1032(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.Phone)
}
