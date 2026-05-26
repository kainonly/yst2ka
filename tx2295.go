package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2295Dto struct {
	OrgRespTraceNum string `json:"orgRespTraceNum"`       // 原通联订单号
	CloseReason     string `json:"closeReason,omitempty"` // 关单原因
}

func NewTx2295Dto(orgRespTraceNum string) *Tx2295Dto {
	return &Tx2295Dto{
		OrgRespTraceNum: orgRespTraceNum,
	}
}

func (x *Tx2295Dto) SetCloseReason(v string) *Tx2295Dto {
	x.CloseReason = v
	return x
}

type Tx2295Result struct {
	ReqTraceNum     string `json:"reqTraceNum"`               // 商户订单号
	RespTraceNum    string `json:"respTraceNum"`              // 通联订单号
	CloseResult     string `json:"closeResult"`               // 订单关闭结果
	CloseFinishTime string `json:"closeFinishTime,omitempty"` // 订单关闭完成时间
	Result          string `json:"result,omitempty"`          // 订单状态
	RespCode        string `json:"respCode"`                  // 业务返回码
	RespMsg         string `json:"respMsg"`                   // 业务返回说明
}

func (x *Yst2Ka) Tx2295(ctx context.Context, dto *Tx2295Dto) (_ *Tx2295Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2295`, data); err != nil {
		return
	}

	var result Tx2295Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
