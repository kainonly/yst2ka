package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateIDCardAndBankCardTogether(t *testing.T) {
	count := 10

	for i := 0; i < count; i++ {
		idCard := GenerateIDCard()
		bankCard := GenerateBankCard()

		t.Logf("%d. 身份证号: %s, 银行卡号: %s", i+1, idCard, bankCard)

		assert.Equal(t, 18, len(idCard), "身份证号长度应为18位")
		assert.True(t, ValidateIDCard(idCard), "身份证号校验码应正确: "+idCard)

		assert.Equal(t, 19, len(bankCard), "银行卡号长度应为19位")
		assert.Equal(t, "621026", bankCard[:6], "银行卡号应以621026开头")
		assert.True(t, ValidateBankCard(bankCard), "银行卡号校验码应正确: "+bankCard)

		idInfo := GetIDCardInfo(idCard)
		bankInfo := GetBankCardInfo(bankCard)

		assert.True(t, idInfo["valid"].(bool), "身份证号信息应标记为有效")
		assert.True(t, bankInfo["valid"].(bool), "银行卡号信息应标记为有效")
	}
}
