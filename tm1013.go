package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1013Dto struct {
	ReqTraceNum      string `json:"reqTraceNum"`                // 商户请求流水号
	SignNum          string `json:"signNum"`                    // 商户会员编号
	MemberRole       string `json:"memberRole,omitempty"`       // 会员角色
	EnterpriseName   string `json:"enterpriseName"`             // 企业名称
	EnterpriseNature string `json:"enterpriseNature,omitempty"` // 企业性质，1-企业 2-个体工商户 3-事业单位
	JumpPageType     string `json:"jumpPageType,omitempty"`     // 跳转页面类型，1-H5 页面 2-小程序页面
	BackURL          string `json:"backUrl"`                    // 异步通知地址
	JumpURL          string `json:"jumpUrl,omitempty"`          // 成功跳转返回页面地址
}

func NewTm1013Dto(reqTraceNum string, signNum string, enterpriseName string, backURL string) *Tm1013Dto {
	return &Tm1013Dto{
		ReqTraceNum:    reqTraceNum,
		SignNum:        signNum,
		EnterpriseName: enterpriseName,
		BackURL:        backURL,
	}
}

func (x *Tm1013Dto) SetMemberRole(v string) *Tm1013Dto {
	x.MemberRole = v
	return x
}

func (x *Tm1013Dto) SetEnterpriseNature(v string) *Tm1013Dto {
	x.EnterpriseNature = v
	return x
}

func (x *Tm1013Dto) SetJumpPageType(v string) *Tm1013Dto {
	x.JumpPageType = v
	return x
}

func (x *Tm1013Dto) SetJumpURL(v string) *Tm1013Dto {
	x.JumpURL = v
	return x
}

type Tm1013Result struct {
	RespTraceNum  string `json:"respTraceNum"`            // 响应流水号
	SignNum       string `json:"signNum,omitempty"`       // 商户会员编号
	RespCode      string `json:"respCode"`                // 业务返回码
	RespMsg       string `json:"respMsg,omitempty"`       // 失败原因
	RegInviteLink string `json:"regInviteLink,omitempty"` // 会员注册链接
}

type Tm1013JumpURLResult struct {
	ReqTraceNum              string     `json:"reqTraceNum"`                        // 商户请求流水号
	SignNum                  string     `json:"signNum"`                            // 商户会员编号
	EnterpriseName           string     `json:"enterpriseName"`                     // 企业名称
	AuditResult              string     `json:"auditResult"`                        // 审核结果，2-审核成功 3-审核失败
	AuditTime                string     `json:"auditTime"`                          // 审核时间
	Remark                   string     `json:"remark,omitempty"`                   // 人工审核备注
	AuditResultMsg           string     `json:"auditResultMsg,omitempty"`           // 失败原因
	EnterpriseCompareResult  string     `json:"enterpriseCompareResult,omitempty"`  // OCR 识别与工商认证信息是否一致
	LegalPersonCompareResult string     `json:"legalPersonCompareResult,omitempty"` // OCR 识别与法人实名信息是否一致
	Phone                    string     `json:"phone,omitempty"`                    // 绑定手机号
	BindPhoneType            string     `json:"bindPhoneType,omitempty"`            // 绑定手机号类型
	WithdrawAgreeStatus      SignStatus `json:"withdrawAgreeStatus,omitempty"`      // 企业/个人主体账户提现协议状态
	WithdrawAgreeNo          string     `json:"withdrawAgreeNo,omitempty"`          // 企业/个人主体账户提现协议签约编号
	MembershipNo             string     `json:"membershipNo,omitempty"`             // 会员关系证明函签约编号
	MembershipStatus         SignStatus `json:"membershipStatus,omitempty"`         // 会员关系证明函协议状态
}

func (x *Yst2Ka) Tm1013(ctx context.Context, dto *Tm1013Dto) (_ *Tm1013Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `1013`, data); err != nil {
		return
	}

	var result Tm1013Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
