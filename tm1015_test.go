package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1015(t *testing.T) {
	ctx := context.TODO()
	agreementNo := ``
	agreeMerchant := ``
	if agreementNo == `` || agreeMerchant == `` {
		t.Skip("请先准备有效的协议号、签约商户号和对应银行卡信息后再执行真实请求测试")
	}

	cerNum, err := v.Encrypt(`310101199001010011`)
	assert.NoError(t, err)
	acctNum, err := v.Encrypt(`6222021234567890123`)
	assert.NoError(t, err)

	dto := yst2ka.NewTm1015Dto(
		Num(`X`, cfg.PersonCode, `0`),
		cfg.PersonCode,
		`张三`,
		yst2ka.CerTypeIdentityCard,
		cerNum,
		acctNum,
		cfg.Phone,
		yst2ka.BindTypeAllinpayTongAgreementPay,
		agreementNo,
		agreeMerchant,
	)

	r, err := client.Tm1015(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`signNum:`, r.SignNum)
	}
}
