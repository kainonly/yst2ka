package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/kainonly/yst2ka/util"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1010(t *testing.T) {
	// TM1010：入会/会员签约类接口的基础用例
	// 该用例主要覆盖：敏感字段加密 -> DTO 构建 -> 发起请求 -> 校验响应无错误
	ctx := context.TODO()
	// 生成本次请求的业务流水号/订单号等（具体格式由 Num() 与 cfg 决定）
	num := Num(`X`, cfg.PersonCode, `0`)

	// 证件号码
	// v.Encrypt：对敏感信息做加密，避免明文传输
	cerNum, err := v.Encrypt(`110102200305048508`)
	assert.NoError(t, err)

	// 银行卡号（本用例不绑定银行卡，保留示例以便需要时开启）
	//acctNum, err := v.Encrypt(`6210262695475575477`)
	//assert.NoError(t, err)

	// 构建 TM1010 请求参数
	// NewTm1010Dto(...)：创建 DTO；后续通过链式方法补充可选字段
	dto := yst2ka.NewTm1010Dto(num, cfg.PersonCode, `李一四`, `1`, cerNum).
		// MemberRole：会员角色/归属（示例：门店）
		SetMemberRole(`门店`).
		// Phone：手机号（从测试配置读取，避免写死在用例中）
		SetPhone(cfg.Phone)
	//SetBindType(`8`).
	//SetAcctNum(acctNum)

	// 发起 TM1010 请求
	r, err := client.Tm1010(ctx, dto)
	assert.NoError(t, err)

	// 打印关键响应字段，便于联调时快速定位问题
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251203151027101000526339
	t.Log(`signNum:`, r.SignNum)
}

func TestYst2Ka_Tm1010_Beta(t *testing.T) {
	// Beta 用例：在基础入会的同时，演示“绑定银行卡”流程（随机生成卡号再加密）
	ctx := context.TODO()
	num := Num(`X`, cfg.PersonCode, `0`)

	// 证件号码
	// 与基础用例一致：证件号属于敏感字段，需要加密
	cerNum, err := v.Encrypt(`110102200305048508`)
	assert.NoError(t, err)

	// 银行卡号
	// util.GenerateBankCard：生成一张随机银行卡号用于测试；随后同样加密后再发送
	acctNum, err := v.Encrypt(util.GenerateBankCard())
	assert.NoError(t, err)

	// 绑定银行卡：SetBindType + SetAcctNum 一起使用
	dto := yst2ka.NewTm1010Dto(num, cfg.PersonCode, `李一四`, `1`, cerNum).
		SetMemberRole(`门店`).
		SetPhone(cfg.Phone).
		SetBindType(`8`).
		SetAcctNum(acctNum)

	r, err := client.Tm1010(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum) // 20251203151027101000526339
	t.Log(`signNum:`, r.SignNum)
}
