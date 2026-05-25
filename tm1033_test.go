package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1033(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.EnterpriseCode, `0`)

	acctNum, err := v.Encrypt(`123426789159100`)
	assert.NoError(t, err)

	bankAcctDetail := yst2ka.Tm1033BankAcctDetail{
		AcctAttr:           `1`,
		AcctNum:            acctNum,
		OpenBankNo:         `01020000`,
		OpenBankBranchName: `中国工商银行上海滩分行`,
		PayBankNumber:      `123456789123`,
		OpenBankProvince:   `上海市`,
		OpenBankCity:       `上海市`,
	}

	dto := yst2ka.NewTm1033Dto(num, cfg.EnterpriseCode, bankAcctDetail)
	r, err := client.Tm1033(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
