package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx4037Dto struct {
	VspCusID string `json:"vspCusid"` // 收银宝商户号
	TrxAmt   int64  `json:"trxamt"`   // 交易金额，不含营销补贴金额
}

func NewTx4037Dto(vspCusID string, trxAmt int64) *Tx4037Dto {
	return &Tx4037Dto{
		VspCusID: vspCusID,
		TrxAmt:   trxAmt,
	}
}

type Tx4037Result struct {
	RespCode string `json:"respCode"`        // 业务返回码
	RespMsg  string `json:"respMsg"`         // 业务返回说明
	Amt3     string `json:"amt3,omitempty"`  // 3 期每期还款
	Fee3     string `json:"fee3,omitempty"`  // 3 期每期利息
	Amt6     string `json:"amt6,omitempty"`  // 6 期每期还款
	Fee6     string `json:"fee6,omitempty"`  // 6 期每期利息
	Amt12    string `json:"amt12,omitempty"` // 12 期每期还款
	Fee12    string `json:"fee12,omitempty"` // 12 期每期利息
}

func (x *Yst2Ka) Tx4037(ctx context.Context, dto *Tx4037Dto) (_ *Tx4037Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `4037`, data); err != nil {
		return
	}

	var result Tx4037Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
