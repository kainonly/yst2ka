package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBusinessLicense(t *testing.T) {
	code := GenerateBusinessLicense()
	t.Log("生成的统一社会信用代码:", code)

	// 验证长度
	assert.Equal(t, 18, len(code), "统一社会信用代码长度应为18位")

	// 验证校验码
	assert.True(t, ValidateBusinessLicense(code), "统一社会信用代码校验码应正确")
}

func TestGenerateBusinessLicenseWithType(t *testing.T) {
	// 测试生成企业类型
	enterpriseCode := GenerateBusinessLicenseWithType("enterprise")
	t.Log("生成的企业统一社会信用代码:", enterpriseCode)
	assert.Equal(t, 18, len(enterpriseCode), "代码长度应为18位")
	assert.True(t, ValidateBusinessLicense(enterpriseCode), "代码校验应正确")
	assert.Equal(t, "1", string(enterpriseCode[1]), "第2位应为1（企业）")

	info := GetBusinessLicenseInfo(enterpriseCode)
	assert.Equal(t, "企业", info["org_type_name"], "应为企业类型")

	// 测试生成个体工商户类型
	individualCode := GenerateBusinessLicenseWithType("individual")
	t.Log("生成的个体工商户统一社会信用代码:", individualCode)
	assert.Equal(t, 18, len(individualCode), "代码长度应为18位")
	assert.True(t, ValidateBusinessLicense(individualCode), "代码校验应正确")
	assert.Equal(t, "2", string(individualCode[1]), "第2位应为2（个体工商户）")

	info = GetBusinessLicenseInfo(individualCode)
	assert.Equal(t, "个体工商户", info["org_type_name"], "应为个体工商户类型")

	// 测试生成农民专业合作社类型
	cooperativeCode := GenerateBusinessLicenseWithType("cooperative")
	t.Log("生成的农民专业合作社统一社会信用代码:", cooperativeCode)
	assert.Equal(t, 18, len(cooperativeCode), "代码长度应为18位")
	assert.True(t, ValidateBusinessLicense(cooperativeCode), "代码校验应正确")
	assert.Equal(t, "3", string(cooperativeCode[1]), "第2位应为3（农民专业合作社）")

	info = GetBusinessLicenseInfo(cooperativeCode)
	assert.Equal(t, "农民专业合作社", info["org_type_name"], "应为农民专业合作社类型")
}

func TestGenerateBusinessLicenseWithArea(t *testing.T) {
	// 测试北京地区（110000）
	beijingCode := GenerateBusinessLicenseWithArea("110101")
	t.Log("生成的北京地区统一社会信用代码:", beijingCode)

	assert.Equal(t, 18, len(beijingCode), "代码长度应为18位")
	assert.True(t, ValidateBusinessLicense(beijingCode), "代码校验应正确")
	assert.Equal(t, "110101", beijingCode[2:8], "地区代码应为110101")

	info := GetBusinessLicenseInfo(beijingCode)
	assert.Equal(t, "110101", info["area_code"], "地区代码应为110101")

	// 测试上海地区（310000）
	shanghaiCode := GenerateBusinessLicenseWithArea("310115")
	t.Log("生成的上海地区统一社会信用代码:", shanghaiCode)

	assert.Equal(t, 18, len(shanghaiCode), "代码长度应为18位")
	assert.True(t, ValidateBusinessLicense(shanghaiCode), "代码校验应正确")
	assert.Equal(t, "310115", shanghaiCode[2:8], "地区代码应为310115")
}

func TestValidateBusinessLicense(t *testing.T) {
	testCases := []struct {
		name  string
		code  string
		valid bool
	}{
		{"有效代码1", "91310115MA1K3E9C6X", true},
		{"有效代码2", "92420324MA4D68J28J", true},
		{"有效代码3", "91110108MA01GKN831", true},
		{"无效代码-长度错误", "9131011", false},
		{"无效代码-校验码错误", "91310115MA1K3E9C6Y", false},
		{"无效代码-包含非法字符", "9131011SMA1K3E9C6X", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateBusinessLicense(tc.code)
			assert.Equal(t, tc.valid, result, tc.name)
		})
	}
}

func TestGenerateBusinessLicenseBatch(t *testing.T) {
	count := 10
	codes := GenerateBusinessLicenseBatch(count)

	t.Log("生成的测试统一社会信用代码:")
	for i, code := range codes {
		info := GetBusinessLicenseInfo(code)
		t.Logf("%d. %s (类型:%s, 地区:%s)", i+1, code, info["org_type_name"], info["area_code"])
	}

	// 验证数量
	assert.Equal(t, count, len(codes), "应生成指定数量的代码")

	// 验证每个代码
	for _, code := range codes {
		assert.Equal(t, 18, len(code), "代码长度应为18位")
		assert.True(t, ValidateBusinessLicense(code), "代码校验应正确: "+code)
	}
}

