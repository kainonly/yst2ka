package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1022(t *testing.T) {
	ctx := context.TODO()
	code := `SUP10000`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1022Dto(num, code, `https://notify.kainonly.com:8443`).
		SetLegpCerFront(`3320251120151991415567696941058`).
		SetLegpCerBack(`3320251120151991413871184859138`).
		SetUnifiedSocialCredit(`3320251120151991415604724256769`)

	r, err := client.Tm1022(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
}
