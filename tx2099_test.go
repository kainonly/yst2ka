package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2099(t *testing.T) {
	ctx := context.TODO()
	signNum := ``
	acctNum := ``
	inSignNum := ``
	inAcctNum := ``
	if signNum == `` || acctNum == `` || inSignNum == `` || inAcctNum == `` {
		t.Skip("请先准备批量转账的转出与转入支付账户信息后再执行真实请求测试")
	}

	transfer := yst2ka.NewTx2099TransferList(Num(`X`, cfg.EnterpriseCode, `1`), inSignNum, inAcctNum, 100)
	dto := yst2ka.NewTx2099Dto(Num(`B`, cfg.EnterpriseCode, `0`), signNum, acctNum, `1`, []yst2ka.Tx2099TransferList{*transfer})
	r, err := client.Tx2099(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`batchNo:`, r.BatchNo)
		t.Log(`authWay:`, r.AuthWay)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
	}
}
