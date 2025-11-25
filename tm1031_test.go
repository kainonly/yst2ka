package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1031(t *testing.T) {
	ctx := context.TODO()
	code := `2705wx100002`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1031Dto(num, code, cfg.Phone)
	r, err := client.Tm1031(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251125134601103000403719
	t.Log(`signNum:`, r.SignNum)
}
