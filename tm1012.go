package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1012Dto struct {
	ReqTraceNum       string `json:"reqTraceNum"`       // 请求流水号
	SignNum           string `json:"signNum"`           // 商户会员编号
	ApplyRespTraceNum string `json:"applyRespTraceNum"` // 申请响应业务关联流水号
	Phone             string `json:"phone"`             // 银行预留手机号
	ValidDate         string `json:"validDate"`         // 有效期
	Cvv2              string `json:"cvv2"`              // CVV2
	VerifyCode        string `json:"verifyCode"`        // 短信验证码
}

func NewTm1012Dto(reqTraceNum string, signNum string, applyRespTraceNum string, phone string, verifyCode string) *Tm1011Dto {
	return &Tm1011Dto{
		ReqTraceNum:       reqTraceNum,
		SignNum:           signNum,
		ApplyRespTraceNum: applyRespTraceNum,
		Phone:             phone,
		VerifyCode:        verifyCode,
	}
}

func (x *Tm1012Dto) SetValidDate(v string) *Tm1012Dto {
	x.ValidDate = v
	return x
}

func (x *Tm1012Dto) SetCvv2(v string) *Tm1012Dto {
	x.Cvv2 = v
	return x
}

type Tm1012Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	SignNum      string `json:"signNum"`      // 商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
	AgreementNo  string `json:"agreementNo"`  // 签约协议号
}

func (x *Yst2Ka) Tm1012(ctx context.Context, dto *Tm1011Dto) (_ *Tm1012Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1012`, data); err != nil {
		return
	}

	var result Tm1012Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
