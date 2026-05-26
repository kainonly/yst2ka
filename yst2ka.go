package yst2ka

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/emmansun/gmsm/sm2"
	"github.com/kainonly/go/help"
	"resty.dev/v3"
)

type Yst2Ka struct {
	Option *Option
	Client *resty.Client
	priKey *sm2.PrivateKey
	pubKey *ecdsa.PublicKey
}

type Option struct {
	BaseURL           string `yaml:"base_url" env:"BASE_URL"`
	PrivateKey        string `yaml:"private_key" env:"PRIVATE_KEY"`
	AllinpayPublicKey string `yaml:"allinpay_public_key" env:"ALLINPAY_PUBLIC_KEY"`
	AppID             string `yaml:"app_id" env:"APP_ID"`
	SpAppID           string `json:"sp_app_id" env:"SP_APP_ID"`
	SecretKey         string `yaml:"secret_key" env:"SECRET_KEY"`
}

func NewYst2Ka(opt Option) (x *Yst2Ka, err error) {
	x = &Yst2Ka{
		Option: &opt,
		Client: resty.New().SetBaseURL(opt.BaseURL),
	}
	if x.priKey, err = help.PrivKeySM2FromBase64(opt.PrivateKey); err != nil {
		return
	}
	if x.pubKey, err = help.PubKeySM2FromBase64(opt.AllinpayPublicKey); err != nil {
		return
	}
	return
}

func (x *Yst2Ka) Debug() {
	x.Client.EnableTrace()
}

func (x *Yst2Ka) GetPublicKey() *ecdsa.PublicKey {
	return x.pubKey
}

func (x *Yst2Ka) GetPrivateKey() *sm2.PrivateKey {
	return x.priKey
}

type M = map[string]any

// Shared dictionary types reused across multiple trades.
type PayMode map[string]any

func NewPayMode() PayMode {
	return PayMode{}
}

type MemberType string

const (
	MemberTypeEnterprise MemberType = "2" // 企业会员
	MemberTypePersonal   MemberType = "3" // 个人会员
)

type MemberStatus string

const (
	MemberStatusPendingActivation MemberStatus = "0" // 待生效
	MemberStatusValid             MemberStatus = "1" // 有效
	MemberStatusAuditFailed       MemberStatus = "2" // 审核失败
	MemberStatusRiskBlacklist     MemberStatus = "3" // 风控黑名单
	MemberStatusBlacklist         MemberStatus = "4" // 黑名单
	MemberStatusCanceled          MemberStatus = "5" // 注销
	MemberStatusBankAuditing      MemberStatus = "6" // 银行审核中
	MemberStatusBankAuditFailed   MemberStatus = "7" // 银行审核失败
)

type CerType string

const (
	CerTypeIdentityCard                 CerType = "1"  // 身份证
	CerTypePassport                     CerType = "2"  // 护照
	CerTypeOfficerCertificate           CerType = "3"  // 军官证
	CerTypeHomeReturnPermit             CerType = "4"  // 回乡证
	CerTypeTaiwanCompatriotPermit       CerType = "5"  // 台胞证
	CerTypePoliceCertificate            CerType = "6"  // 警官证
	CerTypeSoldierCertificate           CerType = "7"  // 士兵证
	CerTypeHouseholdBook                CerType = "8"  // 户口簿
	CerTypeHKMacaoMainlandTravelPermit  CerType = "9"  // 港澳居民来往内地通行证
	CerTypeTemporaryIdentityCard        CerType = "10" // 临时身份证
	CerTypeForeignerResidencePermit     CerType = "11" // 外国人居留证
	CerTypeHKMacaoTaiwanResidencePermit CerType = "12" // 港澳台居民居住证
	CerTypeUnifiedSocialCreditCode      CerType = "13" // 统一社会信用代码证
	CerTypeTaiwanMainlandTravelPermit   CerType = "14" // 台湾同胞来往内地通行证
	CerTypeBusinessLicense              CerType = "15" // 营业执照
	CerTypeOrganizationCodeCertificate  CerType = "16" // 组织机构代码
	CerTypeTaxRegistrationCertificate   CerType = "17" // 税务登记证
	CerTypeForeignPassport              CerType = "18" // 外国护照
	CerTypeEntryExitPermit              CerType = "19" // 出入境通行证
	CerTypeHKMacaoResidencePermit       CerType = "20" // 港澳居民居住证
	CerTypeTaiwanResidencePermit        CerType = "21" // 台湾居民居住证
	CerTypeTravelPermit                 CerType = "22" // 旅行证
	CerTypeOther                        CerType = "99" // 其他证件
)

type BindStatus string

const (
	BindStatusBound    BindStatus = "1" // 已绑定
	BindStatusReleased BindStatus = "2" // 已解除
)

type CardType string

const (
	CardTypeDebit  CardType = "0" // 借记卡
	CardTypeCredit CardType = "1" // 信用卡
)

type SignStatus string

