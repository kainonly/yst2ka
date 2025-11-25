package util

import (
	"fmt"
	"math/rand"
	"time"
)

// 银行卡号前缀(621026开头,19位卡号)
const bankCardPrefix = "621026"

// GenerateBankCard 生成19位银行卡号,第四位为0,交易成功
func GenerateBankCard() string {
	rand.Seed(time.Now().UnixNano())

	// 前缀:621026 (6位)
	// 生成后面12位随机数字(不包括最后一位校验码)
	randomPart := ""
	for i := 0; i < 12; i++ {
		randomPart += fmt.Sprintf("%d", rand.Intn(10))
	}

	// 前18位
	card18 := bankCardPrefix + randomPart

	// 计算Luhn校验码
	checkDigit := calculateLuhnCheckDigit(card18)

	return card18 + fmt.Sprintf("%d", checkDigit)
}

// GenerateBankCardWithFourthDigit 生成指定第四位数字的银行卡号
// fourthDigit: 第四位数字(0-9)
// 第四位为0表示交易成功
func GenerateBankCardWithFourthDigit(fourthDigit int) string {
	if fourthDigit < 0 || fourthDigit > 9 {
		fourthDigit = 0
	}

	rand.Seed(time.Now().UnixNano())

	// 前三位:621
	prefix := "621"

	// 第四位:指定数字
	fourthDigitStr := fmt.Sprintf("%d", fourthDigit)

	// 后面两位:26 (构成完整的621026前缀)
	suffix := "26"

	// 生成后面12位随机数字
	randomPart := ""
	for i := 0; i < 12; i++ {
		randomPart += fmt.Sprintf("%d", rand.Intn(10))
	}

	// 前18位
	card18 := prefix + fourthDigitStr + suffix + randomPart

	// 计算Luhn校验码
	checkDigit := calculateLuhnCheckDigit(card18)

	return card18 + fmt.Sprintf("%d", checkDigit)
}

// GenerateSuccessBankCard 生成交易成功的银行卡号(第四位为0)
func GenerateSuccessBankCard() string {
	return GenerateBankCardWithFourthDigit(0)
}

// calculateLuhnCheckDigit 计算Luhn校验码
func calculateLuhnCheckDigit(cardNumber string) int {
	sum := 0
	isDouble := true

	// 从右往左遍历
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if isDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isDouble = !isDouble
	}

	// 计算校验码
	checkDigit := (10 - (sum % 10)) % 10
	return checkDigit
}

// ValidateBankCard 验证银行卡号(Luhn算法)
func ValidateBankCard(cardNumber string) bool {
	if len(cardNumber) == 0 {
		return false
	}

	sum := 0
	isDouble := false

	// 从右往左遍历
	for i := len(cardNumber) - 1; i >= 0; i-- {
		if cardNumber[i] < '0' || cardNumber[i] > '9' {
			return false
		}

		digit := int(cardNumber[i] - '0')

		if isDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isDouble = !isDouble
	}

	return sum%10 == 0
}

// GenerateBankCardBatch 批量生成银行卡号
func GenerateBankCardBatch(count int) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond) // 确保随机性
		result[i] = GenerateBankCard()
	}
	return result
}

// GenerateSuccessBankCardBatch 批量生成交易成功的银行卡号(第四位为0)
func GenerateSuccessBankCardBatch(count int) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond) // 确保随机性
		result[i] = GenerateSuccessBankCard()
	}
	return result
}

// GetBankCardInfo 获取银行卡号信息
func GetBankCardInfo(cardNumber string) map[string]interface{} {
	info := make(map[string]interface{})

	if len(cardNumber) == 0 {
		info["valid"] = false
		return info
	}

	info["valid"] = ValidateBankCard(cardNumber)
	info["length"] = len(cardNumber)

	if len(cardNumber) >= 6 {
		info["prefix"] = cardNumber[:6]
	}

	if len(cardNumber) >= 4 {
		info["fourth_digit"] = string(cardNumber[3])

		// 第四位为0表示交易成功
		if cardNumber[3] == '0' {
			info["transaction_status"] = "success"
		} else {
			info["transaction_status"] = "unknown"
		}
	}

	return info
}
