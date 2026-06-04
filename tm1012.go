package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1012Dto struct {
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户请求流水号
	SignNum      string `json:"signNum"`                // 商户会员编号
	MemberRole   string `json:"memberRole,omitempty"`   // 会员角色
	Name         string `json:"name"`                   // 姓名
	JumpPageType string `json:"jumpPageType,omitempty"` // 跳转页面类型，1-H5 页面 2-小程序页面
	BackURL      string `json:"backUrl"`                // 异步通知地址
	JumpURL      string `json:"jumpUrl,omitempty"`      // 成功跳转返回页面地址
}

func NewTm1012Dto(reqTraceNum string, signNum string, name string, backURL string) *Tm1012Dto {
	return &Tm1012Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		Name:        name,
		BackURL:     backURL,
	}
}

func (x *Tm1012Dto) SetMemberRole(v string) *Tm1012Dto {
	x.MemberRole = v
	return x
}

func (x *Tm1012Dto) SetJumpPageType(v string) *Tm1012Dto {
	x.JumpPageType = v
	return x
}

func (x *Tm1012Dto) SetJumpURL(v string) *Tm1012Dto {
	x.JumpURL = v
	return x
}

type Tm1012Result struct {
	RespTraceNum  string `json:"respTraceNum"`            // 响应流水号
	SignNum       string `json:"signNum,omitempty"`       // 商户会员编号
	RespCode      string `json:"respCode"`                // 业务返回码
	RespMsg       string `json:"respMsg,omitempty"`       // 失败原因
	RegInviteLink string `json:"regInviteLink,omitempty"` // 会员注册链接
}

type Tm1012JumpURLResult struct {
	ReqTraceNum         string `json:"reqTraceNum"`                   // 商户请求流水号
	SignNum             string `json:"signNum"`                       // 商户会员编号
	Name                string `json:"name"`                          // 会员名称
	AuditResult         string `json:"auditResult"`                   // 审核结果，2-审核成功 3-审核失败
	AuditTime           string `json:"auditTime"`                     // 审核时间
	Remark              string `json:"remark,omitempty"`              // 人工审核备注
	AuditResultMsg      string `json:"auditResultMsg,omitempty"`      // 失败原因
	Phone               string `json:"phone,omitempty"`               // 绑定手机号
	WithdrawAgreeStatus string `json:"withdrawAgreeStatus,omitempty"` // 账户提现协议状态
	WithdrawAgreeNo     string `json:"withdrawAgreeNo,omitempty"`     // 账户提现协议签约编号
	MembershipNo        string `json:"membershipNo,omitempty"`        // 会员关系证明函签约编号
	MembershipStatus    string `json:"membershipStatus,omitempty"`    // 会员关系证明函协议状态
}

func (x *Yst2Ka) Tm1012(ctx context.Context, dto *Tm1012Dto) (_ *Tm1012Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `1012`, data); err != nil {
		return
	}

	var result Tm1012Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
