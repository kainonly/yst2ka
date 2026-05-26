package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2089Dto struct {
	ReqTraceNum     string               `json:"reqTraceNum"`               // 商户订单号
	SignNum         string               `json:"signNum,omitempty"`         // 商户会员编号-付款人
	ReceiverList    []Tx2089ReceiverList `json:"receiverList"`              // 收款人列表
	GoodsType       string               `json:"goodsType,omitempty"`       // 商品类型
	BizGoodsNo      string               `json:"bizGoodsNo,omitempty"`      // 商户商品编号
	OrderAmount     int64                `json:"orderAmount"`               // 订单金额
	PayAmount       int64                `json:"payAmount,omitempty"`       // 支付金额
	PromotionAmount int64                `json:"promotionAmount,omitempty"` // 营销金额
	ReqsUrl         string               `json:"reqsUrl,omitempty"`         // 前台通知地址
	RespUrl         string               `json:"respUrl,omitempty"`         // 后台通知地址
	OrderValidTime  string               `json:"orderValidTime,omitempty"`  // 订单过期时间
	PayMode         PayMode              `json:"payMode"`                   // 支付模式
	GoodsName       string               `json:"goodsName,omitempty"`       // 商品名称
	Summary         string               `json:"summary,omitempty"`         // 摘要
	ExtendParams    string               `json:"extendParams,omitempty"`    // 扩展参数
	TxDistrictCode  string               `json:"txDistrictCode,omitempty"`  // 交易所在省市
	GoodsDesc       string               `json:"goodsDesc,omitempty"`       // 商品描述
}

func NewTx2089Dto(reqTraceNum string, receiverList []Tx2089ReceiverList, orderAmount int64, payMode PayMode) *Tx2089Dto {
	return &Tx2089Dto{
		ReqTraceNum:  reqTraceNum,
		ReceiverList: receiverList,
		OrderAmount:  orderAmount,
		PayMode:      payMode,
	}
}

func (x *Tx2089Dto) SetSignNum(v string) *Tx2089Dto {
	x.SignNum = v
	return x
}

func (x *Tx2089Dto) SetGoodsType(v string) *Tx2089Dto {
	x.GoodsType = v
	return x
}

func (x *Tx2089Dto) SetBizGoodsNo(v string) *Tx2089Dto {
	x.BizGoodsNo = v
	return x
}

func (x *Tx2089Dto) SetPayAmount(v int64) *Tx2089Dto {
	x.PayAmount = v
	return x
}

func (x *Tx2089Dto) SetPromotionAmount(v int64) *Tx2089Dto {
	x.PromotionAmount = v
	return x
}

func (x *Tx2089Dto) SetReqsUrl(v string) *Tx2089Dto {
	x.ReqsUrl = v
	return x
}

func (x *Tx2089Dto) SetRespUrl(v string) *Tx2089Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2089Dto) SetOrderValidTime(v string) *Tx2089Dto {
	x.OrderValidTime = v
	return x
}

func (x *Tx2089Dto) SetPayMode(v PayMode) *Tx2089Dto {
	x.PayMode = v
	return x
}

func (x *Tx2089Dto) SetGoodsName(v string) *Tx2089Dto {
	x.GoodsName = v
	return x
}

func (x *Tx2089Dto) SetSummary(v string) *Tx2089Dto {
	x.Summary = v
	return x
}

func (x *Tx2089Dto) SetExtendParams(v string) *Tx2089Dto {
	x.ExtendParams = v
	return x
}

func (x *Tx2089Dto) SetTxDistrictCode(v string) *Tx2089Dto {
	x.TxDistrictCode = v
	return x
}

func (x *Tx2089Dto) SetGoodsDesc(v string) *Tx2089Dto {
	x.GoodsDesc = v
	return x
}

type Tx2089ReceiverList struct {
	SignNum string `json:"signNum"` // 商户会员编号
	Amount  int64  `json:"amount"`  // 收款金额
}

func NewTx2089ReceiverList(signNum string, amount int64) *Tx2089ReceiverList {
	return &Tx2089ReceiverList{
		SignNum: signNum,
		Amount:  amount,
	}
}

type Tx2089ChannelParamInfo map[string]any

type Tx2089ChnlFrontParamInfo map[string]any

type Tx2089Result struct {
	Result             string                   `json:"result,omitempty"`             // 订单状态
	RespTraceNum       string                   `json:"respTraceNum"`                 // 通联订单号
	ReqTraceNum        string                   `json:"reqTraceNum"`                  // 商户订单号
	ExtendParams       string                   `json:"extendParams,omitempty"`       // 扩展参数
	ChannelParamInfo   Tx2089ChannelParamInfo   `json:"channelParamInfo,omitempty"`   // 渠道参数信息（支付详情）
	ChnlFrontParamInfo Tx2089ChnlFrontParamInfo `json:"chnlFrontParamInfo,omitempty"` // 渠道参数信息（前端支付参数）
	RespCode           string                   `json:"respCode"`                     // 业务返回码
	RespMsg            string                   `json:"respMsg"`                      // 业务返回说明
	IsPreConsume       string                   `json:"isPreConsume,omitempty"`       // 是否微信订单预消费
}

func (x *Yst2Ka) Tx2089(ctx context.Context, dto *Tx2089Dto) (_ *Tx2089Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2089`, data); err != nil {
		return
	}

	var result Tx2089Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
