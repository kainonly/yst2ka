package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1050(t *testing.T) {
	ctx := context.TODO()
	code := `bf10006`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1050Dto(num, code, `李一四`, `1`,
		`https://notify.kainonly.com:8443`,
	)
	r, err := client.Tm1050(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.Phone)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