const (
	SignStatusUnsigned   SignStatus = "0" // 未签约
	SignStatusSuccess    SignStatus = "1" // 签约成功
	SignStatusFailed     SignStatus = "2" // 签约失败
	SignStatusProcessing SignStatus = "3" // 签约中
)

// 当前数据字典页仅列出 6、7、8、99 四种绑卡方式。
type BindType string

const (
	BindTypeAllinpayTongAgreementPay BindType = "6"  // 通联通协议支付签约
	BindTypeCashierQuickPayAgreement BindType = "7"  // 收银宝快捷支付签约（有银行范围）
	BindTypeBankCardFourFactor       BindType = "8"  // 银行卡四要素验证（全部银行）
	BindTypeBackendBindBankCard      BindType = "99" // 后台绑定银行卡
)

type MerchantType string

const (
	MerchantTypeCompany                        MerchantType = "1"  // 公司（不包含国有企业）
	MerchantTypePartnership                    MerchantType = "2"  // 合伙企业
	MerchantTypeSoleProprietorship             MerchantType = "3"  // 个人独资企业
	MerchantTypeStateOwnedEnterprise           MerchantType = "4"  // 受政府控制的企业（国有企业）
	MerchantTypePublicInstitution              MerchantType = "5"  // 事业单位（不包含参照公务员法管理的事业单位）
	MerchantTypeGovernmentAgency               MerchantType = "6"  // 党政军政协等机关、参公事业单位
	MerchantTypeProfessionalServiceInstitution MerchantType = "7"  // 不具备法人资格的专业服务机构
	MerchantTypeFarmerCooperative              MerchantType = "8"  // 经营农林渔牧产业的非公司制农民专业合作组织
	MerchantTypeSocialOrganization             MerchantType = "9"  // 社会团体、基金会、社会服务机构和外国商会
	MerchantTypeInternationalOrganization      MerchantType = "10" // 政府间国际组织、外国政府驻华使领馆及办事处等机构及组织
)

type AccountType string

const (
	AccountTypeBookkeeping       AccountType = "01" // 簿记账户
	AccountTypeAppMarketing      AccountType = "02" // 应用营销账户
	AccountTypeAppGuarantee      AccountType = "03" // 应用担保账户
	AccountTypePendingSettlement AccountType = "08" // 待结算户
	AccountTypePay               AccountType = "11" // 支付账户
)

type PayAccountStatus string

const (
	PayAccountStatusPendingActivation PayAccountStatus = "0" // 待激活
	PayAccountStatusNormal            PayAccountStatus = "2" // 正常
	PayAccountStatusFrozen            PayAccountStatus = "3" // 冻结
	PayAccountStatusBusinessSuspended PayAccountStatus = "4" // 业务暂停
	PayAccountStatusShutdown          PayAccountStatus = "5" // 关停
	PayAccountStatusCaseFrozen        PayAccountStatus = "6" // 涉案冻结
	PayAccountStatusCanceled          PayAccountStatus = "7" // 注销
)

type OrderStatus int

const (
	OrderStatusProcessing OrderStatus = 0 // 进行中
	OrderStatusSuccess    OrderStatus = 1 // 交易成功
	OrderStatusFailed     OrderStatus = 2 // 交易失败
)

type ChannelTradeType string

