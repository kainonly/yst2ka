package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1035Dto struct {
	ReqTraceNum        string                   `json:"reqTraceNum"`        // 商户请求流水号
	SignNum            string                   `json:"signNum"`            // 商户会员编号
	NotifyUrl          string                   `json:"notifyUrl"`          // 企业会员信息修改结果通知地址
	EnterpriseBaseInfo Tm1035EnterpriseBaseInfo `json:"enterpriseBaseInfo"` // 企业基础信息
}

type Tm1035EnterpriseBaseInfo struct {
	EnterpriseName            string `json:"enterpriseName"`            // 企业名称（含括号请使用中文括号）
	AddressCode               string `json:"addressCode"`               // 地区码（注册地址）
	EnterpriseAdress          string `json:"enterpriseAdress"`          // 企业注册地址
	BusLicenseValidate        string `json:"busLicenseValidate"`        // 营业证件有效期
	LegalPersonName           string `json:"legalPersonName"`           // 法人姓名
	LegalPersonCerType        string `json:"legalPersonCerType"`        // 法人证件类型
	LegalPersonCerNum         string `json:"legalPersonCerNum"`         // 法人证件号码（SM4加密）
	IDValidateStart           string `json:"idValidateStart"`           // 法人证件有效期开始日期，格式：9999-12-31
	IDValidateEnd             string `json:"idValidateEnd"`             // 法人证件有效截止日期，长期有效上送 9999-12-31
	LegalPersonPhone          string `json:"legalPersonPhone"`          // 法人手机号（需与绑卡手机号一致）
	PublicAcctName            string `json:"publicAcctName"`            // 对公户名
	LegpCerFrontFileID        string `json:"legpCerFrontFileId"`        // 法人身份证（肖像面）文件ID
	LegpCerBackFileID         string `json:"legpCerBackFileId"`         // 法人身份证（国徽面）文件ID
	UnifiedSocialCreditFileID string `json:"unifiedSocialCreditFileId"` // 统一信用证文件ID
}

func NewTm1035Dto(reqTraceNum string, signNum string, notifyUrl string, enterpriseBaseInfo Tm1035EnterpriseBaseInfo) *Tm1035Dto {
	return &Tm1035Dto{
		ReqTraceNum:        reqTraceNum,
		SignNum:            signNum,
		NotifyUrl:          notifyUrl,
		EnterpriseBaseInfo: enterpriseBaseInfo,
	}
}

type Tm1035Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号（业务正常处理返回）
	SignNum      string `json:"signNum"`      // 商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 失败原因
}

func (x *Yst2Ka) Tm1035(ctx context.Context, dto *Tm1035Dto) (_ *Tm1035Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1035`, data); err != nil {
		return
	}

	var result Tm1035Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
