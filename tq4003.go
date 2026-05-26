package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq4003Dto struct {
	BatchNo      string `json:"batchNo,omitempty"`      // 批次号，和 respTraceNum 二选一必填
	RespTraceNum string `json:"respTraceNum,omitempty"` // 通联订单号，和 batchNo 二选一必填
}

func NewTq4003Dto() *Tq4003Dto {
	return &Tq4003Dto{}
}

func (x *Tq4003Dto) SetBatchNo(v string) *Tq4003Dto {
	x.BatchNo = v
	return x
}

func (x *Tq4003Dto) SetRespTraceNum(v string) *Tq4003Dto {
	x.RespTraceNum = v
	return x
}

type Tq4003Result struct {
	FileURL      string `json:"fileUrl,omitempty"`      // 电子回单地址
	RespTraceNum string `json:"respTraceNum,omitempty"` // 通联订单号
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
}

func (x *Yst2Ka) Tq4003(ctx context.Context, dto *Tq4003Dto) (_ *Tq4003Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tq/handle`, `4003`, data); err != nil {
		return
	}

	var result Tq4003Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
