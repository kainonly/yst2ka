package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3007Dto struct {
	BatchNo string `json:"batchNo"` // 批次号
}

func NewTq3007Dto(batchNo string) *Tq3007Dto {
	return &Tq3007Dto{BatchNo: batchNo}
}

type Tq3007Result struct {
	BatchNo    string             `json:"batchNo"`              // 批次号
	ResultList []Tq3007ResultItem `json:"resultList,omitempty"` // 处理结果列表
	Status     string             `json:"status"`               // 批次处理状态
	RespCode   string             `json:"respCode"`             // 业务返回码
	RespMsg    string             `json:"respMsg,omitempty"`    // 业务返回说明
}

type Tq3007ResultItem struct {
	RespTraceNum string `json:"respTraceNum"`       // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`        // 商户订单号
	OrderAmount  string `json:"orderAmount"`        // 转账金额，单位分
	Result       string `json:"result,omitempty"`   // 订单状态
	ErrorMsg     string `json:"errorMsg,omitempty"` // 订单失败原因说明
	Summary      string `json:"summary,omitempty"`  // 摘要
}

func (x *Yst2Ka) Tq3007(ctx context.Context, dto *Tq3007Dto) (_ *Tq3007Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tq/handle`, `3007`, data); err != nil {
		return
	}

	var result Tq3007Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
