package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBankCard(t *testing.T) {
	cardNumber := GenerateBankCard()
	t.Log("生成的银行卡号:", cardNumber)

	// 验证长度
	assert.Equal(t, 19, len(cardNumber), "银行卡号长度应为19位")

	// 验证前缀
	assert.Equal(t, "621026", cardNumber[:6], "银行卡号应以621026开头")

	// 验证Luhn校验码
	assert.True(t, ValidateBankCard(cardNumber), "银行卡号校验码应正确")
}

func TestGenerateSuccessBankCard(t *testing.T) {
	cardNumber := GenerateSuccessBankCard()
	t.Log("生成的交易成功银行卡号:", cardNumber)

	// 验证长度
	assert.Equal(t, 19, len(cardNumber), "银行卡号长度应为19位")

	// 验证前缀
	assert.Equal(t, "621026", cardNumber[:6], "银行卡号应以621026开头")

	// 验证第四位为0
	assert.Equal(t, "0", string(cardNumber[3]), "第四位应为0")

	// 验证Luhn校验码
	assert.True(t, ValidateBankCard(cardNumber), "银行卡号校验码应正确")

	// 验证交易状态
	info := GetBankCardInfo(cardNumber)
	assert.Equal(t, "success", info["transaction_status"], "交易状态应为成功")
}

func TestGenerateBankCardWithFourthDigit(t *testing.T) {
	// 测试生成第四位为0的银行卡号
	cardNumber := GenerateBankCardWithFourthDigit(0)
	t.Log("生成的第四位为0的银行卡号:", cardNumber)

	assert.Equal(t, 19, len(cardNumber), "银行卡号长度应为19位")
	assert.Equal(t, "621026", cardNumber[:6], "银行卡号应以621026开头")
	assert.Equal(t, "0", string(cardNumber[3]), "第四位应为0")
	assert.True(t, ValidateBankCard(cardNumber), "银行卡号校验码应正确")

	// 测试生成第四位为5的银行卡号
	cardNumber = GenerateBankCardWithFourthDigit(5)
	t.Log("生成的第四位为5的银行卡号:", cardNumber)

	assert.Equal(t, 19, len(cardNumber), "银行卡号长度应为19位")
	assert.Equal(t, "5", string(cardNumber[3]), "第四位应为5")
	assert.True(t, ValidateBankCard(cardNumber), "银行卡号校验码应正确")
}

func TestValidateBankCard(t *testing.T) {
	testCases := []struct {
		name       string
		cardNumber string
		valid      bool
	}{
		{"有效银行卡号1", "6210261234567890128", true},
		{"有效银行卡号2", "6210267890123456782", true},
		{"无效银行卡号-校验码错误", "6210261234567890127", false},
		{"无效银行卡号-包含非法字符", "621026123456789012A", false},
		{"空字符串", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateBankCard(tc.cardNumber)
			assert.Equal(t, tc.valid, result, tc.name)
		})
	}
}

func TestGenerateSuccessBankCardBatch(t *testing.T) {
	count := 10
	cardNumbers := GenerateSuccessBankCardBatch(count)

	t.Log("生成的交易成功银行卡号:")
	for i, cardNumber := range cardNumbers {
		info := GetBankCardInfo(cardNumber)
		t.Logf("%d. %s (第四位:%s, 交易状态:%s)",
			i+1,
			cardNumber,
			info["fourth_digit"],
			info["transaction_status"])
	}

	// 验证数量
	assert.Equal(t, count, len(cardNumbers), "应生成指定数量的银行卡号")

	// 验证每个银行卡号
	for _, cardNumber := range cardNumbers {
		assert.Equal(t, 19, len(cardNumber), "银行卡号长度应为19位")
		assert.Equal(t, "621026", cardNumber[:6], "银行卡号应以621026开头")
		assert.Equal(t, "0", string(cardNumber[3]), "第四位应为0")
		assert.True(t, ValidateBankCard(cardNumber), "银行卡号校验码应正确: "+cardNumber)
	}
}

func TestGenerateBankCardBatch(t *testing.T) {
	count := 5
	cardNumbers := GenerateBankCardBatch(count)

	t.Log("生成的银行卡号:")
	for i, cardNumber := range cardNumbers {
		info := GetBankCardInfo(cardNumber)
		t.Logf("%d. %s (前缀:%s, 长度:%d, 有效:%v)",
			i+1,
			cardNumber,
			info["prefix"],
			info["length"],
			info["valid"])
	}

	// 验证数量
	assert.Equal(t, count, len(cardNumbers), "应生成指定数量的银行卡号")

	// 验证每个银行卡号
	for _, cardNumber := range cardNumbers {
		assert.Equal(t, 19, len(cardNumber), "银行卡号长度应为19位")
		assert.Equal(t, "621026", cardNumber[:6], "银行卡号应以621026开头")
		assert.True(t, ValidateBankCard(cardNumber), "银行卡号校验码应正确: "+cardNumber)
	}
}

func TestGetBankCardInfo(t *testing.T) {
	cardNumber := GenerateSuccessBankCard()
	info := GetBankCardInfo(cardNumber)

	t.Logf("银行卡号: %s", cardNumber)
	t.Logf("详细信息: %+v", info)

	assert.True(t, info["valid"].(bool), "应为有效银行卡号")
	assert.Equal(t, 19, info["length"].(int), "长度应为19")
	assert.Equal(t, "621026", info["prefix"].(string), "前缀应为621026")
	assert.Equal(t, "0", info["fourth_digit"].(string), "第四位应为0")
	assert.Equal(t, "success", info["transaction_status"].(string), "交易状态应为成功")
}

func TestGeneratedBankCardsAreValid(t *testing.T) {
	// 测试生成100个银行卡号，确保都是有效的
	for i := 0; i < 100; i++ {
		cardNumber := GenerateBankCard()
		assert.True(t, ValidateBankCard(cardNumber), "生成的银行卡号应该都是有效的: "+cardNumber)
	}
}

func TestGeneratedSuccessBankCardsHaveFourthDigitZero(t *testing.T) {
	// 测试生成50个交易成功的银行卡号，确保第四位都是0
	successCount := 0
	for i := 0; i < 50; i++ {
		cardNumber := GenerateSuccessBankCard()
		info := GetBankCardInfo(cardNumber)

		if info["fourth_digit"] == "0" && info["transaction_status"] == "success" {
			successCount++
		}
	}

	t.Logf("生成交易成功的银行卡号: %d个", successCount)
	assert.Equal(t, 50, successCount, "应生成50个交易成功的银行卡号")
}

func TestLuhnAlgorithm(t *testing.T) {
	// 测试已知的有效银行卡号
	knownValidCards := []string{
		"4532015112830366", // Visa
		"6011111111111117", // Discover
		"378282246310005",  // American Express
	}

	for _, card := range knownValidCards {
		t.Run("Valid_"+card, func(t *testing.T) {
			assert.True(t, ValidateBankCard(card), "已知有效银行卡号应通过验证")
		})
	}

	// 测试已知的无效银行卡号
	knownInvalidCards := []string{
		"4532015112830367", // 最后一位错误
		"6011111111111118", // 最后一位错误
		"378282246310006",  // 最后一位错误
	}

	for _, card := range knownInvalidCards {
		t.Run("Invalid_"+card, func(t *testing.T) {
			assert.False(t, ValidateBankCard(card), "已知无效银行卡号应无法通过验证")
		})
	}
}
