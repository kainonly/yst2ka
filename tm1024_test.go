package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1024(t *testing.T) {
	ctx := context.TODO()
	sybMerchantCode := ``
	dto := yst2ka.NewTm1024Dto(Num(`X`, cfg.PersonCode, `0`), cfg.PersonCode, `set`).
		SetSybMerchantCode(sybMerchantCode)

	r, err := client.Tm1024(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`sybMerchantCodeArray:`, r.SybMerchantCodeArray)
}
