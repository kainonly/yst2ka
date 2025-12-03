package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1023(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTm1023Dto(`yunBizUserId_B2C`)

	var r yst2ka.Tm1023Result[yst2ka.Tm1023BalanceDetail]
	err := client.Tm1023(ctx, dto, &r)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`balanceDetail:`, r.BalanceDetail)
}

func TestYst2Ka_GetPlatformBalanceDetail(t *testing.T) {
	ctx := context.TODO()
	balance, err := client.GetPlatformBalanceDetail(ctx)
	assert.NoError(t, err)

	t.Log(`acctNum:`, balance.AcctNum)
	t.Log(`acctType:`, balance.AcctType)
	t.Log(`totalAmt:`, balance.TotalAmt)
	t.Log(`transitAmt:`, balance.TransitAmt)
	t.Log(`availableAmt:`, balance.AvailableAmt)
	t.Log(`yesAmt:`, balance.YesAmt)
	t.Log(`retentionLimitAmt:`, balance.RetentionLimitAmt)
}

func TestYst2Ka_GetMemberBalanceDetails(t *testing.T) {
	ctx := context.TODO()
	details, err := client.GetMemberBalanceDetails(ctx, `SUP10000`)
	assert.NoError(t, err)

	for i, detail := range details {
		t.Logf(`========== %d Start ==========`, i)
		t.Log(`acctNum:`, detail.AcctNum)
		t.Log(`acctType:`, detail.AcctType)
		t.Log(`totalAmt:`, detail.TotalAmt)
		t.Log(`transitAmt:`, detail.TransitAmt)
		t.Log(`availableAmt:`, detail.AvailableAmt)
		t.Log(`yesAmt:`, detail.YesAmt)
		t.Log(`retentionLimitAmt:`, detail.RetentionLimitAmt)
		t.Logf(`========== %d End ==========`, i)
	}
}
