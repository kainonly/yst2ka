package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2294(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	orgRespTraceNum := `20260604110147208501591069`

	dto := yst2ka.NewTx2294Dto(num, 1).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRespUrl(v.Notify(`/tx2294/callback`))

	r, err := client.Tx2294(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)
}
