package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1032Dto struct {
	ReqTraceNum       string `json:"reqTraceNum"`       // 商户请求流水号
	SignNum           string `json:"signNum"`           // 商户会员编号
	ApplyRespTraceNum string `json:"applyRespTraceNum"` // 申请响应业务关联流水号
	Phone             string `json:"phone"`             // 绑定或解绑手机
	VerifyCode        string `json:"verifyCode"`        // 短信验证码
}

func NewTm1032Dto(reqTraceNum string, signNum string, applyRespTraceNum string, phone string, verifyCode string) *Tm1032Dto {
	return &Tm1032Dto{
		ReqTraceNum:       reqTraceNum,
		SignNum:           signNum,
		ApplyRespTraceNum: applyRespTraceNum,
		Phone:             phone,
		VerifyCode:        verifyCode,
	}
}

type Tm1032Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	Phone        string `json:"phone"`        // 绑定或解绑手机号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tm1032(ctx context.Context, dto *Tm1032Dto) (_ *Tm1032Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1032`, data); err != nil {
		return
	}

	var result Tm1032Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
