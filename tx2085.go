package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2085Dto struct {
	ReqTraceNum     string            `json:"reqTraceNum"`               // 商户订单号
	ReceiverSignNum string            `json:"receiverSignNum"`           // 商户会员编号-收款人
	OrderAmount     int64             `json:"orderAmount"`               // 订单金额
	SignNum         string            `json:"signNum,omitempty"`         // 商户会员编号-付款人
	PayAmount       int64             `json:"payAmount,omitempty"`       // 支付金额
	PromotionAmount int64             `json:"promotionAmount,omitempty"` // 营销金额
	CouponAmount    int64             `json:"couponAmount,omitempty"`    // 平台抽佣金额
	PayMode         M                 `json:"payMode,omitempty"`         // 支付模式
	SepDetailSource string            `json:"sepDetailSource,omitempty"` // 分账规则来源
	SepDetail       []Tx2085SepDetail `json:"sepDetail,omitempty"`       // 分账规则
	ReqsUrl         string            `json:"reqsUrl,omitempty"`         // 前台通知地址
	RespUrl         string            `json:"respUrl,omitempty"`         // 后台通知地址
	OrderValidTime  string            `json:"orderValidTime,omitempty"`  // 订单过期时间
	GoodsName       string            `json:"goodsName,omitempty"`       // 商品名称
	ExtendParams    string            `json:"extendParams,omitempty"`    // 扩展参数
	TxDistrictCode  string            `json:"txDistrictCode,omitempty"`  // 交易所在省市
	GoodsDesc       string            `json:"goodsDesc,omitempty"`       // 商品描述
}

func NewTx2085Dto(reqTraceNum string, receiverSignNum string, orderAmount int64) *Tx2085Dto {
	return &Tx2085Dto{
		ReqTraceNum:     reqTraceNum,
		ReceiverSignNum: receiverSignNum,
		OrderAmount:     orderAmount,
	}
}

func (x *Tx2085Dto) SetSignNum(v string) *Tx2085Dto {
	x.SignNum = v
	return x
}

func (x *Tx2085Dto) SetPayAmount(v int64) *Tx2085Dto {
	x.PayAmount = v
	return x
}

func (x *Tx2085Dto) SetPromotionAmount(v int64) *Tx2085Dto {
	x.PromotionAmount = v
	return x
}

func (x *Tx2085Dto) SetCouponAmount(v int64) *Tx2085Dto {
	x.CouponAmount = v
	return x
}

func (x *Tx2085Dto) SetPayMode(v M) *Tx2085Dto {
	x.PayMode = v
	return x
}

func (x *Tx2085Dto) SetSepDetailSource(v string) *Tx2085Dto {
	x.SepDetailSource = v
	return x
}

func (x *Tx2085Dto) SetSepDetail(v []Tx2085SepDetail) *Tx2085Dto {
	x.SepDetail = v
	return x
}

func (x *Tx2085Dto) SetReqsUrl(v string) *Tx2085Dto {
	x.ReqsUrl = v
	return x
}

func (x *Tx2085Dto) SetRespUrl(v string) *Tx2085Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2085Dto) SetOrderValidTime(v string) *Tx2085Dto {
	x.OrderValidTime = v
	return x
}

func (x *Tx2085Dto) SetGoodsName(v string) *Tx2085Dto {
	x.GoodsName = v
	return x
}

func (x *Tx2085Dto) SetExtendParams(v string) *Tx2085Dto {
	x.ExtendParams = v
	return x
}

func (x *Tx2085Dto) SetTxDistrictCode(v string) *Tx2085Dto {
	x.TxDistrictCode = v
	return x
}

func (x *Tx2085Dto) SetGoodsDesc(v string) *Tx2085Dto {
	x.GoodsDesc = v
	return x
}

type Tx2085SepDetail struct {
	SignNum string `json:"signNum"`          // 商户会员编号
	Amount  int64  `json:"amount"`           // 分账金额
	Remark  string `json:"remark,omitempty"` // 备注
}

func NewTx2085SepDetail(signNum string, amount int64) *Tx2085SepDetail {
	return &Tx2085SepDetail{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2085SepDetail) SetRemark(v string) *Tx2085SepDetail {
	x.Remark = v
	return x
}

type Tx2085ChannelParamInfo map[string]any

type Tx2085ChnlFrontParamInfo map[string]any

type Tx2085Result struct {
	Result             string `json:"result,omitempty"`             // 订单状态
	RespTraceNum       string `json:"respTraceNum"`                 // 通联订单号
	ReqTraceNum        string `json:"reqTraceNum"`                  // 商户订单号
	ExtendParams       string `json:"extendParams,omitempty"`       // 扩展参数
	ChannelParamInfo   string `json:"channelParamInfo,omitempty"`   // 渠道参数信息（支付详情）
	ChnlFrontParamInfo string `json:"chnlFrontParamInfo,omitempty"` // 渠道参数信息（前端支付参数）
	RespCode           string `json:"respCode"`                     // 业务返回码
	RespMsg            string `json:"respMsg"`                      // 业务返回说明
	IsPreConsume       string `json:"isPreConsume,omitempty"`       // 是否微信订单预消费
}

func (x *Yst2Ka) Tx2085(ctx context.Context, dto *Tx2085Dto) (_ *Tx2085Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tx/handle`, `2085`, data); err != nil {
		return
	}

	var result Tx2085Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
