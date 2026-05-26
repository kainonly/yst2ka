package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1053Dto struct {
	ReqTraceNum   string `json:"reqTraceNum"`   // 商户请求流水号
	SignNum       string `json:"signNum"`       // 商户会员编号
	MemberName    string `json:"memberName"`    // 签约会员名称
	AgreementType string `json:"agreementType"` // 协议类型
	JumpPageType  string `json:"jumpPageType"`  // 跳转页面类型
	JumpUrl       string `json:"jumpUrl"`       // 前台跳转地址
}

func NewTm1053Dto(reqTraceNum string, signNum string, memberName string, agreementType string) *Tm1053Dto {
	return &Tm1053Dto{
		ReqTraceNum:   reqTraceNum,
		SignNum:       signNum,
		MemberName:    memberName,
		AgreementType: agreementType,
	}
}

func (x *Tm1053Dto) SetJumpPageType(v string) *Tm1053Dto {
	x.JumpPageType = v
	return x
}

func (x *Tm1053Dto) SetJumpUrl(v string) *Tm1053Dto {
	x.JumpUrl = v
	return x
}

type Tm1053Result struct {
	RespTraceNum     string `json:"respTraceNum"`     // 响应流水号（业务正常处理返回）
	SignAgreementUrl string `json:"signAgreementUrl"` // 协议签约地址
	RespCode         string `json:"respCode"`         // 业务返回码
	RespMsg          string `json:"respMsg"`          // 失败原因
}

func (x *Yst2Ka) Tm1053(ctx context.Context, dto *Tm1053Dto) (_ *Tm1053Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1053`, data); err != nil {
		return
	}

	var result Tm1053Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
