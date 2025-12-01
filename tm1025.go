package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1025Dto struct {
	ReqTraceNum            string                      `json:"reqTraceNum"`            // 请求流水号
	SignNum                string                      `json:"signNum"`                // 商户会员编号
	MemberRole             string                      `json:"memberRole"`             // 会员角色
	EnterpriseNature       string                      `json:"enterpriseNature"`       // 企业性质
	NotifyUrl              string                      `json:"notifyUrl"`              // 企业会员审核结果通知地址
	EnterpriseBaseInfo     Tm1025EnterpriseBaseInfo    `json:"enterpriseBaseInfo"`     // 企业基本信息
	LegaAndBeneficiaryInfo Tm1025LegAndBeneficiaryInfo `json:"legaAndBeneficiaryInfo"` // 受益人信息
	BankAcctDetail         Tm1025BankAcctDetail        `json:"bankAcctDetail"`         // 银行账户信息
	Attachments            Tm1025Attachments           `json:"attachments"`            // 开户附件材料
}

func NewTm1025Dto(
	reqTraceNum string,
	signNum string,
	notifyUrl string,
	enterpriseBaseInfo Tm1025EnterpriseBaseInfo,
	legaAndBeneficiaryInfo Tm1025LegAndBeneficiaryInfo,
	bankAcctDetail Tm1025BankAcctDetail,
	attachments Tm1025Attachments,
) *Tm1025Dto {
	return &Tm1025Dto{
		ReqTraceNum:            reqTraceNum,
		SignNum:                signNum,
		NotifyUrl:              notifyUrl,
		EnterpriseBaseInfo:     enterpriseBaseInfo,
		LegaAndBeneficiaryInfo: legaAndBeneficiaryInfo,
		BankAcctDetail:         bankAcctDetail,
		Attachments:            attachments,
	}
}

func (t *Tm1025Dto) SetMemberRole(memberRole string) *Tm1025Dto {
	t.MemberRole = memberRole
	return t
}

func (t *Tm1025Dto) SetEnterpriseNature(enterpriseNature string) *Tm1025Dto {
	t.EnterpriseNature = enterpriseNature
	return t
}

type Tm1025EnterpriseBaseInfo struct {
	EnterpriseName      string `json:"enterpriseName"`      // 企业名称
	AddressCode         string `json:"addressCode"`         // 地区码
	EnterpriseAdress    string `json:"enterpriseAdress"`    // 企业地址
	UnifiedSocialCredit string `json:"unifiedSocialCredit"` // 统一社会信用
	BusLicenseValidate  string `json:"busLicenseValidate"`  // 营业证件有效期
	LegalPersonName     string `json:"legalPersonName"`     // 法人姓名
	LegalPersonCerType  string `json:"legalPersonCerType"`  // 法人证件类型
	LegalPersonCerNum   string `json:"legalPersonCerNum"`   // 法人证件号码
	IdValidateStart     string `json:"idValidateStart"`     // 法人证件有效开始日期
	IdValidateEnd       string `json:"idValidateEnd"`       // 法人证件有效截止日期
	LegalPersonPhone    string `json:"legalPersonPhone"`    // 法人手机号码
	BusiScope           string `json:"busiScope"`           // 经营内容
}

func NewTm1025EnterpriseBaseInfo(
	busLicenseValidate string,
	idValidateStart string,
	idValidateEnd string,
	busiScope string,
) *Tm1025EnterpriseBaseInfo {
	return &Tm1025EnterpriseBaseInfo{
		BusLicenseValidate: busLicenseValidate,
		IdValidateStart:    idValidateStart,
		IdValidateEnd:      idValidateEnd,
		BusiScope:          busiScope,
	}
}

