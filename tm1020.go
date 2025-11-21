package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1020Dto struct {
	ReqTraceNum        string             `json:"reqTraceNum"`        // 商户请求流水号
	SignNum            string             `json:"signNum"`            // 商户会员编号
	MemberRole         string             `json:"memberRole"`         // 会员角色
	NotifyUrl          string             `json:"notifyUrl"`          // 企业会员审核结果通知地址
	EnterpriseBaseInfo EnterpriseBaseInfo `json:"enterpriseBaseInfo"` // 企业基本信息
	BankAcctDetail     BankAcctDetail     `json:"bankAcctDetail"`     // 银行账户信息
}

type EnterpriseBaseInfo struct {
	EnterpriseName            string `json:"enterpriseName"`            // 企业名称
	EnterpriseNature          string `json:"enterpriseNature"`          // 企业性质
	AddressCode               string `json:"addressCode"`               // 省市地区码
	EnterpriseAdress          string `json:"enterpriseAdress"`          // 企业注册地址
	UnifiedSocialCredit       string `json:"unifiedSocialCredit"`       // 统一社会信用
	BusLicenseValidate        string `json:"busLicenseValidate"`        // 营业证件有效期
	LegalPersonName           string `json:"legalPersonName"`           // 法人姓名
	LegalPersonCerType        string `json:"legalPersonCerType"`        // 法人证件类型
	LegalPersonCerNum         string `json:"legalPersonCerNum"`         // 法人证件号码
	IdValidateStart           string `json:"idValidateStart"`           // 法人证件有效开始日期
	IdValidateEnd             string `json:"idValidateEnd"`             // 法人证件有效截止日期
	LegalPersonPhone          string `json:"legalPersonPhone"`          // 法人手机号码
	LegpCerFrontFileId        string `json:"legpCerFrontFileId"`        // 法人身份证（肖像面）
	LegpCerBackFileId         string `json:"legpCerBackFileId"`         // 法人身份证（国徽面）
	UnifiedSocialCreditFileId string `json:"unifiedSocialCreditFileId"` // 统一信用证
}

type BankAcctDetail struct {
	AcctAttr           string `json:"acctAttr"`           // 账户类型
	AcctNum            string `json:"acctNum"`            // 账号
	BankReservePhone   string `json:"bankReservePhone"`   // 银行预留手机
	OpenBankNo         string `json:"openBankNo"`         // 银行代码
	OpenBankBranchName string `json:"openBankBranchName"` // 开户行支行名称
	PayBankNumber      string `json:"payBankNumber"`      // 支付行号，12位数字
	OpenBankProvince   string `json:"openBankProvince"`   // 开户行所在省
	OpenBankCity       string `json:"openBankCity"`       // 开户行所在市
}

func NewTm1020Dto(reqTraceNum string, signNum string, notifyUrl string) *Tm1020Dto {
	return &Tm1020Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		NotifyUrl:   notifyUrl,
	}
}

func (x *Tm1020Dto) SetMemberRole(v string) *Tm1020Dto {
	x.MemberRole = v
	return x
}

func (x *Tm1020Dto) SetEnterpriseBaseInfo(v EnterpriseBaseInfo) *Tm1020Dto {
	x.EnterpriseBaseInfo = v
	return x
}

func (x *Tm1020Dto) SetBankAcctDetail(v BankAcctDetail) *Tm1020Dto {
	x.BankAcctDetail = v
	return x
}

type Tm1020Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	SignNum      string `json:"signNum"`      // 商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tm1020(ctx context.Context, dto *Tm1020Dto) (_ *Tm1020Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1020`, data); err != nil {
		return
	}

	var result Tm1020Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
