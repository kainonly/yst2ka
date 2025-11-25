package util

import (
	"fmt"
	"math/rand"
	"time"
)

// 地区代码(部分常用地区)
var areaCodes = []string{
	"110101", // 北京东城区
	"110102", // 北京西城区
	"310115", // 上海浦东新区
	"310110", // 上海杨浦区
	"440305", // 深圳南山区
	"440306", // 深圳宝安区
	"420106", // 武汉武昌区
	"420111", // 武汉洪山区
	"510107", // 成都武侯区
	"510108", // 成都成华区
	"330106", // 杭州西湖区
	"330108", // 杭州滨江区
	"320102", // 南京玄武区
	"320104", // 南京秦淮区
	"500101", // 重庆万州区
	"500103", // 重庆渝中区
	"610102", // 西安新城区
	"610103", // 西安碑林区
	"350102", // 福州鼓楼区
	"370102", // 济南历下区
}

// 加权因子
var weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// 校验码
var checkCode = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

// GenerateIDCard 生成身份证号
func GenerateIDCard() string {
	rand.Seed(time.Now().UnixNano())

	// 随机选择地区代码
	areaCode := areaCodes[rand.Intn(len(areaCodes))]

	// 生成出生日期 (1970-2005年之间)
	year := 1970 + rand.Intn(35)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28) // 简化处理,使用1-28日
	birthDate := fmt.Sprintf("%04d%02d%02d", year, month, day)

	// 生成顺序码 (奇数为男,偶数为女)
	sequence := fmt.Sprintf("%03d", rand.Intn(999)+1)

	// 前17位
	id17 := areaCode + birthDate + sequence

	// 计算校验码
	checkDigit := calculateIDCardCheckDigit(id17)

	return id17 + checkDigit
}

// GenerateIDCardWithGender 生成指定性别的身份证号
// gender: "male" 或 "female"
func GenerateIDCardWithGender(gender string) string {
	rand.Seed(time.Now().UnixNano())

	// 随机选择地区代码
	areaCode := areaCodes[rand.Intn(len(areaCodes))]

	// 生成出生日期 (1970-2005年之间)
	year := 1970 + rand.Intn(35)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28)
	birthDate := fmt.Sprintf("%04d%02d%02d", year, month, day)

	// 生成顺序码 (奇数为男,偶数为女)
	var sequence int
	if gender == "male" {
		// 生成奇数
		sequence = rand.Intn(500)*2 + 1
	} else {
		// 生成偶数
		sequence = rand.Intn(500) * 2
		if sequence == 0 {
			sequence = 2
		}
	}
	sequenceStr := fmt.Sprintf("%03d", sequence)

	// 前17位
	id17 := areaCode + birthDate + sequenceStr

	// 计算校验码
	checkDigit := calculateIDCardCheckDigit(id17)

	return id17 + checkDigit
}

// GenerateIDCardWithAge 生成指定年龄范围的身份证号
func GenerateIDCardWithAge(minAge, maxAge int) string {
	rand.Seed(time.Now().UnixNano())

	// 随机选择地区代码
	areaCode := areaCodes[rand.Intn(len(areaCodes))]

	// 根据年龄计算出生年份
	currentYear := time.Now().Year()
	year := currentYear - minAge - rand.Intn(maxAge-minAge+1)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28)
	birthDate := fmt.Sprintf("%04d%02d%02d", year, month, day)

	// 生成顺序码
	sequence := fmt.Sprintf("%03d", rand.Intn(999)+1)

	// 前17位
	id17 := areaCode + birthDate + sequence

	// 计算校验码
	checkDigit := calculateIDCardCheckDigit(id17)

	return id17 + checkDigit
}

// calculateIDCardCheckDigit 计算身份证号校验码
func calculateIDCardCheckDigit(id17 string) string {
	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(id17[i]-'0') * weight[i]
	}
	return checkCode[sum%11]
}

// ValidateIDCard 验证身份证号校验码是否正确
func ValidateIDCard(idCard string) bool {
	if len(idCard) != 18 {
		return false
	}

	// 计算校验码
	sum := 0
	for i := 0; i < 17; i++ {
		if idCard[i] < '0' || idCard[i] > '9' {
			return false
		}
		sum += int(idCard[i]-'0') * weight[i]
	}
	expectedCheck := checkCode[sum%11]
	actualCheck := string(idCard[17])

	return expectedCheck == actualCheck
}

// GenerateIDCardBatch 批量生成身份证号
func GenerateIDCardBatch(count int) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond) // 确保随机性
		result[i] = GenerateIDCard()
	}
	return result
}

// GetIDCardInfo 获取身份证号信息
func GetIDCardInfo(idCard string) map[string]interface{} {
	info := make(map[string]interface{})

	if len(idCard) != 18 {
		info["valid"] = false
		return info
	}

	info["valid"] = ValidateIDCard(idCard)
	info["area_code"] = idCard[:6]
	info["birth_date"] = idCard[6:14]
	info["year"] = idCard[6:10]
	info["month"] = idCard[10:12]
	info["day"] = idCard[12:14]
	info["sequence"] = idCard[14:17]

	// 判断性别
	sequenceNum := int(idCard[16] - '0')
	if sequenceNum%2 == 0 {
		info["gender"] = "女"
	} else {
		info["gender"] = "男"
	}

	// 计算年龄
	birthYear := 0
	fmt.Sscanf(idCard[6:10], "%d", &birthYear)
	currentYear := time.Now().Year()
	info["age"] = currentYear - birthYear

	return info
}
