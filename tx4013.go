package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx4013Dto struct {
	DigID       string         `json:"digId,omitempty"`       // 应用机构ID
	OrgID       string         `json:"orgId,omitempty"`       // 应用机构号
	OrgAppID    string         `json:"orgAppId,omitempty"`    // 机构appid
	OrderNo     string         `json:"orderNo"`               // 通联订单号
	Code        string         `json:"code"`                  // 收银宝分配的二维码编号
	AgreeIDList string         `json:"agreeIdList,omitempty"` // 快捷协议号列表
	MktInfo     *Tx4013MktInfo `json:"mktInfo,omitempty"`     // 营销信息
	ReqsURL     string         `json:"reqsUrl"`               // 支付成功跳转地址
	NoPayType   string         `json:"noPayType,omitempty"`   // 需屏蔽支付类型
}

type Tx4013MktInfo struct {
	OutUserID string `json:"outUserid,omitempty"` // 商户平台会员号
	OpeID     string `json:"opeid,omitempty"`     // 基础营销活动发起方
	MktAppID  string `json:"mktAppid,omitempty"`  // 基础营销appid
}

func NewTx4013Dto(orderNo string, code string, reqsURL string) *Tx4013Dto {
	return &Tx4013Dto{
		OrderNo: orderNo,
		Code:    code,
		ReqsURL: reqsURL,
	}
}

func NewTx4013QuickH5Dto(digID string, orgID string, orgAppID string, orderNo string, code string, reqsURL string) *Tx4013Dto {
	return &Tx4013Dto{
		DigID:    digID,
		OrgID:    orgID,
		OrgAppID: orgAppID,
		OrderNo:  orderNo,
		Code:     code,
		ReqsURL:  reqsURL,
	}
}

func NewTx4013MktInfo() *Tx4013MktInfo {
	return &Tx4013MktInfo{}
}

func (x *Tx4013Dto) SetDigID(v string) *Tx4013Dto {
	x.DigID = v
	return x
}

func (x *Tx4013Dto) SetOrgID(v string) *Tx4013Dto {
	x.OrgID = v
	return x
}

func (x *Tx4013Dto) SetOrgAppID(v string) *Tx4013Dto {
	x.OrgAppID = v
	return x
}

func (x *Tx4013Dto) SetAgreeIDList(v string) *Tx4013Dto {
	x.AgreeIDList = v
	return x
}

func (x *Tx4013Dto) SetMktInfo(v *Tx4013MktInfo) *Tx4013Dto {
	x.MktInfo = v
	return x
}

func (x *Tx4013Dto) SetNoPayType(v string) *Tx4013Dto {
	x.NoPayType = v
	return x
}

func (x *Tx4013MktInfo) SetOutUserID(v string) *Tx4013MktInfo {
	x.OutUserID = v
	return x
}

func (x *Tx4013MktInfo) SetOpeID(v string) *Tx4013MktInfo {
	x.OpeID = v
	return x
}

func (x *Tx4013MktInfo) SetMktAppID(v string) *Tx4013MktInfo {
	x.MktAppID = v
	return x
}

type Tx4013Result struct {
	RespCode           string `json:"respCode"`             // 业务返回码
	RespMsg            string `json:"respMsg"`              // 业务返回说明
	PayAmount          int64  `json:"payAmount"`            // 支付金额
	ChnlFrontParamInfo string `json:"chnlFrontParamInfo"`   // 当面付支付链接
	OrderNo            string `json:"orderNo,omitempty"`    // 通联订单号
	TrxReserve         string `json:"trxReserve,omitempty"` // 业务备注信息
}

func (x *Yst2Ka) Tx4013(ctx context.Context, dto *Tx4013Dto) (_ *Tx4013Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `4013`, data); err != nil {
		return
	}

	var result Tx4013Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
