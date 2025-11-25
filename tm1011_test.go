package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1011(t *testing.T) {
	ctx := context.TODO()
	code := `bf10005`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1011Dto(num, code,
		`20251125092409101000401842`,
		cfg.Phone,
		`277733`,
	)
	r, err := client.Tm1011(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`agreementNo:`, r.AgreementNo)
}
