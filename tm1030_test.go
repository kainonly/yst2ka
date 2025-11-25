package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1030(t *testing.T) {
	ctx := context.TODO()
	code := `bf10006`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1030Dto(num, code, cfg.Phone).
		SetNotifyUrl(`https://notify.kainonly.com:8443`)

	r, err := client.Tm1030(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251125135128103000403759
	t.Log(`signNum:`, r.SignNum)
	t.Log(`signUrl:`, r.SignUrl)
}
