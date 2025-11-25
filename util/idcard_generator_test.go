package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateIDCard(t *testing.T) {
	idCard := GenerateIDCard()
	t.Log("生成的身份证号:", idCard)

	// 验证长度
	assert.Equal(t, 18, len(idCard), "身份证号长度应为18位")

	// 验证校验码
	assert.True(t, ValidateIDCard(idCard), "身份证号校验码应正确")
}

func TestGenerateIDCardWithGender(t *testing.T) {
	// 测试生成男性身份证号
	maleCard := GenerateIDCardWithGender("male")
	t.Log("生成的男性身份证号:", maleCard)
	assert.Equal(t, 18, len(maleCard), "身份证号长度应为18位")
	assert.True(t, ValidateIDCard(maleCard), "身份证号校验码应正确")

	info := GetIDCardInfo(maleCard)
	assert.Equal(t, "男", info["gender"], "应为男性")

	// 测试生成女性身份证号
	femaleCard := GenerateIDCardWithGender("female")
	t.Log("生成的女性身份证号:", femaleCard)
	assert.Equal(t, 18, len(femaleCard), "身份证号长度应为18位")
	assert.True(t, ValidateIDCard(femaleCard), "身份证号校验码应正确")

	info = GetIDCardInfo(femaleCard)
	assert.Equal(t, "女", info["gender"], "应为女性")
}

func TestGenerateIDCardWithAge(t *testing.T) {
	// 生成25-35岁的身份证号
	idCard := GenerateIDCardWithAge(25, 35)
	t.Log("生成的25-35岁身份证号:", idCard)

	assert.Equal(t, 18, len(idCard), "身份证号长度应为18位")
	assert.True(t, ValidateIDCard(idCard), "身份证号校验码应正确")

	info := GetIDCardInfo(idCard)
	age := info["age"].(int)
	assert.True(t, age >= 25 && age <= 35, "年龄应在25-35岁之间")
	t.Logf("实际年龄: %d岁", age)
}

func TestValidateIDCard(t *testing.T) {
	testCases := []struct {
		name   string
		idCard string
		valid  bool
	}{
		{"有效身份证号1", "110101199003074477", true},
		{"有效身份证号2", "310115199112203229", true},
		{"有效身份证号3", "42010619740525383X", true},
		{"无效身份证号-长度错误", "12345678901234567", false},
		{"无效身份证号-校验码错误", "110101199003074478", false},
		{"无效身份证号-包含非法字符", "11010119900307447A", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateIDCard(tc.idCard)
			assert.Equal(t, tc.valid, result, tc.name)
		})
	}
}

func TestGenerateIDCardBatch(t *testing.T) {
	count := 10
	idCards := GenerateIDCardBatch(count)

	t.Log("生成的测试身份证号:")
	for i, idCard := range idCards {
		info := GetIDCardInfo(idCard)
		t.Logf("%d. %s (性别:%s, 年龄:%d岁)", i+1, idCard, info["gender"], info["age"])
	}

	// 验证数量
	assert.Equal(t, count, len(idCards), "应生成指定数量的身份证号")

	// 验证每个身份证号
	for _, idCard := range idCards {
		assert.Equal(t, 18, len(idCard), "身份证号长度应为18位")
		assert.True(t, ValidateIDCard(idCard), "身份证号校验码应正确: "+idCard)
	}
}

func TestGetIDCardInfo(t *testing.T) {
	idCard := GenerateIDCard()
	info := GetIDCardInfo(idCard)

	t.Logf("身份证号: %s", idCard)
	t.Logf("详细信息: %+v", info)

	assert.True(t, info["valid"].(bool), "应为有效身份证号")
	assert.NotEmpty(t, info["area_code"], "地区代码不应为空")
	assert.NotEmpty(t, info["birth_date"], "出生日期不应为空")
	assert.NotEmpty(t, info["gender"], "性别不应为空")
	assert.True(t, info["age"].(int) > 0, "年龄应大于0")
}

func TestGeneratedIDCardsAreValid(t *testing.T) {
	// 测试生成100个身份证号，确保都是有效的
	for i := 0; i < 100; i++ {
		idCard := GenerateIDCard()
		assert.True(t, ValidateIDCard(idCard), "生成的身份证号应该都是有效的: "+idCard)
	}
}

func TestGenerateMultipleGenders(t *testing.T) {
	// 测试生成多个不同性别的身份证号
	maleCount := 0
	femaleCount := 0

	for i := 0; i < 50; i++ {
		maleCard := GenerateIDCardWithGender("male")
		info := GetIDCardInfo(maleCard)
		if info["gender"] == "男" {
			maleCount++
		}

		femaleCard := GenerateIDCardWithGender("female")
		info = GetIDCardInfo(femaleCard)
		if info["gender"] == "女" {
			femaleCount++
		}
	}

	t.Logf("生成男性身份证号: %d个", maleCount)
	t.Logf("生成女性身份证号: %d个", femaleCount)

	assert.Equal(t, 50, maleCount, "应生成50个男性身份证号")
	assert.Equal(t, 50, femaleCount, "应生成50个女性身份证号")
}
