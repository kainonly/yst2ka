package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2090(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	orgRespTraceNum := ``
	receiverSignNum := ``

	dto := yst2ka.NewTx2090Dto(num,
		[]*yst2ka.Tx2090Receiver{
			yst2ka.NewTx2090Receiver(receiverSignNum, 100),
		},
	).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRespUrl(v.Notify(`/tx2090/callback`))

	r, err := client.Tx2090(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
}
