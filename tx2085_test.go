package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2085(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	receiverSignNum := `T1000`
	dto := yst2ka.NewTx2085Dto(num, receiverSignNum, 1).
		SetPayMode(PayMode).
		SetPayAmount(1).
		SetPromotionAmount(0).
		SetReqsUrl(v.Notify(`/tx2085/return`)).
		SetRespUrl(v.Notify(`/tx2085/callback`)).
		SetGoodsName(`测试商品`)

	r, err := client.Tx2085(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`chnlFrontParamInfo:`, r.ChnlFrontParamInfo)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)
	t.Log(`isPreConsume:`, r.IsPreConsume)
}
