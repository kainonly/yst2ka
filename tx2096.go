package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2096Dto struct {
	ReceiverSignNum    string `json:"receiverSignNum"`              // 充值会员编号
	ReqTraceNum        string `json:"reqTraceNum"`                  // 商户订单号
	OrderAmount        int64  `json:"orderAmount"`                  // 订单金额
	PlatAcctType       string `json:"platAcctType,omitempty"`       // 平台账户类型
	SignNum            string `json:"signNum,omitempty"`            // 付款人会员编号
	PayAmount          int64  `json:"payAmount,omitempty"`          // 支付金额
	PromotionAmount    int64  `json:"promotionAmount,omitempty"`    // 营销金额
	CouponAmount       int64  `json:"couponAmount,omitempty"`       // 抽佣金额
	IsHandleChannelFee string `json:"isHandleChannelFee,omitempty"` // 是否处理渠道手续费
	PayMode            M      `json:"payMode,omitempty"`            // 支付模式
	ReqsUrl            string `json:"reqsUrl,omitempty"`            // 前台通知地址
	RespUrl            string `json:"respUrl,omitempty"`            // 后台通知地址
	OrderValidTime     string `json:"orderValidTime,omitempty"`     // 订单过期时间
	GoodsName          string `json:"goodsName,omitempty"`          // 商品名称
	GoodsDesc          string `json:"goodsDesc,omitempty"`          // 商品描述
	TxDistrictCode     string `json:"txDistrictCode,omitempty"`     // 交易所在省市
	Summary            string `json:"summary,omitempty"`            // 摘要
	ExtendParams       string `json:"extendParams,omitempty"`       // 扩展参数
}

func NewTx2096Dto(receiverSignNum string, reqTraceNum string, orderAmount int64) *Tx2096Dto {
	return &Tx2096Dto{
		ReceiverSignNum: receiverSignNum,
		ReqTraceNum:     reqTraceNum,
		OrderAmount:     orderAmount,
	}
}

func (x *Tx2096Dto) SetPlatAcctType(v string) *Tx2096Dto {
	x.PlatAcctType = v
	return x
}

func (x *Tx2096Dto) SetSignNum(v string) *Tx2096Dto {
	x.SignNum = v
	return x
}

func (x *Tx2096Dto) SetPayAmount(v int64) *Tx2096Dto {
	x.PayAmount = v
	return x
}

func (x *Tx2096Dto) SetPromotionAmount(v int64) *Tx2096Dto {
	x.PromotionAmount = v
	return x
}

func (x *Tx2096Dto) SetCouponAmount(v int64) *Tx2096Dto {
	x.CouponAmount = v
	return x
}

func (x *Tx2096Dto) SetIsHandleChannelFee(v string) *Tx2096Dto {
	x.IsHandleChannelFee = v
	return x
}

func (x *Tx2096Dto) SetPayMode(v M) *Tx2096Dto {
	x.PayMode = v
	return x
}

func (x *Tx2096Dto) SetReqsUrl(v string) *Tx2096Dto {
	x.ReqsUrl = v
	return x
}

func (x *Tx2096Dto) SetRespUrl(v string) *Tx2096Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2096Dto) SetOrderValidTime(v string) *Tx2096Dto {
	x.OrderValidTime = v
	return x
}

func (x *Tx2096Dto) SetGoodsName(v string) *Tx2096Dto {
	x.GoodsName = v
	return x
}

func (x *Tx2096Dto) SetGoodsDesc(v string) *Tx2096Dto {
	x.GoodsDesc = v
	return x
}

func (x *Tx2096Dto) SetTxDistrictCode(v string) *Tx2096Dto {
	x.TxDistrictCode = v
	return x
}

func (x *Tx2096Dto) SetSummary(v string) *Tx2096Dto {
	x.Summary = v
	return x
}

func (x *Tx2096Dto) SetExtendParams(v string) *Tx2096Dto {
	x.ExtendParams = v
	return x
}

type Tx2096ChannelParamInfo map[string]any

type Tx2096ChnlFrontParamInfo map[string]any

type Tx2096Result struct {
	Result             string                   `json:"result,omitempty"`             // 订单状态
	RespTraceNum       string                   `json:"respTraceNum"`                 // 通联订单号
	ReqTraceNum        string                   `json:"reqTraceNum"`                  // 商户订单号
	ExtendParams       string                   `json:"extendParams,omitempty"`       // 扩展参数
	ChannelParamInfo   Tx2096ChannelParamInfo   `json:"channelParamInfo,omitempty"`   // 渠道参数信息（支付详情）
	ChnlFrontParamInfo Tx2096ChnlFrontParamInfo `json:"chnlFrontParamInfo,omitempty"` // 渠道参数信息（前端支付参数）
	RespCode           string                   `json:"respCode"`                     // 业务返回码
	RespMsg            string                   `json:"respMsg"`                      // 业务返回说明
	IsPreConsume       string                   `json:"isPreConsume,omitempty"`       // 是否微信订单预消费
}

func (x *Yst2Ka) Tx2096(ctx context.Context, dto *Tx2096Dto) (_ *Tx2096Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2096`, data); err != nil {
		return
	}

	var result Tx2096Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
