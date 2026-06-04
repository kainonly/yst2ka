package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2294Dto struct {
	ReqTraceNum              string                          `json:"reqTraceNum"`                        // 商户订单号-退款订单号
	OrderAmount              int64                           `json:"orderAmount"`                        // 退款总金额
	OrgReqTraceNum           string                          `json:"orgReqTraceNum,omitempty"`           // 原商户订单号
	OrgTransDate             string                          `json:"orgTransDate,omitempty"`             // 原订单创建日期
	OrgRespTraceNum          string                          `json:"orgRespTraceNum,omitempty"`          // 原通联订单号
	PromotionAmount          int64                           `json:"promotionAmount,omitempty"`          // 营销退款金额
	RefundDetail             []Tx2294RefundDetail            `json:"refundDetail,omitempty"`             // 订单退款列表
	SvcRefundDetail          []Tx2294SvcRefundDetail         `json:"svcRefundDetail,omitempty"`          // 储值卡核销退款列表
	IsFundAllocation         string                          `json:"isFundAllocation,omitempty"`         // 是否需要调拨资金
	IsAdvancePay             string                          `json:"isAdvancePay,omitempty"`             // 是否允许收款人垫资
	RespUrl                  string                          `json:"respUrl,omitempty"`                  // 后台通知地址
	ChnlDiscAmt              Tx2294ChnlDiscAmt               `json:"chnlDiscAmt,omitempty"`              // 优惠信息
	Summary                  string                          `json:"summary,omitempty"`                  // 摘要
	ExtendParams             string                          `json:"extendParams,omitempty"`             // 扩展信息
	WechatPayB2bRefundDetail *Tx2294WechatPayB2bRefundDetail `json:"wechatPayB2bRefundDetail,omitempty"` // B2b门店助手退款详情
	Remark                   string                          `json:"remark,omitempty"`                   // 业务备注
	Reason                   string                          `json:"reason,omitempty"`                   // 退款理由
}

func NewTx2294Dto(reqTraceNum string, orderAmount int64) *Tx2294Dto {
	return &Tx2294Dto{
		ReqTraceNum: reqTraceNum,
		OrderAmount: orderAmount,
	}
}

func (x *Tx2294Dto) SetOrgReqTraceNum(v string) *Tx2294Dto {
	x.OrgReqTraceNum = v
	return x
}

func (x *Tx2294Dto) SetOrgTransDate(v string) *Tx2294Dto {
	x.OrgTransDate = v
	return x
}

func (x *Tx2294Dto) SetOrgRespTraceNum(v string) *Tx2294Dto {
	x.OrgRespTraceNum = v
	return x
}

func (x *Tx2294Dto) SetPromotionAmount(v int64) *Tx2294Dto {
	x.PromotionAmount = v
	return x
}

func (x *Tx2294Dto) SetRefundDetail(v []Tx2294RefundDetail) *Tx2294Dto {
	x.RefundDetail = v
	return x
}

func (x *Tx2294Dto) SetSvcRefundDetail(v []Tx2294SvcRefundDetail) *Tx2294Dto {
	x.SvcRefundDetail = v
	return x
}

func (x *Tx2294Dto) SetIsFundAllocation(v string) *Tx2294Dto {
	x.IsFundAllocation = v
	return x
}

func (x *Tx2294Dto) SetIsAdvancePay(v string) *Tx2294Dto {
	x.IsAdvancePay = v
	return x
}

func (x *Tx2294Dto) SetRespUrl(v string) *Tx2294Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2294Dto) SetChnlDiscAmt(v Tx2294ChnlDiscAmt) *Tx2294Dto {
	x.ChnlDiscAmt = v
	return x
}

func (x *Tx2294Dto) SetSummary(v string) *Tx2294Dto {
	x.Summary = v
	return x
}

func (x *Tx2294Dto) SetExtendParams(v string) *Tx2294Dto {
	x.ExtendParams = v
	return x
}

func (x *Tx2294Dto) SetWechatPayB2bRefundDetail(v *Tx2294WechatPayB2bRefundDetail) *Tx2294Dto {
	x.WechatPayB2bRefundDetail = v
	return x
}

func (x *Tx2294Dto) SetRemark(v string) *Tx2294Dto {
	x.Remark = v
	return x
}

func (x *Tx2294Dto) SetReason(v string) *Tx2294Dto {
	x.Reason = v
	return x
}

