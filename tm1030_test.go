package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1030(t *testing.T) {
	ctx := context.TODO()
	code := `10000`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1030Dto(num, code, `15501875915`)
	r, err := client.Tm1030(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
}
