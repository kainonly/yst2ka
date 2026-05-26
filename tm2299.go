package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm2299Dto struct {
	ReqTraceNum string `json:"reqTraceNum"`       // 商户请求流水号
	OpenBankNo  string `json:"openBankNo"`        // 开户银行编号
	OrderAmount int64  `json:"orderAmount"`       // 调拨金额，单位分
	CusID       string `json:"cusId"`             // 收银宝商户号
	RespURL     string `json:"respUrl,omitempty"` // 调拨结果通知地址
}

func NewTm2299Dto(reqTraceNum string, openBankNo string, orderAmount int64, cusID string) *Tm2299Dto {
	return &Tm2299Dto{
		ReqTraceNum: reqTraceNum,
		OpenBankNo:  openBankNo,
		OrderAmount: orderAmount,
		CusID:       cusID,
	}
}

func (x *Tm2299Dto) SetRespURL(v string) *Tm2299Dto {
	x.RespURL = v
	return x
}

type Tm2299Result struct {
	RespCode     string `json:"respCode"`         // 业务返回码
	RespMsg      string `json:"respMsg"`          // 业务返回说明
	RespTraceNum string `json:"respTraceNum"`     // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`      // 商户订单号
	Result       string `json:"result,omitempty"` // 订单状态
}

func (x *Yst2Ka) Tm2299(ctx context.Context, dto *Tm2299Dto) (_ *Tm2299Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `2299`, data); err != nil {
		return
	}

	var result Tm2299Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
