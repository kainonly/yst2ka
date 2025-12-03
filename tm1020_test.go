package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1020(t *testing.T) {
	ctx := context.TODO()
	code := `VAN1002`
	num := Num(`EN`, code, `0`)

	legalPersonCerNum, err := v.Encrypt(`51370119380325580x`)
	assert.NoError(t, err)

	acctNum, err := v.Encrypt(`123426789159100`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1020Dto(num, code, `https://notify.kainonly.com:8443/tm1020/callback`).
		SetMemberRole(`门店`).
		SetEnterpriseBaseInfo(*yst2ka.NewTm1020EnterpriseBaseInfo(
			"竹溪县子怡鞋店",
			"310115",
			"上海市浦东新区金桥镇",
			"92420324MA4D68J28J",
			"王三华",
			"1",
			legalPersonCerNum,
			"12312341234",
		).
			SetEnterpriseNature("2").
			SetBusLicenseValidate("9999-12-31").
			SetIdValidateStart("2023-12-31").
			SetIdValidateEnd("9999-12-31")).
		SetBankAcctDetail(*yst2ka.NewTm1020BankAcctDetail(
			acctNum,
			"上海市",
			"上海市",
		).
			SetAcctAttr("1").
			SetBankReservePhone("12312341234").
			SetOpenBankNo("01020000").
			SetOpenBankBranchName("中国工商银行上海滩分行").
			SetPayBankNumber("123456789123"))

	r, err := client.Tm1020(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
