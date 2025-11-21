package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1011(t *testing.T) {
	ctx := context.TODO()
	code := `2705wxl00001`
	num := Num(`X`, code, `0`)

	dto := yst2ka.NewTm1011Dto(num, code,
		`20240304111308101000950414`,
		`15201933462`,
		`277733`,
	)
	r, err := x.Tm1011(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`agreementNo:`, r.AgreementNo)
}
