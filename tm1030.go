package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1030Dto struct {
	ReqTraceNum string         `json:"reqTraceNum"` // 商户请求流水号
	SignNum     string         `json:"signNum"`     // 商户会员编号
	Phone       string         `json:"phone"`       // 绑定手机号
	PhoneType   string         `json:"phoneType"`   // 绑定手机类型
	JumpUrl     string         `json:"jumpUrl"`     // 前端跳转地址
	NotifyUrl   string         `json:"notifyUrl"`   // 签约结果通知地址
	AuthPerInfo map[string]any `json:"authPerInfo"` // 被授权人信息
}

type AuthPerInfo struct {
	AuthPerName    string `json:"authPerName"`    // 被授权人姓名
	AuthPerCerNum  string `json:"authPerCerNum"`  // 被授权人证件号
	AuthPerCerType string `json:"authPerCerType"` // 被授权人证件类型
}

func NewTm1030Dto(reqTraceNum string, signNum string, phone string) *Tm1030Dto {
	return &Tm1030Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		Phone:       phone,
	}
}

func (x *Tm1030Dto) SetPhoneType(v string) *Tm1030Dto {
	x.PhoneType = v
	return x
}

func (x *Tm1030Dto) SetJumpUrl(v string) *Tm1030Dto {
	x.JumpUrl = v
	return x
}

func (x *Tm1030Dto) SetNotifyUrl(v string) *Tm1030Dto {
	x.NotifyUrl = v
	return x
}

type Tm1030Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	SignNum      string `json:"signNum"`      // 商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
	SignUrl      string `json:"signUrl"`      // 授权手机号签约链接
}

func (x *Yst2Ka) Tm1030(ctx context.Context, dto *Tm1010Dto) (_ *Tm1030Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1030`, data); err != nil {
		return
	}

	var result Tm1030Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
