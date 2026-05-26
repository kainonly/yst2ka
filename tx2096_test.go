package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2096(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	receiverSignNum := ``
	payMode := yst2ka.NewPayMode()
	if receiverSignNum == `` || len(payMode) == 0 {
		t.Skip("请先根据充值场景填写有效的充值会员编号和 payMode 后再执行真实请求测试")
	}

	dto := yst2ka.NewTx2096Dto(receiverSignNum, num, 100).
		SetPayMode(payMode).
		SetReqsUrl(v.Notify(`/tx2096/return`)).
		SetRespUrl(v.Notify(`/tx2096/callback`)).
		SetGoodsName(`测试商品`)

	r, err := client.Tx2096(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`result:`, r.Result)
		t.Log(`chnlFrontParamInfo:`, r.ChnlFrontParamInfo)
		t.Log(`channelParamInfo:`, r.ChannelParamInfo)
		t.Log(`isPreConsume:`, r.IsPreConsume)
	}
}
