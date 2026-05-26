package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx3010Dto struct {
	VerifyCode   string `json:"verifyCode"`             // 短信验证码
	BatchNo      string `json:"batchNo,omitempty"`      // 批次号
	RespTraceNum string `json:"respTraceNum,omitempty"` // 通联订单号
}

func NewTx3010Dto(verifyCode string) *Tx3010Dto {
	return &Tx3010Dto{
		VerifyCode: verifyCode,
	}
}

func (x *Tx3010Dto) SetBatchNo(v string) *Tx3010Dto {
	x.BatchNo = v
	return x
}

func (x *Tx3010Dto) SetRespTraceNum(v string) *Tx3010Dto {
	x.RespTraceNum = v
	return x
}

type Tx3010Result struct {
	Result       string `json:"result,omitempty"`       // 订单状态
	BatchNo      string `json:"batchNo,omitempty"`      // 批次号
	ReqTraceNum  string `json:"reqTraceNum,omitempty"`  // 商户订单号
	RespTraceNum string `json:"respTraceNum,omitempty"` // 通联订单号
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
}

func (x *Yst2Ka) Tx3010(ctx context.Context, dto *Tx3010Dto) (_ *Tx3010Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `3010`, data); err != nil {
		return
	}

	var result Tx3010Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
