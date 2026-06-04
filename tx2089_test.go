package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2089(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	receivers := []yst2ka.Tx2089Receiver{
		*yst2ka.NewTx2089Receiver(`T1000`, 1),
	}
	dto := yst2ka.NewTx2089Dto(num, receivers, 1, PayMode).
		SetPayAmount(1).
		SetPromotionAmount(0).
		SetReqsUrl(v.Notify(`/tx2089/return`)).
		SetRespUrl(v.Notify(`/tx2089/callback`)).
		SetGoodsName(`测试商品`)

	r, err := client.Tx2089(ctx, dto)
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
