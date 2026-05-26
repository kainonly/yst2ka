package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2293Dto struct {
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户请求流水号
	OpenBankNo   string `json:"openBankNo"`             // 开户银行编号
	OrderAmount  int64  `json:"orderAmount"`            // 调拨资金，单位分
	RespURL      string `json:"respUrl,omitempty"`      // 调拨结果通知地址
	ExtendParams string `json:"extendParams,omitempty"` // 备注
}

func NewTx2293Dto(reqTraceNum string, openBankNo string, orderAmount int64) *Tx2293Dto {
	return &Tx2293Dto{
		ReqTraceNum: reqTraceNum,
		OpenBankNo:  openBankNo,
		OrderAmount: orderAmount,
	}
}

func (x *Tx2293Dto) SetRespURL(v string) *Tx2293Dto {
	x.RespURL = v
	return x
}

func (x *Tx2293Dto) SetExtendParams(v string) *Tx2293Dto {
	x.ExtendParams = v
	return x
}

type Tx2293Result struct {
	Result       string `json:"result"`       // 订单状态，0-进行中 1-交易成功 2-交易失败
	RespTraceNum string `json:"respTraceNum"` // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`  // 商户订单号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tx2293(ctx context.Context, dto *Tx2293Dto) (_ *Tx2293Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2293`, data); err != nil {
		return
	}

	var result Tx2293Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
