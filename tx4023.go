package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx4023Dto struct {
	OrgRespTraceNum string `json:"orgRespTraceNum"` // 原通联订单号
	BizLink         string `json:"bizLink"`         // 跳转业务链接
	OpType          string `json:"opType"`          // 口令类型，01-花呗分期 02-支付宝
}

func NewTx4023Dto(orgRespTraceNum string, bizLink string, opType string) *Tx4023Dto {
	return &Tx4023Dto{
		OrgRespTraceNum: orgRespTraceNum,
		BizLink:         bizLink,
		OpType:          opType,
	}
}

type Tx4023Result struct {
	RespCode   string `json:"respCode"`             // 业务返回码
	RespMsg    string `json:"respMsg"`              // 业务返回说明
	ShareToken string `json:"shareToken,omitempty"` // 吱口令
	ExpireDate string `json:"expireDate,omitempty"` // 有效期
}

func (x *Yst2Ka) Tx4023(ctx context.Context, dto *Tx4023Dto) (_ *Tx4023Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `4023`, data); err != nil {
		return
	}

	var result Tx4023Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
