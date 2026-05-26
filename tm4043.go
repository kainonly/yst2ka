package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm4043Dto struct {
	ReqTraceNum string `json:"reqTraceNum"`        // 请求流水号
	SybOrgID    string `json:"sybOrgId,omitempty"` // 集团/代理商商户号
	CusID       string `json:"cusId"`              // 收银宝商户号
	AuthCode    string `json:"authCode"`           // 授权码（付款码）
	AuthType    string `json:"authType"`           // 授权码类型，01-微信付款码 02-银联 userAuth
	Identify    string `json:"identify,omitempty"` // 云闪付 UA 标识
	SubAppID    string `json:"subAppid,omitempty"` // 微信支付 appid
}

func NewTm4043Dto(reqTraceNum string, cusID string, authCode string, authType string) *Tm4043Dto {
	return &Tm4043Dto{
		ReqTraceNum: reqTraceNum,
		CusID:       cusID,
		AuthCode:    authCode,
		AuthType:    authType,
	}
}

func (x *Tm4043Dto) SetSybOrgID(v string) *Tm4043Dto {
	x.SybOrgID = v
	return x
}

func (x *Tm4043Dto) SetIdentify(v string) *Tm4043Dto {
	x.Identify = v
	return x
}

func (x *Tm4043Dto) SetSubAppID(v string) *Tm4043Dto {
	x.SubAppID = v
	return x
}

type Tm4043Result struct {
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
	RespTraceNum string `json:"respTraceNum"` // 通联流水号
	CusID        string `json:"cusId"`        // 收银宝商户号
	Acct         string `json:"acct"`         // 支付平台用户标识
}

func (x *Yst2Ka) Tm4043(ctx context.Context, dto *Tm4043Dto) (_ *Tm4043Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `4043`, data); err != nil {
		return
	}

	var result Tm4043Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
