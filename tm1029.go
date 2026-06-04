package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1029Dto struct {
	ReqTraceNum string `json:"reqTraceNum"` // 商户请求流水号
	SignNum     string `json:"signNum"`     // 商户会员编号
	NotifyUrl   string `json:"notifyUrl"`   // 个人会员支付账户开户结果通知地址
	JumpUrl     string `json:"jumpUrl"`     // 前端回调地址，H5 开户流程结束后应跳转回的页面地址
}

func NewTm1029Dto(reqTraceNum string, signNum string, notifyUrl string) *Tm1029Dto {
	return &Tm1029Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		NotifyUrl:   notifyUrl,
	}
}

func (x *Tm1029Dto) SetJumpUrl(v string) *Tm1029Dto {
	x.JumpUrl = v
	return x
}

type Tm1029Result struct {
	RespTraceNum   string `json:"respTraceNum,omitempty"`   // 响应流水号
	RespCode       string `json:"respCode"`                 // 业务返回码
	RespMsg        string `json:"respMsg,omitempty"`        // 业务返回说明
	SignNum        string `json:"signNum"`                  // 商户会员编号
	OpenAcctStatus string `json:"openAcctStatus,omitempty"` // 开户受理状态，0：已受理 1：受理成功 2：受理失败
	OpenAcctUrl    string `json:"openAcctUrl,omitempty"`    // 个人支付账户开户H5链接
}

type Tm1029NotifyResult struct {
	ReqTraceNum      string `json:"reqTraceNum"`               // 商户请求流水号
	RespTraceNum     string `json:"respTraceNum,omitempty"`    // 响应流水号
	SignNum          string `json:"signNum"`                   // 商户会员编号
	OpenAccoutResult string `json:"openAccoutResult"`          // 支付账户开户结果，0：开户成功 1：开户失败
	PayAcctNo        string `json:"payAcctNo,omitempty"`       // 支付账户号
	PayAcctNoStatus  string `json:"payAcctNoStatus,omitempty"` // 支付账户状态
	PayAcctType      string `json:"payAcctType,omitempty"`     // 支付账户分类，1：1类 2：2类 3：3类
	RespMsg          string `json:"respMsg,omitempty"`         // 失败说明
}

func (x *Yst2Ka) Tm1029(ctx context.Context, dto *Tm1029Dto) (_ *Tm1029Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1029`, data); err != nil {
		return
	}

	var result Tm1029Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