const (
	ChannelTradeTypeConsume                   ChannelTradeType = "VSP001" // 消费
	ChannelTradeTypeConsumeCancel             ChannelTradeType = "VSP002" // 消费撤销
	ChannelTradeTypeRefund                    ChannelTradeType = "VSP003" // 退货/消费退货
	ChannelTradeTypePreAuth                   ChannelTradeType = "VSP004" // 预授权
	ChannelTradeTypePreAuthCancel             ChannelTradeType = "VSP005" // 预授权撤销
	ChannelTradeTypePreAuthComplete           ChannelTradeType = "VSP006" // 预授权完成
	ChannelTradeTypePreAuthCompleteCancel     ChannelTradeType = "VSP007" // 预授权完成撤销
	ChannelTradeTypeManualRefundRegistration  ChannelTradeType = "VSP008" // 手工退货登记
	ChannelTradeTypeScanPreConsume            ChannelTradeType = "VSP011" // 扫码预消费
	ChannelTradeTypeScanPreConsumeComplete    ChannelTradeType = "VSP013" // 扫码预消费完成
	ChannelTradeTypeScanPreConsumeRefund      ChannelTradeType = "VSP014" // 扫码预消费完成退货
	ChannelTradeTypeConsumeReversal           ChannelTradeType = "CMN001" // 消费冲正
	ChannelTradeTypePreAuthReversal           ChannelTradeType = "CMN002" // 预授权冲正
	ChannelTradeTypeQuickPay                  ChannelTradeType = "VSP301" // 快捷支付
	ChannelTradeTypeQuickPayCancel            ChannelTradeType = "VSP302" // 快捷支付撤销
	ChannelTradeTypeQuickPayRefund            ChannelTradeType = "VSP303" // 快捷支付退货
	ChannelTradeTypeWeChatPay                 ChannelTradeType = "VSP501" // 微信支付
	ChannelTradeTypeWeChatCancel              ChannelTradeType = "VSP502" // 微信撤销
	ChannelTradeTypeWeChatRefund              ChannelTradeType = "VSP503" // 微信退款
	ChannelTradeTypeAlipayPay                 ChannelTradeType = "VSP511" // 支付宝支付
	ChannelTradeTypeAlipayCancel              ChannelTradeType = "VSP512" // 支付宝撤销
	ChannelTradeTypeAlipayRefund              ChannelTradeType = "VSP513" // 支付宝退货
	ChannelTradeTypeGatewayPay                ChannelTradeType = "VSP531" // 网关支付
	ChannelTradeTypeGatewayCancel             ChannelTradeType = "VSP532" // 网关撤销
	ChannelTradeTypeGatewayRefund             ChannelTradeType = "VSP533" // 网关退货
	ChannelTradeTypeGatewayPayB2B             ChannelTradeType = "VSP534" // 网关支付(B2B)
	ChannelTradeTypeGatewayPayApp             ChannelTradeType = "VSP535" // 网关支付(APP)
	ChannelTradeTypeScanPayRefund             ChannelTradeType = "VSP543" // 扫码支付退货
	ChannelTradeTypeUnionPayQRCodePay         ChannelTradeType = "VSP551" // 银联扫码支付
	ChannelTradeTypeUnionPayQRCodeCancel      ChannelTradeType = "VSP552" // 银联扫码撤销
	ChannelTradeTypeUnionPayQRCodeRefund      ChannelTradeType = "VSP553" // 银联扫码退货
	ChannelTradeTypeCloudQuickPassApp         ChannelTradeType = "VSP591" // 云闪付APP
	ChannelTradeTypeCloudQuickPassCancel      ChannelTradeType = "VSP592" // 云闪付撤销
	ChannelTradeTypeCloudQuickPassRefund      ChannelTradeType = "VSP593" // 云闪付退货
	ChannelTradeTypeInstallmentPay            ChannelTradeType = "VSP631" // 聚分期支付
	ChannelTradeTypeInstallmentCancel         ChannelTradeType = "VSP632" // 聚分期撤销
	ChannelTradeTypeInstallmentRefund         ChannelTradeType = "VSP633" // 聚分期退货
	ChannelTradeTypeWeChatOrderPreConsume     ChannelTradeType = "VSP681" // 微信订单预消费
	ChannelTradeTypeWeChatOrderRefund         ChannelTradeType = "VSP682" // 微信订单退款
	ChannelTradeTypeWeChatOrderComplete       ChannelTradeType = "VSP683" // 微信订单完成
	ChannelTradeTypeWeChatOrderCompleteRefund ChannelTradeType = "VSP684" // 微信订单完成退款
)

func (x *Yst2Ka) SetNow(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, "now", ts)
}

func (x *Yst2Ka) GetNow(ctx context.Context) time.Time {
	return ctx.Value("now").(time.Time)
}

type ResponseBody struct {
	Code    string `json:"code"`    // 调用结果返回码
	Msg     string `json:"msg"`     // 调用结果返回码描述
	Sign    string `json:"sign"`    // 商户请求参数的签名串
	BizData string `json:"bizData"` // 返回参数的集合
}

func (x *Yst2Ka) Request(ctx context.Context, path string, code string, data string) (_ string, err error) {
	now := x.GetNow(ctx)
	body := M{
		"appId":     x.Option.AppID,
		"bizData":   data,
		"charset":   "UTF-8",
		"format":    "json",
		"transCode": code,
		"transDate": now.Format(`20060102`),
		"transTime": now.Format(`150405`),
		"version":   "1.0",
	}

	var signature string
	if signature, err = help.Sm2Sign(x.priKey, help.MapToSignText(body)); err != nil {
		return
	}
	body["sign"] = signature
	body["signType"] = "SM3withSM2"

	var resp *resty.Response
	if resp, err = x.Client.R().
		SetContext(ctx).
		SetBody(body).
		Post(path); err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = help.E(0, `第三方接口响应失败!`)
		return
	}
	var content M
	if err = sonic.Unmarshal(resp.Bytes(), &content); err != nil {
		return
	}

	if content["code"] != "00000" {
		err = help.E(0, fmt.Sprintf(`第三方请求失败![%s]: %s`, content["code"], content["msg"]))
		return
	}

	sign := content["sign"].(string)
	delete(content, "sign")
	delete(content, "signType")

	var verify bool
	if verify, err = help.Sm2Verify(x.pubKey, help.MapToSignText(content), sign); err != nil {
		return
	}
	if !verify {
		err = help.E(0, `第三方响应内容签名存在不一致!`)
		return
	}

	return content["bizData"].(string), nil
}
