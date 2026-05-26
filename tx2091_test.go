package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2091(t *testing.T) {
	ctx := context.TODO()
	batchNo := Num(`B`, cfg.PersonCode, `0`)
	reqTraceNum := Num(`X`, cfg.PersonCode, `0`)

	orgRespTraceNum := ``
	receiverSignNum := ``
	if orgRespTraceNum == `` || receiverSignNum == `` {
		t.Skip("请先准备有效的担保消费申请通联订单号和收款会员；该接口测试链路需要人工前置条件")
	}

	dto := yst2ka.NewTx2091Dto(batchNo,
		[]yst2ka.Tx2091ApplyList{*yst2ka.NewTx2091ApplyList(
			reqTraceNum,
			[]yst2ka.Tx2091ApplyInfo{*yst2ka.NewTx2091ApplyInfo(100).SetOrgRespTraceNum(orgRespTraceNum)},
			receiverSignNum,
			100,
		)},
	).SetRespUrl(v.Notify(`/tx2091/callback`))

	r, err := client.Tx2091(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`batchNo:`, r.BatchNo)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
	}
}