func TestGetBusinessLicenseInfo(t *testing.T) {
	code := GenerateBusinessLicense()
	info := GetBusinessLicenseInfo(code)

	t.Logf("统一社会信用代码: %s", code)
	t.Logf("详细信息: %+v", info)

	assert.True(t, info["valid"].(bool), "应为有效代码")
	assert.NotEmpty(t, info["dept_code"], "登记管理部门代码不应为空")
	assert.NotEmpty(t, info["dept_name"], "登记管理部门名称不应为空")
	assert.NotEmpty(t, info["org_type_code"], "机构类别代码不应为空")
	assert.NotEmpty(t, info["org_type_name"], "机构类别名称不应为空")
	assert.NotEmpty(t, info["area_code"], "行政区划代码不应为空")
	assert.NotEmpty(t, info["org_code"], "主体标识码不应为空")
	assert.NotEmpty(t, info["check_code"], "校验码不应为空")
}

func TestGeneratedBusinessLicensesAreValid(t *testing.T) {
	// 测试生成100个统一社会信用代码，确保都是有效的
	for i := 0; i < 100; i++ {
		code := GenerateBusinessLicense()
		assert.True(t, ValidateBusinessLicense(code), "生成的代码应该都是有效的: "+code)
	}
}

func TestGenerateMultipleTypes(t *testing.T) {
	// 测试生成多个不同类型的统一社会信用代码
	enterpriseCount := 0
	individualCount := 0
	cooperativeCount := 0

	for i := 0; i < 30; i++ {
		enterpriseCode := GenerateBusinessLicenseWithType("enterprise")
		info := GetBusinessLicenseInfo(enterpriseCode)
		if info["org_type_name"] == "企业" {
			enterpriseCount++
		}

		individualCode := GenerateBusinessLicenseWithType("individual")
		info = GetBusinessLicenseInfo(individualCode)
		if info["org_type_name"] == "个体工商户" {
			individualCount++
		}

		cooperativeCode := GenerateBusinessLicenseWithType("cooperative")
		info = GetBusinessLicenseInfo(cooperativeCode)
		if info["org_type_name"] == "农民专业合作社" {
			cooperativeCount++
		}
	}

	t.Logf("生成企业代码: %d个", enterpriseCount)
	t.Logf("生成个体工商户代码: %d个", individualCount)
	t.Logf("生成农民专业合作社代码: %d个", cooperativeCount)

	assert.Equal(t, 30, enterpriseCount, "应生成30个企业代码")
	assert.Equal(t, 30, individualCount, "应生成30个个体工商户代码")
	assert.Equal(t, 30, cooperativeCount, "应生成30个农民专业合作社代码")
}

func TestBusinessLicenseFormat(t *testing.T) {
	code := GenerateBusinessLicense()

	// 第1位应为9（工商）
	assert.Equal(t, "9", string(code[0]), "第1位应为9（工商）")

	// 第2位应为1、2或3
	secondChar := string(code[1])
	assert.True(t, secondChar == "1" || secondChar == "2" || secondChar == "3",
		"第2位应为1、2或3")

	// 第3-8位为行政区划码（数字）
	areaCode := code[2:8]
	assert.Equal(t, 6, len(areaCode), "行政区划码应为6位")

	// 第9-17位为主体标识码
	orgCode := code[8:17]
	assert.Equal(t, 9, len(orgCode), "主体标识码应为9位")

	// 第18位为校验码
	checkCode := string(code[17])
	assert.Equal(t, 1, len(checkCode), "校验码应为1位")
}

func TestValidateRealBusinessLicenses(t *testing.T) {
	// 测试一些真实的统一社会信用代码
	realCodes := []string{
		"91310115MA1K3E9C6X",
		"92420324MA4D68J28J",
		"91110108MA01GKN831",
	}

	for _, code := range realCodes {
		t.Run(code, func(t *testing.T) {
			assert.True(t, ValidateBusinessLicense(code), "真实代码应验证通过: "+code)
			info := GetBusinessLicenseInfo(code)
			t.Logf("代码 %s 信息: %+v", code, info)
			assert.True(t, info["valid"].(bool), "代码应为有效")
		})
	}
}
