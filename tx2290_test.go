package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2290(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	signNum := `T1002`
	acctNum := `6210262695475575477`

	dto := yst2ka.NewTx2290Dto(signNum, num, 1, acctNum).
		SetRespUrl(v.Notify(`/tx2290/callback`))

	r, err := client.Tx2290(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`chnlTradeCode:`, r.ChnlTradeCode)
}
