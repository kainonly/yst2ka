package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1031(t *testing.T) {
	ctx := context.TODO()
	code := `bf10006`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1031Dto(num, code, cfg.Phone)
	r, err := client.Tm1031(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251127085331103100400341
	t.Log(`signNum:`, r.SignNum)
}
