package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1027Dto struct {
	SignNum  string `json:"signNum"`  // 商户会员编号
	InfoType string `json:"InfoType"` // 查询信息类型
}

func NewTm1027Dto(signNum string, infoType string) *Tm1027Dto {
	return &Tm1027Dto{
		SignNum:  signNum,
		InfoType: infoType,
	}
}

type Tm1027Result[T Tm1027PersonInfo | Tm1027EnterpriseInfo | M] struct {
	RespCode          string                  `json:"respCode"`                    // 业务返回码
	RespMsg           string                  `json:"respMsg"`                     // 业务返回说明
	SignNum           string                  `json:"signNum"`                     // 商户会员编号
	MemberBasicInfo   T                       `json:"memberBasicInfo,omitempty"`   // 会员基本信息
	AcctInfo          []Tm1027AcctInfo        `json:"acctInfo,omitempty"`          // 银行账户信息
	AgreementArray    []Tm1027Agreement       `json:"agreementArray,omitempty"`    // 协议信息
	OcrResultJson     Tm1027OcrResultJson     `json:"ocrResultJson,omitempty"`     // 影印件OCR核对结果
	BindPhoneJson     Tm1027BindPhoneJson     `json:"bindPhoneJson,omitempty"`     // 绑定手机号信息
	PayAcctOpenJson   Tm1027PayAcctOpenJson   `json:"payAcctOpenJson,omitempty"`   // 支付账户开户信息
	PayAcctAuditJson  Tm1027PayAcctAuditJson  `json:"payAcctAuditJson,omitempty"`  // 支付账户审核结果详情
	BankSubAcctInfo   M                       `json:"bankSubAcctInfo,omitempty"`   // 银行子账户信息
	SettleAcctInfo    Tm1027SettleAcctInfo    `json:"settleAcctInfo,omitempty"`    // 待结算户信息
	MemberControlInfo Tm1027MemberControlInfo `json:"memberControlInfo,omitempty"` // 会员交易控制类型
}

type Tm1027PersonInfo struct {
	Name             string `json:"name"`             // 姓名
	CerType          string `json:"cerType"`          // 证件类型
	CerNum           string `json:"cerNum"`           // 身份证号码
	IsWithdraw       string `json:"isWithdraw"`       // 是否可提现
	Phone            string `json:"phone"`            // 绑定手机
	IdValidStartDate string `json:"idValidStartDate"` // 证件有效开始日期
	IdValidEndDate   string `json:"idValidEndDate"`   // 证件有效截止日期
	RegisterTime     string `json:"registerTime"`     // 注册时间
	IsRealNameAuth   string `json:"isRealNameAuth"`   // 是否实名认证
	RealNameAuthTime string `json:"realNameAuthTime"` // 实名认证时间
	MemberStatus     string `json:"memberStatus"`     // 会员状态
	MemberRole       string `json:"memberRole"`       // 会员角色
	MemberType       string `json:"memberType"`       // 会员类型
}

type Tm1027EnterpriseInfo struct {
	MemberRole          string `json:"memberRole"`
	MemberType          string `json:"memberType"`
	EnterpriseName      string `json:"enterpriseName"`
	AddressCode         string `json:"addressCode"`
	EnterpriseAdress    string `json:"enterpriseAdress"`
	EnterpriseNature    string `json:"enterpriseNature"`
	UnifiedSocialCredit string `json:"unifiedSocialCredit"`
	BusLicenseValidDate string `json:"busLicenseValidDate"`
	Phone               string `json:"phone"`
	LegalPersonName     string `json:"legalPersonName"`
	LegalPersonCerType  string `json:"legalPersonCerType"`
	LegalPersonCerNum   string `json:"legalPersonCerNum"`
	IdValidStartDate    string `json:"idValidStartDate"`
	IdValidEndDate      string `json:"idValidEndDate"`
	LegalPersonPhone    string `json:"legalPersonPhone"`
	MemberStatus        string `json:"memberStatus"`
	AuditTime           string `json:"auditTime"`
	IsWithdraw          string `json:"isWithdraw"`
	RespMsg             string `json:"respMsg"`
}

type Tm1027AcctInfo struct {
	BankCardNo         string `json:"bankCardNo"`         // 银行卡号
	BankAccountName    string `json:"bankAccountName"`    // 银行户名
	BankName           string `json:"bankName"`           // 银行名称
	BindTime           string `json:"bindTime"`           // 绑定时间
	CardType           string `json:"cardType"`           // 银行卡类型
	BindStatus         string `json:"bindStatus"`         // 绑定状态
	BankReservePhone   string `json:"bankReservePhone"`   // 银行预留手机号码
	BindType           string `json:"bindType"`           // 绑卡方式
	AcctAttr           string `json:"acctAttr"`           // 银行卡/账户属性
	OpenBankBranchName string `json:"openBankBranchName"` // 开户行支行名称
	PayBankNumber      string `json:"payBankNumber"`      // 支付行号
	OpenBankProvince   string `json:"openBankProvince"`   // 开户行所在省
	OpenBankCity       string `json:"openBankCity"`       // 开户行所在市
	IsSpecifyAcct      string `json:"isSpecifyAcct"`      // 是否为支付账户指定出入金银行账户
}

type Tm1027Agreement struct {
	SignAccount       string            `json:"signAccount"`       // 签约户名
	AgreementType     string            `json:"agreementType"`     // 协议类型
	SignResult        string            `json:"signResult"`        // 签约结果
	AgreeNo           string            `json:"agreeNo"`           // 协议编号
	SignTime          string            `json:"signTime"`          // 签约时间
	AnotherMemberInfo map[string]string `json:"anotherMemberInfo"` // 另一方(收款方/分账方)签约信息
}

