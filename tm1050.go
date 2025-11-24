package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1050Dto struct {
	ReqTraceNum   string `json:"reqTraceNum"`   // 商户请求流水号
	SignNum       string `json:"signNum"`       // 签约会员编号
	MemberName    string `json:"memberName"`    // 签约会员名称
	AgreementType string `json:"agreementType"` // 协议类型
	JumpPageType  string `json:"jumpPageType"`  // 跳转页面类型
	JumpUrl       string `json:"jumpUrl"`       // 前台跳转地址
	NotifyUrl     string `json:"notifyUrl"`     // 签约结果通知地址
}

func NewTm1050Dto(reqTraceNum string, signNum string, memberName string, agreementType string, notifyUrl string) *Tm1050Dto {
	return &Tm1050Dto{
		ReqTraceNum:   reqTraceNum,
		SignNum:       signNum,
		MemberName:    memberName,
		AgreementType: agreementType,
		NotifyUrl:     notifyUrl,
	}
}

type Tm1050Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	Phone        string `json:"phone"`        // 绑定或解绑手机号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tm1050(ctx context.Context, dto *Tm1050Dto) (_ *Tm1050Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1050`, data); err != nil {
		return
	}

	var result Tm1050Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