func (t *Tm1025EnterpriseBaseInfo) SetEnterpriseName(enterpriseName string) *Tm1025EnterpriseBaseInfo {
	t.EnterpriseName = enterpriseName
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetAddressCode(addressCode string) *Tm1025EnterpriseBaseInfo {
	t.AddressCode = addressCode
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetEnterpriseAdress(enterpriseAdress string) *Tm1025EnterpriseBaseInfo {
	t.EnterpriseAdress = enterpriseAdress
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetUnifiedSocialCredit(unifiedSocialCredit string) *Tm1025EnterpriseBaseInfo {
	t.UnifiedSocialCredit = unifiedSocialCredit
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetLegalPersonName(legalPersonName string) *Tm1025EnterpriseBaseInfo {
	t.LegalPersonName = legalPersonName
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetLegalPersonCerType(legalPersonCerType string) *Tm1025EnterpriseBaseInfo {
	t.LegalPersonCerType = legalPersonCerType
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetLegalPersonCerNum(legalPersonCerNum string) *Tm1025EnterpriseBaseInfo {
	t.LegalPersonCerNum = legalPersonCerNum
	return t
}

func (t *Tm1025EnterpriseBaseInfo) SetLegalPersonPhone(legalPersonPhone string) *Tm1025EnterpriseBaseInfo {
	t.LegalPersonPhone = legalPersonPhone
	return t
}

type Tm1025LegAndBeneficiaryInfo struct {
	LegalCountry                string `json:"legalCountry"`                // 法人国籍
	LegalSex                    string `json:"legalSex"`                    // 法人性别
	LegalCareer                 string `json:"legalCareer"`                 // 法人职业
	LegalAddress                string `json:"legalAddress"`                // 法人住址
	MerchantType                string `json:"merchantType"`                // 商户类型
	BeneficiaryJudgmentCriteria string `json:"beneficiaryJudgmentCriteria"` // 受益所有人判定标准
	BeneficiaryJudgmentFile     string `json:"beneficiaryJudgmentFile"`     // 受益所有人证明材料类型
	LegalIsBeneficiary          string `json:"legalIsBeneficiary"`          // 法人是否受益人
	LegalIsShareholder          string `json:"legalIsShareholder"`          // 法人是否为股东人
	BeneficiaryCerType          string `json:"beneficiaryCerType"`          // 受益人证件类型
	BeneficiaryName             string `json:"beneficiaryName"`             // 受益人姓名
	BeneficiaryCerNum           string `json:"beneficiaryCerNum"`           // 受益人证件号码
	BeneficiaryCerValidate      string `json:"beneficiaryCerValidate"`      // 受益人证件有效期
	IsSeniorManagement          string `json:"isSeniorManagement"`          // 是否为高管
	BeneficiaryAddress          string `json:"beneficiaryAddress"`          // 受益人住址
	ShareholderName             string `json:"shareholderName"`             // 控股股东姓名
	ShareholderCerNum           string `json:"shareholderCerNum"`           // 控股股东证件号码
	ShareholderCerValidate      string `json:"shareholderCerValidate"`      // 控股股东证件有效期
}

func NewTm1025LegAndBeneficiaryInfo(
	legalCountry string,
	legalSex string,
	legalCareer string,
	legalAddress string,
	merchantType string,
	legalIsBeneficiary string,
	legalIsShareholder string,
	beneficiaryCerType string,
	beneficiaryName string,
	beneficiaryCerNum string,
	beneficiaryCerValidate string,
	isSeniorManagement string,
	beneficiaryAddress string,
	shareholderName string,
	shareholderCerNum string,
	shareholderCerValidate string,
) *Tm1025LegAndBeneficiaryInfo {
	return &Tm1025LegAndBeneficiaryInfo{
		LegalCountry:           legalCountry,
		LegalSex:               legalSex,
		LegalCareer:            legalCareer,
		LegalAddress:           legalAddress,
		MerchantType:           merchantType,
		LegalIsBeneficiary:     legalIsBeneficiary,
		LegalIsShareholder:     legalIsShareholder,
		BeneficiaryCerType:     beneficiaryCerType,
		BeneficiaryName:        beneficiaryName,
		BeneficiaryCerNum:      beneficiaryCerNum,
		BeneficiaryCerValidate: beneficiaryCerValidate,
		IsSeniorManagement:     isSeniorManagement,
		BeneficiaryAddress:     beneficiaryAddress,
		ShareholderName:        shareholderName,
		ShareholderCerNum:      shareholderCerNum,
		ShareholderCerValidate: shareholderCerValidate,
	}
}

func (t *Tm1025LegAndBeneficiaryInfo) SetBeneficiaryJudgmentCriteria(beneficiaryJudgmentCriteria string) *Tm1025LegAndBeneficiaryInfo {
	t.BeneficiaryJudgmentCriteria = beneficiaryJudgmentCriteria
	return t
}

func (t *Tm1025LegAndBeneficiaryInfo) SetBeneficiaryJudgmentFile(beneficiaryJudgmentFile string) *Tm1025LegAndBeneficiaryInfo {
	t.BeneficiaryJudgmentFile = beneficiaryJudgmentFile
	return t
}

type Tm1025BankAcctDetail struct {
	AcctAttr           string `json:"acctAttr"`           // 账户类型
	AcctNum            string `json:"acctNum"`            // 账号
	BankReservePhone   string `json:"bankReservePhone"`   // 银行预留手机
	OpenBankNo         string `json:"openBankNo"`         // 银行代码
	OpenBankBranchName string `json:"openBankBranchName"` // 开户行支行名称
	PayBankNumber      string `json:"payBankNumber"`      // 支付行号
	OpenBankProvince   string `json:"openBankProvince"`   // 开户行所在省
	OpenBankCity       string `json:"openBankCity"`       // 开户行所在市
}

func NewTm1025BankAcctDetail(
	acctNum string,
	openBankProvince string,
	openBankCity string,
) *Tm1025BankAcctDetail {
	return &Tm1025BankAcctDetail{
		AcctNum:          acctNum,
		OpenBankProvince: openBankProvince,
		OpenBankCity:     openBankCity,
	}
}

func (t *Tm1025BankAcctDetail) SetAcctAttr(acctAttr string) *Tm1025BankAcctDetail {
	t.AcctAttr = acctAttr
	return t
}

func (t *Tm1025BankAcctDetail) SetBankReservePhone(bankReservePhone string) *Tm1025BankAcctDetail {
	t.BankReservePhone = bankReservePhone
	return t
}

func (t *Tm1025BankAcctDetail) SetOpenBankNo(openBankNo string) *Tm1025BankAcctDetail {
	t.OpenBankNo = openBankNo
	return t
}

func (t *Tm1025BankAcctDetail) SetOpenBankBranchName(openBankBranchName string) *Tm1025BankAcctDetail {
	t.OpenBankBranchName = openBankBranchName
	return t
}

func (t *Tm1025BankAcctDetail) SetPayBankNumber(payBankNumber string) *Tm1025BankAcctDetail {
	t.PayBankNumber = payBankNumber
	return t
}

type Tm1025Attachments struct {
	UnifiedSocialCreditPhoto         string `json:"unifiedSocialCreditPhoto"`         // 统一信用证照片
	LegalNationalEmblemPhoto         string `json:"legalNationalEmblemPhoto"`         // 法人证件(国徽面)
	LegalFacePhoto                   string `json:"legalFacePhoto"`                   // 法人证件(肤像面)
	SettleAcctPhoto                  string `json:"settleAcctPhoto"`                  // 结算账户照
	BusinessDoorHeadPhoto            string `json:"businessDoorHeadPhoto"`            // 经营门头照片
	BusinessInteriorPhoto            string `json:"businessInteriorPhoto"`            // 经营内景照片
	AccountManagerWithDoorPhoto      string `json:"accountManagerWithDoorPhoto"`      // 客户经理与门头照
	AccountManagerHoldingIdCardPhoto string `json:"accountManagerHoldingIdCardPhoto"` // 客户经理手持身份证照片
	BeneficiaryFile                  string `json:"beneficiaryFile"`                  // 受益所有人证明材料
}

func NewTm1025Attachments(
	settleAcctPhoto string,
	businessDoorHeadPhoto string,
	businessInteriorPhoto string,
	accountManagerWithDoorPhoto string,
	accountManagerHoldingIdCardPhoto string,
) *Tm1025Attachments {
	return &Tm1025Attachments{
		SettleAcctPhoto:                  settleAcctPhoto,
		BusinessDoorHeadPhoto:            businessDoorHeadPhoto,
		BusinessInteriorPhoto:            businessInteriorPhoto,
		AccountManagerWithDoorPhoto:      accountManagerWithDoorPhoto,
		AccountManagerHoldingIdCardPhoto: accountManagerHoldingIdCardPhoto,
	}
}

func (t *Tm1025Attachments) SetUnifiedSocialCreditPhoto(unifiedSocialCreditPhoto string) *Tm1025Attachments {
	t.UnifiedSocialCreditPhoto = unifiedSocialCreditPhoto
	return t
}

func (t *Tm1025Attachments) SetLegalNationalEmblemPhoto(legalNationalEmblemPhoto string) *Tm1025Attachments {
	t.LegalNationalEmblemPhoto = legalNationalEmblemPhoto
	return t
}

func (t *Tm1025Attachments) SetLegalFacePhoto(legalFacePhoto string) *Tm1025Attachments {
	t.LegalFacePhoto = legalFacePhoto
	return t
}

func (t *Tm1025Attachments) SetBeneficiaryFile(beneficiaryFile string) *Tm1025Attachments {
	t.BeneficiaryFile = beneficiaryFile
	return t
}

type Tm1025Result struct {
	RespTraceNum   string `json:"respTraceNum"`   // 响应流水号
	RespCode       string `json:"respCode"`       // 业务返回码
	RespMsg        string `json:"respMsg"`        // 业务返回说明
	SignNum        string `json:"signNum"`        // 商户会员编号
	OpenAcctStatus string `json:"openAcctStatus"` // 开户受理状态
}

func (x *Yst2Ka) Tm1025(ctx context.Context, dto *Tm1025Dto) (_ *Tm1025Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1025`, data); err != nil {
		return
	}

	var result Tm1025Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
