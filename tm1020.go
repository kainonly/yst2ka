package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1020Dto struct {
	ReqTraceNum        string                   `json:"reqTraceNum"`        // 商户请求流水号
	SignNum            string                   `json:"signNum"`            // 商户会员编号
	MemberRole         string                   `json:"memberRole"`         // 会员角色
	NotifyUrl          string                   `json:"notifyUrl"`          // 企业会员审核结果通知地址
	EnterpriseBaseInfo Tm1020EnterpriseBaseInfo `json:"enterpriseBaseInfo"` // 企业基本信息
	BankAcctDetail     Tm1020BankAcctDetail     `json:"bankAcctDetail"`     // 银行账户信息
}

type Tm1020EnterpriseBaseInfo struct {
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

func NewTm1020EnterpriseBaseInfo(
	enterpriseName string,
	addressCode string,
	enterpriseAdress string,
	unifiedSocialCredit string,
	legalPersonName string,
	legalPersonCerType string,
	legalPersonCerNum string,
	legalPersonPhone string,
) *Tm1020EnterpriseBaseInfo {
	return &Tm1020EnterpriseBaseInfo{
		EnterpriseName:      enterpriseName,
		AddressCode:         addressCode,
		EnterpriseAdress:    enterpriseAdress,
		UnifiedSocialCredit: unifiedSocialCredit,
		LegalPersonName:     legalPersonName,
		LegalPersonCerType:  legalPersonCerType,
		LegalPersonCerNum:   legalPersonCerNum,
		LegalPersonPhone:    legalPersonPhone,
	}
}

func (t *Tm1020EnterpriseBaseInfo) SetEnterpriseNature(enterpriseNature string) *Tm1020EnterpriseBaseInfo {
	t.EnterpriseNature = enterpriseNature
	return t
}

func (t *Tm1020EnterpriseBaseInfo) SetBusLicenseValidate(busLicenseValidate string) *Tm1020EnterpriseBaseInfo {
	t.BusLicenseValidate = busLicenseValidate
	return t
}

func (t *Tm1020EnterpriseBaseInfo) SetIdValidateStart(idValidateStart string) *Tm1020EnterpriseBaseInfo {
	t.IdValidateStart = idValidateStart
	return t
}

func (t *Tm1020EnterpriseBaseInfo) SetIdValidateEnd(idValidateEnd string) *Tm1020EnterpriseBaseInfo {
	t.IdValidateEnd = idValidateEnd
	return t
}

func (t *Tm1020EnterpriseBaseInfo) SetLegpCerFrontFileId(legpCerFrontFileId string) *Tm1020EnterpriseBaseInfo {
	t.LegpCerFrontFileId = legpCerFrontFileId
	return t
}

func (t *Tm1020EnterpriseBaseInfo) SetLegpCerBackFileId(legpCerBackFileId string) *Tm1020EnterpriseBaseInfo {
	t.LegpCerBackFileId = legpCerBackFileId
	return t
}

func (t *Tm1020EnterpriseBaseInfo) SetUnifiedSocialCreditFileId(unifiedSocialCreditFileId string) *Tm1020EnterpriseBaseInfo {
	t.UnifiedSocialCreditFileId = unifiedSocialCreditFileId
	return t
}

type Tm1020BankAcctDetail struct {
	AcctAttr           string `json:"acctAttr"`           // 账户类型
	AcctNum            string `json:"acctNum"`            // 账号
	BankReservePhone   string `json:"bankReservePhone"`   // 银行预留手机
	OpenBankNo         string `json:"openBankNo"`         // 银行代码
	OpenBankBranchName string `json:"openBankBranchName"` // 开户行支行名称
	PayBankNumber      string `json:"payBankNumber"`      // 支付行号，12位数字
	OpenBankProvince   string `json:"openBankProvince"`   // 开户行所在省
	OpenBankCity       string `json:"openBankCity"`       // 开户行所在市
}

func NewTm1020BankAcctDetail(
	acctNum string,
	openBankProvince string,
	openBankCity string,
) *Tm1020BankAcctDetail {
	return &Tm1020BankAcctDetail{
		AcctNum:          acctNum,
		OpenBankProvince: openBankProvince,
		OpenBankCity:     openBankCity,
	}
}

func (t *Tm1020BankAcctDetail) SetAcctAttr(acctAttr string) *Tm1020BankAcctDetail {
	t.AcctAttr = acctAttr
	return t
}

func (t *Tm1020BankAcctDetail) SetBankReservePhone(bankReservePhone string) *Tm1020BankAcctDetail {
	t.BankReservePhone = bankReservePhone
	return t
}

func (t *Tm1020BankAcctDetail) SetOpenBankNo(openBankNo string) *Tm1020BankAcctDetail {
	t.OpenBankNo = openBankNo
	return t
}

func (t *Tm1020BankAcctDetail) SetOpenBankBranchName(openBankBranchName string) *Tm1020BankAcctDetail {
	t.OpenBankBranchName = openBankBranchName
	return t
}

func (t *Tm1020BankAcctDetail) SetPayBankNumber(payBankNumber string) *Tm1020BankAcctDetail {
	t.PayBankNumber = payBankNumber
	return t
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

func (x *Tm1020Dto) SetEnterpriseBaseInfo(v Tm1020EnterpriseBaseInfo) *Tm1020Dto {
	x.EnterpriseBaseInfo = v
	return x
}

func (x *Tm1020Dto) SetBankAcctDetail(v Tm1020BankAcctDetail) *Tm1020Dto {
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
