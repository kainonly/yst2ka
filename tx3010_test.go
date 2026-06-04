package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx3010(t *testing.T) {
	ctx := context.TODO()

	respTraceNum := ``
	verifyCode := ``

	dto := yst2ka.NewTx3010Dto(verifyCode).
		SetRespTraceNum(respTraceNum)

	r, err := client.Tx3010(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`batchNo:`, r.BatchNo)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
}
