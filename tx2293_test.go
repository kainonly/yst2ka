package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2293(t *testing.T) {
	ctx := context.TODO()
	openBankNo := ``

	dto := yst2ka.NewTx2293Dto(Num(`X`, cfg.EnterpriseCode, `0`), openBankNo, 100)
	r, err := client.Tx2293(ctx, dto)
	assert.NoError(t, err)

	t.Log(`result:`, r.Result)
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
