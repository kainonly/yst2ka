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

	transfer := yst2ka.NewTx2099Transfer(Num(`X`, cfg.EnterpriseCode, `1`), inSignNum, inAcctNum, 1)
	dto := yst2ka.NewTx2099Dto(Num(`B`, cfg.EnterpriseCode, `0`), signNum, acctNum, `1`, []yst2ka.Tx2099Transfer{*transfer})
	r, err := client.Tx2099(ctx, dto)
	assert.NoError(t, err)

	t.Log(`batchNo:`, r.BatchNo)
	t.Log(`authWay:`, r.AuthWay)
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
}