type Tm1027OcrResultJson struct {
	EnterpriseCompareResult  string `json:"enterpriseCompareResult"`  // OCR识别与企业工商认证信息是否一致
	LegalPersonCompareResult string `json:"legalPersonCompareResult"` // OCR识别与企业法人实名信息是否一致
}

type Tm1027BindPhoneJson struct {
	IsBind string `json:"isBind"` // 是否已绑定手机
	Phone  string `json:"phone"`  // 绑定手机
}

type Tm1027PayAcctOpenJson struct {
	CusId           string `json:"cusId"`           // 统一客户号
	PayAcctNo       string `json:"payAcctNo"`       // 支付账户号
	PayAcctNoStatus string `json:"payAcctNoStatus"` // 支付账户状态
	OpenAcctTime    string `json:"openAcctTime"`    // 开户时间
}

type Tm1027PayAcctAuditJson struct {
	EnterpriseVerifyResult    string `json:"enterpriseVerifyResult"`    // 工商验证
	LegalIdCardVerifyResult   string `json:"legalIdCardVerifyResult"`   // 法人验证
	BankAcctVerifyResult      string `json:"bankAcctVerifyResult"`      // 银行结算账户验证
	UnifiedCreditPhotoResult  string `json:"unifiedCreditPhotoResult"`  // 统一信用证照片验证
	LegalCerPhotoResult       string `json:"legalCerPhotoResult"`       // 法人证件照片验证
	SettleAcctPhotoResult     string `json:"settleAcctPhotoResult"`     // 结算账户照
	BusOutdoorPhotoResult     string `json:"busOutdoorPhotoResult"`     // 经营门头照片
	BusInnerPhotoResult       string `json:"busInnerPhotoResult"`       // 经营内景照
	AcctManOutdoorPhotoResult string `json:"acctManOutdoorPhotoResult"` // 客户经理与门头照
	AcctManWithIdPhotoResult  string `json:"acctManWithIdPhotoResult"`  // 客户经理手持身份证照片
	BusCoopConfirmResult      string `json:"busCoopConfirmResult"`      // 客户业务合作确认函
	NonNatBenfitInfoResult    string `json:"nonNatBenfitInfoResult"`    // 非自然人客户受益所有人信息登记表
	TlPayAcctNoAgreeResult    string `json:"tlPayAcctNoAgreeResult"`    // 通联单位支付账户服务协议
}

type Tm1027SettleAcctInfo struct {
	VspCusId string `json:"vspCusId"` // 收银宝商户号
	AcctNo   string `json:"acctNo"`   // 待结算账户号
	Status   string `json:"status"`   // 待结算户状态
}

type Tm1027MemberControlInfo struct {
	SepOutFlag         string `json:"sepOutFlag"`         // 分账出金
	SepInFlag          string `json:"sepInFlag"`          // 分账入金
	MemberWithdrawFlag string `json:"memberWithdrawFlag"` // 会员提现
}

func (x *Yst2Ka) Tm1027(ctx context.Context, dto *Tm1027Dto, i any) (err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1027`, data); err != nil {
		return
	}

	if err = sonic.UnmarshalString(bizData, i); err != nil {
		return
	}
	return
}

func (x *Yst2Ka) GetPersonInfo(ctx context.Context, signNum string) (info Tm1027PersonInfo, err error) {
	var r Tm1027Result[Tm1027PersonInfo]
	dto := NewTm1027Dto(signNum, "1")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.MemberBasicInfo
	return
}

func (x *Yst2Ka) GetEnterpriseInfo(ctx context.Context, signNum string) (info Tm1027EnterpriseInfo, err error) {
	var r Tm1027Result[Tm1027EnterpriseInfo]
	dto := NewTm1027Dto(signNum, "1")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.MemberBasicInfo
	return
}

func (x *Yst2Ka) GetAcctInfos(ctx context.Context, signNum string) (infos []Tm1027AcctInfo, err error) {
	var r Tm1027Result[map[string]any]
	dto := NewTm1027Dto(signNum, "2")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	infos = r.AcctInfo
	return
}

func (x *Yst2Ka) GetAgreements(ctx context.Context, signNum string) (agreements []Tm1027Agreement, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "3")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	agreements = r.AgreementArray
	return
}

func (x *Yst2Ka) GetOcrResultJson(ctx context.Context, signNum string) (ocr Tm1027OcrResultJson, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "4")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	ocr = r.OcrResultJson
	return
}

func (x *Yst2Ka) GetBindPhoneJson(ctx context.Context, signNum string) (bindPhone Tm1027BindPhoneJson, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "5")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	bindPhone = r.BindPhoneJson
	return
}

func (x *Yst2Ka) GetPayAcctOpenJson(ctx context.Context, signNum string) (info Tm1027PayAcctOpenJson, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "6")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.PayAcctOpenJson
	return
}

func (x *Yst2Ka) GetPayAcctAuditJson(ctx context.Context, signNum string) (info Tm1027PayAcctAuditJson, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "7")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.PayAcctAuditJson
	return
}

func (x *Yst2Ka) GetbankSubAcctInfo(ctx context.Context, signNum string) (info M, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "8")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.BankSubAcctInfo
	return
}

func (x *Yst2Ka) GetSettleAcctInfo(ctx context.Context, signNum string) (info Tm1027SettleAcctInfo, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "9")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.SettleAcctInfo
	return
}

func (x *Yst2Ka) GetMemberControlInfo(ctx context.Context, signNum string) (info Tm1027MemberControlInfo, err error) {
	var r Tm1027Result[M]
	dto := NewTm1027Dto(signNum, "10")
	if err = x.Tm1027(ctx, dto, &r); err != nil {
		return
	}
	info = r.MemberControlInfo
	return
}
