package util

import (
	"math/rand"
	"strings"
	"time"
)

// 统一社会信用代码字符集（不含I、O、Z、S、V）
const socialCreditChars = "0123456789ABCDEFGHJKLMNPQRTUWXY"

// 统一社会信用代码权重
var socialCreditWeights = []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}

// 登记管理部门代码
var registrationDeptCodes = map[string]string{
	"1": "机构编制",
	"5": "民政",
	"9": "工商",
	"Y": "其他",
}

// 机构类别代码
var organizationTypeCodes = map[string]string{
	"1": "企业",
	"2": "个体工商户",
	"3": "农民专业合作社",
	"9": "其他",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateBusinessLicense 生成统一社会信用代码（18位）
func GenerateBusinessLicense() string {
	// 第1位：登记管理部门代码（通常为9-工商）
	deptCode := "9"

	// 第2位：机构类别代码（1-企业，2-个体工商户，3-农民专业合作社）
	orgType := []string{"1", "2", "3"}[rand.Intn(3)]

	// 第3-8位：登记管理机关行政区划码（6位）
	areaCode := generateAreaCode()

	// 第9-17位：主体标识码（组织机构代码）（9位）
	orgCode := generateOrgCode()

	// 前17位
	code := deptCode + orgType + areaCode + orgCode

	// 第18位：校验码
	checkCode := calculateSocialCreditCheckCode(code)

	return code + checkCode
}

// GenerateBusinessLicenseWithType 生成指定类型的统一社会信用代码
// orgType: "enterprise"(企业), "individual"(个体工商户), "cooperative"(农民专业合作社)
func GenerateBusinessLicenseWithType(orgType string) string {
	deptCode := "9"

	var typeCode string
	switch orgType {
	case "enterprise":
		typeCode = "1"
	case "individual":
		typeCode = "2"
	case "cooperative":
		typeCode = "3"
	default:
		typeCode = "1"
	}

	areaCode := generateAreaCode()
	orgCode := generateOrgCode()
	code := deptCode + typeCode + areaCode + orgCode
	checkCode := calculateSocialCreditCheckCode(code)

	return code + checkCode
}

// GenerateBusinessLicenseWithArea 生成指定地区的统一社会信用代码
func GenerateBusinessLicenseWithArea(areaCode string) string {
	if len(areaCode) != 6 {
		areaCode = generateAreaCode()
	}

	deptCode := "9"
	orgType := []string{"1", "2", "3"}[rand.Intn(3)]
	orgCode := generateOrgCode()
	code := deptCode + orgType + areaCode + orgCode
	checkCode := calculateSocialCreditCheckCode(code)

	return code + checkCode
}

// ValidateBusinessLicense 验证统一社会信用代码
func ValidateBusinessLicense(code string) bool {
	// 长度检查
	if len(code) != 18 {
		return false
	}

	// 字符检查
	code = strings.ToUpper(code)
	for _, c := range code {
		if !strings.ContainsRune(socialCreditChars, c) {
			return false
		}
	}

	// 校验码验证
	checkCode := calculateSocialCreditCheckCode(code[:17])
	return code[17:18] == checkCode
}

// GetBusinessLicenseInfo 获取统一社会信用代码信息
func GetBusinessLicenseInfo(code string) map[string]interface{} {
	info := make(map[string]interface{})

	if !ValidateBusinessLicense(code) {
		info["valid"] = false
		return info
	}

	info["valid"] = true
	info["code"] = code

	// 登记管理部门
	deptCode := string(code[0])
	info["dept_code"] = deptCode
	info["dept_name"] = registrationDeptCodes[deptCode]

	// 机构类别
	orgTypeCode := string(code[1])
	info["org_type_code"] = orgTypeCode
	info["org_type_name"] = organizationTypeCodes[orgTypeCode]

	// 行政区划
	areaCode := code[2:8]
	info["area_code"] = areaCode

	// 主体标识码
	orgCode := code[8:17]
	info["org_code"] = orgCode

	// 校验码
	checkCode := string(code[17])
	info["check_code"] = checkCode

	return info
}

// GenerateBusinessLicenseBatch 批量生成统一社会信用代码
func GenerateBusinessLicenseBatch(count int) []string {
	codes := make([]string, count)
	for i := 0; i < count; i++ {
		codes[i] = GenerateBusinessLicense()
	}
	return codes
}

// calculateSocialCreditCheckCode 计算统一社会信用代码校验码
func calculateSocialCreditCheckCode(code string) string {
	if len(code) != 17 {
		return ""
	}

	sum := 0
	for i, c := range code {
		charIndex := strings.IndexRune(socialCreditChars, c)
		if charIndex == -1 {
			return ""
		}
		sum += charIndex * socialCreditWeights[i]
	}

	remainder := sum % 31
	checkValue := 31 - remainder

	if checkValue == 31 {
		return "0"
	}
	return string(socialCreditChars[checkValue])
}

// generateAreaCode 生成6位行政区划码
func generateAreaCode() string {
	// 常用省份代码
	provinceCodes := []string{
		"11", "12", "13", "14", "15", // 北京、天津、河北、山西、内蒙古
		"21", "22", "23", // 辽宁、吉林、黑龙江
		"31", "32", "33", "34", "35", "36", "37", // 上海、江苏、浙江、安徽、福建、江西、山东
		"41", "42", "43", "44", "45", "46", // 河南、湖北、湖南、广东、广西、海南
		"50", "51", "52", "53", "54", // 重庆、四川、贵州、云南、西藏
		"61", "62", "63", "64", "65", // 陕西、甘肃、青海、宁夏、新疆
	}

	province := provinceCodes[rand.Intn(len(provinceCodes))]
	city := rand.Intn(20) + 1
	district := rand.Intn(20) + 1

	return province + pad(city, 2) + pad(district, 2)
}

// generateOrgCode 生成9位组织机构代码
func generateOrgCode() string {
	code := ""
	for i := 0; i < 9; i++ {
		code += string(socialCreditChars[rand.Intn(len(socialCreditChars))])
	}
	return code
}

// pad 数字补零
func pad(num int, length int) string {
	str := string(rune('0'+num/10)) + string(rune('0'+num%10))
	if len(str) < length {
		return strings.Repeat("0", length-len(str)) + str
	}
	return str
}
