package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1035(t *testing.T) {
	ctx := context.TODO()

	// 文件ID有效期较短（文档约 30 分钟），请先执行 upload 测试获取后再填入。
	legpCerFrontFileID := ""
	legpCerBackFileID := ""
	unifiedSocialCreditFileID := ""

	legalPersonCerNum, err := v.Encrypt(`51370119380325580x`)
	assert.NoError(t, err)

	num := Num(`X`, cfg.EnterpriseCode, `0`)
	enterpriseBaseInfo := &yst2ka.Tm1035EnterpriseBaseInfo{
		EnterpriseName:            `竹溪县子怡鞋店`,
		AddressCode:               `310115`,
		EnterpriseAdress:          `上海市浦东新区金桥镇`,
		BusLicenseValidate:        `9999-12-31`,
		LegalPersonName:           `王三华`,
		LegalPersonCerType:        `1`,
		LegalPersonCerNum:         legalPersonCerNum,
		IDValidateStart:           `2023-12-31`,
		IDValidateEnd:             `9999-12-31`,
		LegalPersonPhone:          cfg.Phone,
		PublicAcctName:            `竹溪县子怡鞋店`,
		LegpCerFrontFileID:        legpCerFrontFileID,
		LegpCerBackFileID:         legpCerBackFileID,
		UnifiedSocialCreditFileID: unifiedSocialCreditFileID,
	}

	dto := yst2ka.NewTm1035Dto(num, cfg.EnterpriseCode, v.Notify(`/tm1035/callback`), enterpriseBaseInfo)
	r, err := client.Tm1035(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
}
