package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2290Dto struct {
	SignNum         string  `json:"signNum"`                   // 商户会员编号
	ReqTraceNum     string  `json:"reqTraceNum"`               // 商户订单号
	OrderAmount     int64   `json:"orderAmount"`               // 订单金额
	AcctNum         string  `json:"acctNum"`                   // 银行卡号
	AcctType        string  `json:"acctType,omitempty"`        // 提现账户类型
	PayAcctNo       string  `json:"payAcctNo,omitempty"`       // 支付账户号
	CouponAmount    int64   `json:"couponAmount,omitempty"`    // 平台抽佣金额
	RespUrl         string  `json:"respUrl,omitempty"`         // 后台通知地址
	PayMode         PayMode `json:"payMode,omitempty"`         // 支付模式
	ReceiveAcctType string  `json:"receiveAcctType,omitempty"` // 入账账户类型
	WithdrawType    string  `json:"withdrawType,omitempty"`    // 提现方式
	Summary         string  `json:"summary,omitempty"`         // 摘要
	ExtendParams    string  `json:"extendParams,omitempty"`    // 扩展信息
}

func NewTx2290Dto(signNum string, reqTraceNum string, orderAmount int64, acctNum string) *Tx2290Dto {
	return &Tx2290Dto{
		SignNum:     signNum,
		ReqTraceNum: reqTraceNum,
		OrderAmount: orderAmount,
		AcctNum:     acctNum,
	}
}

func (x *Tx2290Dto) SetAcctType(v string) *Tx2290Dto {
	x.AcctType = v
	return x
}

func (x *Tx2290Dto) SetPayAcctNo(v string) *Tx2290Dto {
	x.PayAcctNo = v
	return x
}

func (x *Tx2290Dto) SetCouponAmount(v int64) *Tx2290Dto {
	x.CouponAmount = v
	return x
}

func (x *Tx2290Dto) SetRespUrl(v string) *Tx2290Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2290Dto) SetPayMode(v PayMode) *Tx2290Dto {
	x.PayMode = v
	return x
}

func (x *Tx2290Dto) SetReceiveAcctType(v string) *Tx2290Dto {
	x.ReceiveAcctType = v
	return x
}

func (x *Tx2290Dto) SetWithdrawType(v string) *Tx2290Dto {
	x.WithdrawType = v
	return x
}

func (x *Tx2290Dto) SetSummary(v string) *Tx2290Dto {
	x.Summary = v
	return x
}

func (x *Tx2290Dto) SetExtendParams(v string) *Tx2290Dto {
	x.ExtendParams = v
	return x
}

type Tx2290Result struct {
	Result        string `json:"result,omitempty"`        // 订单状态
	RespTraceNum  string `json:"respTraceNum"`            // 通联订单号
	ReqTraceNum   string `json:"reqTraceNum"`             // 商户订单号（支付订单）
	ChnlTradeCode string `json:"chnlTradeCode,omitempty"` // 收付通渠道银行流水号
	ExtendParams  string `json:"extendParams,omitempty"`  // 扩展信息
	RespCode      string `json:"respCode"`                // 业务返回码
	RespMsg       string `json:"respMsg"`                 // 业务返回说明
}

func (x *Yst2Ka) Tx2290(ctx context.Context, dto *Tx2290Dto) (_ *Tx2290Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2290`, data); err != nil {
		return
	}

	var result Tx2290Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
