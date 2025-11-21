package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/go/help"
	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1020(t *testing.T) {
	ctx := context.TODO()
	code := `SUP10000`
	num := Num(`X`, code, `0`)

	legalPersonCerNum, err := help.SM4Encrypt(secretKey, `51370119380325580x`)
	assert.NoError(t, err)

	acctNum, err := help.SM4Encrypt(secretKey, `123426789159100`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1020Dto(num, code, `https://notify.kainonly.com:8443`).
		SetMemberRole(`门店`).
		SetEnterpriseBaseInfo(yst2ka.EnterpriseBaseInfo{
			EnterpriseName:            "竹溪县子怡鞋店",
			EnterpriseNature:          "2",
			AddressCode:               "310115",
			EnterpriseAdress:          "上海市浦东新区金桥镇",
			UnifiedSocialCredit:       "92420324MA4D68J28J",
			BusLicenseValidate:        "9999-12-31",
			LegalPersonName:           "王三华",
			LegalPersonCerType:        "1",
			LegalPersonCerNum:         legalPersonCerNum,
			IdValidateStart:           "2023-12-31",
			IdValidateEnd:             "9999-12-31",
			LegalPersonPhone:          "12312341234",
			LegpCerFrontFileId:        "",
			LegpCerBackFileId:         "",
			UnifiedSocialCreditFileId: "",
		}).
		SetBankAcctDetail(yst2ka.BankAcctDetail{
			AcctAttr:           "1",
			AcctNum:            acctNum,
			BankReservePhone:   "12312341234",
			OpenBankNo:         "01020000",
			OpenBankBranchName: "中国工商银行上海滩分行",
			PayBankNumber:      "123456789123",
			OpenBankProvince:   "上海市",
			OpenBankCity:       "上海市",
		})

	r, err := x.Tm1020(ctx, dto)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
