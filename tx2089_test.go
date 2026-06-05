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

	// 设定一个担保消费
	receivers := []*yst2ka.Tx2089Receiver{
		yst2ka.NewTx2089Receiver(`T1000`, 11),
	}
	dto := yst2ka.NewTx2089Dto(num, receivers, 11, PayMode).
		SetSignNum(`ANY`).
		SetPayAmount(10).
		SetPromotionAmount(1).
		SetReqsUrl(v.Notify(`tx2089/return`)).
		SetRespUrl(v.Notify(`tx2089/callback`)).
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

	// respCode: 66666
	// respMsg: 业务已受理
	// reqTraceNum: XPS1002-202606041551273140
	// respTraceNum: 20260604155127208901590586
	// result: 0
	// chnlFrontParamInfo: {"chnlPayInfo":"https://syb.allinpay.com/apiweb/h5unionpay/native?key=RPg6fAa%2F8lHy5HpM8g5Q8CzY"}
	// channelParamInfo:
	// isPreConsume:
}

func TestYst2Ka_Tx2089_Refund(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	// 从担保消费中退 1 分
	orgRespTraceNum := `20260604155127208901590586`
	dto := yst2ka.NewTx2294Dto(num, 1).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRefundDetail([]*yst2ka.Tx2294RefundDetail{
			yst2ka.NewTx2294RefundDetail(`T1000`, 1),
		}).
		SetRespUrl(v.Notify(`tx2294/callback`))

	r, err := client.Tx2294(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)

	// respCode: 66666
	// respMsg: 业务已受理
	// reqTraceNum: XPS1002-202606041611360660
	// respTraceNum: 20260604161136229401591234
	// result: 0
	// channelParamInfo:
}

func TestYst2Ka_Tx2089_Sep(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	orgRespTraceNum := `20260604155127208901590586`

	dto := yst2ka.NewTx2090Dto(num,
		[]*yst2ka.Tx2090Receiver{
			yst2ka.NewTx2090Receiver(`T1000`, 9).
				SetSepDetail([]*yst2ka.Tx2090SepDetail{
					yst2ka.NewTx2090SepDetail(`T1001`, 3),
					yst2ka.NewTx2090SepDetail(`T1002`, 3),
					yst2ka.NewTx2090SepDetail(`T1003`, 3),
				}),
		},
	).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRespUrl(v.Notify(`tx2090/callback`))

	r, err := client.Tx2090(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
}

func TestYst2Ka_Tx2089_M(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	// 设定一个担保消费
	receivers := []*yst2ka.Tx2089Receiver{
		yst2ka.NewTx2089Receiver(`T1000`, 4),
		yst2ka.NewTx2089Receiver(`T1001`, 3),
		yst2ka.NewTx2089Receiver(`T1002`, 3),
	}
	dto := yst2ka.NewTx2089Dto(num, receivers, 10, PayMode).
		SetSignNum(`ANY`).
		SetPayAmount(10).
		SetPromotionAmount(0).
		SetReqsUrl(v.Notify(`tx2089/return`)).
		SetRespUrl(v.Notify(`tx2089/callback`)).
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

	// respCode: 66666
	// respMsg: 业务已受理
	// reqTraceNum: XPS1002-202606050848556540
	// respTraceNum: 20260605084855208901596264
	// result: 0
	// chnlFrontParamInfo: {"chnlPayInfo":"https://syb.allinpay.com/apiweb/h5unionpay/native?key=PDWoJ9aIsEzobW7%2BbqKFwrrh"}
	// channelParamInfo:
	// isPreConsume:
}

func TestYst2Ka_Tx2089_M_Refund(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	// 从担保消费中退 1 分
	orgRespTraceNum := `20260605084855208901596264`
	dto := yst2ka.NewTx2294Dto(num, 1).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRefundDetail([]*yst2ka.Tx2294RefundDetail{
			yst2ka.NewTx2294RefundDetail(`T1000`, 1),
		}).
		SetRespUrl(v.Notify(`tx2294/callback`))

	r, err := client.Tx2294(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)

	// respCode: 66666
	// respMsg: 业务已受理
	// reqTraceNum: XPS1002-202606050906329290
	// respTraceNum: 20260605090632229401597079
	// result: 0
	// channelParamInfo:
}

func TestYst2Ka_Tx2089_MOk(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	orgRespTraceNum := `20260605084855208901596264`

	dto := yst2ka.NewTx2090Dto(num,
		[]*yst2ka.Tx2090Receiver{
			yst2ka.NewTx2090Receiver(`T1000`, 3),
			yst2ka.NewTx2090Receiver(`T1001`, 3),
			yst2ka.NewTx2090Receiver(`T1002`, 3),
		},
	).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRespUrl(v.Notify(`tx2090/callback`))

	r, err := client.Tx2090(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)

	// respCode: 66666
	// respMsg: 业务已受理
	// reqTraceNum: XPS1002-202606050908228610
	// respTraceNum: 20260605090822209001597198
	// result: 0
}

func TestYst2Ka_Tx2089_MOk_Query(t *testing.T) {
	ctx := context.TODO()
	dto1 := yst2ka.NewTq3002Dto("20260605084855208901596264")

	r1, err := client.Tq3002(ctx, dto1)
	assert.NoError(t, err)

	t.Log(`result1:`, r1)

	dto2 := yst2ka.NewTq3002Dto("20260605090822209001597198")

	r2, err := client.Tq3002(ctx, dto2)
	assert.NoError(t, err)

	t.Log(`result2:`, r2)
}

func TestYst2Ka_Tx2089_M_Refund_31(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	// 从担保消费确认后中退 3 分
	orgRespTraceNum := `20260605084855208901596264`
	dto := yst2ka.NewTx2294Dto(num, 1).
		SetOrgRespTraceNum(orgRespTraceNum).
		SetRefundDetail([]*yst2ka.Tx2294RefundDetail{
			yst2ka.NewTx2294RefundDetail(`T1002`, 1),
		}).
		SetRespUrl(v.Notify(`tx2294/callback`))

	r, err := client.Tx2294(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)

	// respCode: 66666
	// respMsg: 业务已受理
	// reqTraceNum: XPS1002-202606050906329290
	// respTraceNum: 20260605090632229401597079
	// result: 0
	// channelParamInfo:
}
