package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm4043(t *testing.T) {
	ctx := context.TODO()
	cusID := ``
	authCode := ``

	dto := yst2ka.NewTm4043Dto(Num(`X`, cfg.PersonCode, `0`), cusID, authCode, `01`)
	r, err := client.Tm4043(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`cusId:`, r.CusID)
	t.Log(`acct:`, r.Acct)
}
