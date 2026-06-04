package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1015Dto struct {
	ReqTraceNum   string `json:"reqTraceNum"`          // 请求流水号
	SignNum       string `json:"signNum"`              // 商户会员编号
	MemberRole    string `json:"memberRole,omitempty"` // 会员角色，未注册时按文档上送
	Name          string `json:"name"`                 // 姓名
	CerType       string `json:"cerType"`              // 证件类型
	CerNum        string `json:"cerNum"`               // 证件号码，SM4 加密
	AcctNum       string `json:"acctNum"`              // 银行卡号，SM4 加密
	Phone         string `json:"phone"`                // 银行预留手机
	BindType      string `json:"bindType"`             // 绑卡方式
	AgreementNo   string `json:"agreementNo"`          // 签约协议号
	AgreeMerchant string `json:"agreeMerchant"`        // 签约商户号
}

func NewTm1015Dto(reqTraceNum string, signNum string, name string, cerType string, cerNum string, acctNum string, phone string, bindType string, agreementNo string, agreeMerchant string) *Tm1015Dto {
	return &Tm1015Dto{
		ReqTraceNum:   reqTraceNum,
		SignNum:       signNum,
		Name:          name,
		CerType:       cerType,
		CerNum:        cerNum,
		AcctNum:       acctNum,
		Phone:         phone,
		BindType:      bindType,
		AgreementNo:   agreementNo,
		AgreeMerchant: agreeMerchant,
	}
}

func (x *Tm1015Dto) SetMemberRole(v string) *Tm1015Dto {
	x.MemberRole = v
	return x
}

type Tm1015Result struct {
	RespTraceNum string `json:"respTraceNum,omitempty"` // 响应流水号
	SignNum      string `json:"signNum,omitempty"`      // 商户会员编号
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
}

func (x *Yst2Ka) Tm1015(ctx context.Context, dto *Tm1015Dto) (_ *Tm1015Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `1015`, data); err != nil {
		return
	}

	var result Tm1015Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