type Tx2294RefundDetail struct {
	SignNum      string            `json:"signNum"`                // 商户会员编号
	OrderAmount  int64             `json:"orderAmount"`            // 退款金额
	AcctType     string            `json:"acctType,omitempty"`     // 账户类型
	CouponAmount int64             `json:"couponAmount,omitempty"` // 平台抽佣退款金额
	SepDetail    []Tx2294SepDetail `json:"sepDetail,omitempty"`    // 分账退款列表
}

func NewTx2294RefundDetail(signNum string, orderAmount int64) *Tx2294RefundDetail {
	return &Tx2294RefundDetail{
		SignNum:     signNum,
		OrderAmount: orderAmount,
	}
}

func (x *Tx2294RefundDetail) SetAcctType(v string) *Tx2294RefundDetail {
	x.AcctType = v
	return x
}

func (x *Tx2294RefundDetail) SetCouponAmount(v int64) *Tx2294RefundDetail {
	x.CouponAmount = v
	return x
}

func (x *Tx2294RefundDetail) SetSepDetail(v []Tx2294SepDetail) *Tx2294RefundDetail {
	x.SepDetail = v
	return x
}

type Tx2294SvcRefundDetail struct {
	SignNum      string            `json:"signNum"`                // 商户会员编号
	OrderAmount  int64             `json:"orderAmount"`            // 储值卡核销退款金额
	CouponAmount int64             `json:"couponAmount,omitempty"` // 平台抽佣退款金额
	SepDetail    []Tx2294SepDetail `json:"sepDetail,omitempty"`    // 分账退款列表
}

func NewTx2294SvcRefundDetail(signNum string, orderAmount int64) *Tx2294SvcRefundDetail {
	return &Tx2294SvcRefundDetail{
		SignNum:     signNum,
		OrderAmount: orderAmount,
	}
}

func (x *Tx2294SvcRefundDetail) SetCouponAmount(v int64) *Tx2294SvcRefundDetail {
	x.CouponAmount = v
	return x
}

func (x *Tx2294SvcRefundDetail) SetSepDetail(v []Tx2294SepDetail) *Tx2294SvcRefundDetail {
	x.SepDetail = v
	return x
}

type Tx2294SepDetail struct {
	SignNum string `json:"signNum"`          // 商户会员编号
	Amount  int64  `json:"amount"`           // 分账退款金额
	Remark  string `json:"remark,omitempty"` // 备注
}

func NewTx2294SepDetail(signNum string, amount int64) *Tx2294SepDetail {
	return &Tx2294SepDetail{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2294SepDetail) SetRemark(v string) *Tx2294SepDetail {
	x.Remark = v
	return x
}

type Tx2294ChnlDiscAmt map[string]any

func NewTx2294ChnlDiscAmt() Tx2294ChnlDiscAmt {
	return Tx2294ChnlDiscAmt{}
}

type Tx2294WechatPayB2bRefundDetail struct {
	RefundFrom   string `json:"refundFrom,omitempty"`   // 退款来源
	RefundReason string `json:"refundReason,omitempty"` // 退款原因
	Description  string `json:"description,omitempty"`  // 退款商品描述
}

func NewTx2294WechatPayB2bRefundDetail() *Tx2294WechatPayB2bRefundDetail {
	return &Tx2294WechatPayB2bRefundDetail{}
}

func (x *Tx2294WechatPayB2bRefundDetail) SetRefundFrom(v string) *Tx2294WechatPayB2bRefundDetail {
	x.RefundFrom = v
	return x
}

func (x *Tx2294WechatPayB2bRefundDetail) SetRefundReason(v string) *Tx2294WechatPayB2bRefundDetail {
	x.RefundReason = v
	return x
}

func (x *Tx2294WechatPayB2bRefundDetail) SetDescription(v string) *Tx2294WechatPayB2bRefundDetail {
	x.Description = v
	return x
}

type Tx2294Result struct {
	Result           string `json:"result,omitempty"`           // 订单状态
	ReqTraceNum      string `json:"reqTraceNum"`                // 商户订单号
	RespTraceNum     string `json:"respTraceNum"`               // 通联订单号
	ExtendParams     string `json:"extendParams,omitempty"`     // 扩展信息
	ChannelParamInfo string `json:"channelParamInfo,omitempty"` // 渠道参数信息
	RespCode         string `json:"respCode"`                   // 业务返回码
	RespMsg          string `json:"respMsg"`                    // 业务返回说明
}

func (x *Yst2Ka) Tx2294(ctx context.Context, dto *Tx2294Dto) (_ *Tx2294Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2294`, data); err != nil {
		return
	}

	var result Tx2294Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
