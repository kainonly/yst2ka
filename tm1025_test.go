package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1025(t *testing.T) {
	ctx := context.TODO()
	code := `PAY10000`
	num := Num(`X`, code, `0`)

	legalPersonCerNum, err := v.Encrypt(`51370119380325580x`)
	assert.NoError(t, err)

	enterpriseBaseInfo := yst2ka.NewTm1025EnterpriseBaseInfo(`2028-12-31`, `2021-12-31`, `9999-12-31`, `水产之类`).
		SetAddressCode(`110101`).
		SetEnterpriseName(`竹溪县子怡鞋店`).
		SetEnterpriseAdress(`金沪路55号`).
		SetUnifiedSocialCredit(`92420324MA4D68J28J`).
		SetLegalPersonName(`王三华`).
		SetLegalPersonCerType(`1`).
		SetLegalPersonCerNum(legalPersonCerNum).
		SetLegalPersonPhone(``)

	shareholderCerNum, err := v.Encrypt(`51370119380325580x`)
	assert.NoError(t, err)

	legAndBeneficiaryInfo := yst2ka.NewTm1025LegAndBeneficiaryInfo(
		`中国`, `1`, `1`, `凤岗村`, `01`, `1`, `1`, `1`,
		``, legalPersonCerNum, `9999-12-31`, `1`, `凤岗村`,
		``, shareholderCerNum, `2099-12-31`,
	).
		SetBeneficiaryJudgmentCriteria(`1`).
		SetBeneficiaryJudgmentFile(`61`)

	acctNum, err := v.Encrypt(`123426789159100`)
	assert.NoError(t, err)

	bankAcctDetail := yst2ka.NewTm1025BankAcctDetail(`1`, `上海市`, `上海市`).
		SetOpenBankNo(`01040000`).
		SetAcctAttr(acctNum).
		SetOpenBankBranchName(`中国银行上海滩分行`).
		SetPayBankNumber(`100504100265`)

	attachments := yst2ka.NewTm1025Attachments(
		`3320240402211775150284935442433`,
		`3320240402211775150181608763394`,
		`3320240402211775150181608763394`,
		`3320240402211775150181608763394`,
		`3320240402211775150181608763394`,
	).
		SetUnifiedSocialCreditPhoto(`3320240402211775150284935442433`).
		SetLegalNationalEmblemPhoto(`3320240402211775150589207031809`).
		SetLegalFacePhoto(`3320240402211775150493119721473`).
		SetBeneficiaryFile(`3320240402211775150181608763394`)

	dto := yst2ka.NewTm1025Dto(num, code,
		`https://notify.kainonly.com:8443`,
		enterpriseBaseInfo,
		legAndBeneficiaryInfo,
		bankAcctDetail,
		attachments,
	)

	var r *yst2ka.Tm1025Result
	r, err = client.Tm1025(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`openAcctStatus:`, r.OpenAcctStatus)
}
