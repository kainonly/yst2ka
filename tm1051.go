package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1051Dto[T AcctAgreementJson | PayAgreementJson] struct {
	ReqTraceNum   string `json:"reqTraceNum"`   // 商户请求流水号
	SignNum       string `json:"signNum"`       // 签约会员编号
	MemberName    string `json:"memberName"`    // 签约会员名称
	AgreementType string `json:"agreementType"` // 协议类型
	AgreementJson T      `json:"agreementJson"` // 签约协议信息
	NotifyUrl     string `json:"notifyUrl"`     // 签约结果通知地址
}

type AcctAgreementJson struct {
	PayeeAgreeToken    string            `json:"payeeAgreeToken,omitempty"`    // 签约会员的收款协议文件token 收款方协议
	WithdrawAgreeToken string            `json:"withdrawAgreeToken,omitempty"` // 签约会员的账户提现协议文件token 若会员需要按照企业/个人主体签账户提现协议,则上送文件token
	AuthPerAgreeInfo   *AuthPerAgreeInfo `json:"authPerAgreeInfo,omitempty"`   // 签约会员的授权委托协议信息 上送,则进行签约 具体字段见【线下授权委托协议(绑定手机号)信息】
}

type AuthPerAgreeInfo struct {
	AuthPhone         string `json:"authPhone"`         // 被授权人手机号
	AuthPerName       string `json:"authPerName"`       // 被授权人姓名
	AuthPerCerNum     string `json:"authPerCerNum"`     // 被授权人证件号 SM4 加密
	AuthPerCerType    string `json:"authPerCerType"`    // 被授权人证件类型 见枚举值,支持多种证件类型
	AuthPerAgreeToken string `json:"authPerAgreeToken"` // 签约会员的授权委托协议文件token
}

func NewAuthPerAgreeInfo(authPhone string, authPerName string, authPerCerNum string, authPerCerType string, authPerAgreeToken string) *AuthPerAgreeInfo {
	return &AuthPerAgreeInfo{
		AuthPhone:         authPhone,
		AuthPerName:       authPerName,
		AuthPerCerNum:     authPerCerNum,
		AuthPerCerType:    authPerCerType,
		AuthPerAgreeToken: authPerAgreeToken,
	}
}

func NewAcctAgreementJson() *AcctAgreementJson {
	return &AcctAgreementJson{}
}

func (x *AcctAgreementJson) SetPayeeAgreeToken(v string) *AcctAgreementJson {
	x.PayeeAgreeToken = v
	return x
}

func (x *AcctAgreementJson) SetWithdrawAgreeToken(v string) *AcctAgreementJson {
	x.WithdrawAgreeToken = v
	return x
}

func (x *AcctAgreementJson) SetAuthPerAgreeInfo(v *AuthPerAgreeInfo) *AcctAgreementJson {
	x.AuthPerAgreeInfo = v
	return x
}

type PayAgreementJson struct {
	PayAcctNoOpenAgreeToken  string `json:"payAcctNoOpenAgreeToken,omitempty"`  // 通联支付账户服务协议文件token
	CoopConfirmToken         string `json:"coopConfirmToken,omitempty"`         // 客户业务合作确认函文件token
	NonNatureCusBenefitToken string `json:"nonNatureCusBenefitToken,omitempty"` // 非自然人客户受益所有人信息登记表文件token
}

func NewPayAgreementJson(payAcctNoOpenAgreeToken string, coopConfirmToken string, nonNatureCusBenefitToken string) *PayAgreementJson {
	return &PayAgreementJson{
		PayAcctNoOpenAgreeToken:  payAcctNoOpenAgreeToken,
		CoopConfirmToken:         coopConfirmToken,
		NonNatureCusBenefitToken: nonNatureCusBenefitToken,
	}
}

func NewTm1051Dto[T AcctAgreementJson | PayAgreementJson](reqTraceNum string, signNum string, memberName string,
	agreementType string, agreementJson T, notifyUrl string) *Tm1051Dto[T] {
	return &Tm1051Dto[T]{
		ReqTraceNum:   reqTraceNum,
		SignNum:       signNum,
		MemberName:    memberName,
		AgreementType: agreementType,
		AgreementJson: agreementJson,
		NotifyUrl:     notifyUrl,
	}
}

type Tm1051Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	SignNum      string `json:"signNum"`      // 签约商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tm1051(ctx context.Context, dto any) (_ *Tm1051Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(dto); err != nil {
		return
	}

	println(data)

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1051`, data); err != nil {
		return
	}

	var result Tm1051Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
