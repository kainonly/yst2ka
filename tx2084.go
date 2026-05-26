package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2084Dto struct {
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户订单号
	SignNum      string `json:"signNum"`                // 商户会员编号-转出方
	InSignNum    string `json:"inSignNum"`              // 商户会员编号-转入方
	OrderAmount  int64  `json:"orderAmount"`            // 转账金额
	AcctType     string `json:"acctType,omitempty"`     // 转出账户类型
	AcctNum      string `json:"acctNum,omitempty"`      // 支付账户号-转出方
	InAcctType   string `json:"inAcctType,omitempty"`   // 支付账户类型-转入方
	InAcctNum    string `json:"inAcctNum,omitempty"`    // 支付账户号-转入方
	RespUrl      string `json:"respUrl,omitempty"`      // 后台通知地址
	Summary      string `json:"summary,omitempty"`      // 摘要（透传渠道）
	ExtendParams string `json:"extendParams,omitempty"` // 扩展信息
}

func NewTx2084Dto(reqTraceNum string, signNum string, inSignNum string, orderAmount int64) *Tx2084Dto {
	return &Tx2084Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		InSignNum:   inSignNum,
		OrderAmount: orderAmount,
	}
}

func (x *Tx2084Dto) SetAcctType(v string) *Tx2084Dto {
	x.AcctType = v
	return x
}

func (x *Tx2084Dto) SetAcctNum(v string) *Tx2084Dto {
	x.AcctNum = v
	return x
}

func (x *Tx2084Dto) SetInAcctType(v string) *Tx2084Dto {
	x.InAcctType = v
	return x
}

func (x *Tx2084Dto) SetInAcctNum(v string) *Tx2084Dto {
	x.InAcctNum = v
	return x
}

func (x *Tx2084Dto) SetRespUrl(v string) *Tx2084Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2084Dto) SetSummary(v string) *Tx2084Dto {
	x.Summary = v
	return x
}

func (x *Tx2084Dto) SetExtendParams(v string) *Tx2084Dto {
	x.ExtendParams = v
	return x
}

type Tx2084Result struct {
	Result       string `json:"result,omitempty"`       // 订单状态
	AuthWay      string `json:"authWay,omitempty"`      // 鉴权方式
	RespTraceNum string `json:"respTraceNum"`           // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户订单号
	OrderAmount  int64  `json:"orderAmount"`            // 转账金额
	ExtendParams string `json:"extendParams,omitempty"` // 扩展信息
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
}

func (x *Yst2Ka) Tx2084(ctx context.Context, dto *Tx2084Dto) (_ *Tx2084Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2084`, data); err != nil {
		return
	}

	var result Tx2084Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
