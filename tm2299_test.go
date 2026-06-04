package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm2299(t *testing.T) {
	ctx := context.TODO()
	cusID := ``
	dto := yst2ka.NewTm2299Dto(Num(`X`, cfg.EnterpriseCode, `0`), `1020`, 100, cusID)
	r, err := client.Tm2299(ctx, dto)
	assert.NoError(t, err)

	t.Log(`result:`, r.Result)
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
