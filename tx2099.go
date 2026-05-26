package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2099Dto struct {
	BatchNo      string               `json:"batchNo"`            // 批次号
	SignNum      string               `json:"signNum"`            // 商户会员编号-转出方
	AcctType     string               `json:"acctType,omitempty"` // 转出账户类型
	AcctNum      string               `json:"acctNum"`            // 支付账户号-转出方
	TotalCount   string               `json:"totalCount"`         // 转账笔数
	TransferList []Tx2099TransferList `json:"transferList"`       // 转账列表
	RespURL      string               `json:"respUrl,omitempty"`  // 后台通知地址
	Summary      string               `json:"summary,omitempty"`  // 摘要
}

type Tx2099TransferList struct {
	ReqTraceNum string `json:"reqTraceNum"`          // 商户订单号
	InSignNum   string `json:"inSignNum"`            // 商户会员编号-转入方
	InAcctNum   string `json:"inAcctNum"`            // 支付账户号-转入方
	InAcctType  string `json:"inAcctType,omitempty"` // 转入账户类型
	OrderAmount int64  `json:"orderAmount"`          // 转账金额
	Summary     string `json:"summary,omitempty"`    // 摘要
}

func NewTx2099Dto(batchNo string, signNum string, acctNum string, totalCount string, transferList []Tx2099TransferList) *Tx2099Dto {
	return &Tx2099Dto{
		BatchNo:      batchNo,
		SignNum:      signNum,
		AcctNum:      acctNum,
		TotalCount:   totalCount,
		TransferList: transferList,
	}
}

func NewTx2099TransferList(reqTraceNum string, inSignNum string, inAcctNum string, orderAmount int64) *Tx2099TransferList {
	return &Tx2099TransferList{
		ReqTraceNum: reqTraceNum,
		InSignNum:   inSignNum,
		InAcctNum:   inAcctNum,
		OrderAmount: orderAmount,
	}
}

func (x *Tx2099Dto) SetAcctType(v string) *Tx2099Dto {
	x.AcctType = v
	return x
}

func (x *Tx2099Dto) SetRespURL(v string) *Tx2099Dto {
	x.RespURL = v
	return x
}

func (x *Tx2099Dto) SetSummary(v string) *Tx2099Dto {
	x.Summary = v
	return x
}

func (x *Tx2099TransferList) SetInAcctType(v string) *Tx2099TransferList {
	x.InAcctType = v
	return x
}

func (x *Tx2099TransferList) SetSummary(v string) *Tx2099TransferList {
	x.Summary = v
	return x
}

type Tx2099Result struct {
	RespCode string `json:"respCode"`          // 业务返回码
	RespMsg  string `json:"respMsg"`           // 业务返回说明
	AuthWay  string `json:"authWay,omitempty"` // 鉴权方式
	BatchNo  string `json:"batchNo"`           // 批次号
}

func (x *Yst2Ka) Tx2099(ctx context.Context, dto *Tx2099Dto) (_ *Tx2099Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2099`, data); err != nil {
		return
	}

	var result Tx2099Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
