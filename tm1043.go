package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1043Dto struct {
	ReqTraceNum string `json:"reqTraceNum"`        // 请求流水号
	OpType      string `json:"opType"`             // 操作类型，query-查询 set-设置
	SignNum     string `json:"signNum"`            // 商户会员编号
	AcctType    string `json:"acctType,omitempty"` // 账户类型，01-簿记账户
	Amount      int64  `json:"amount,omitempty"`   // 账户留存额度，单位分，opType=set 时必填
}

func NewTm1043Dto(reqTraceNum string, opType string, signNum string) *Tm1043Dto {
	return &Tm1043Dto{
		ReqTraceNum: reqTraceNum,
		OpType:      opType,
		SignNum:     signNum,
	}
}

func (x *Tm1043Dto) SetAcctType(v string) *Tm1043Dto {
	x.AcctType = v
	return x
}

func (x *Tm1043Dto) SetAmount(v int64) *Tm1043Dto {
	x.Amount = v
	return x
}

type Tm1043Result struct {
	ReqTraceNum       string `json:"reqTraceNum"`                 // 请求流水号
	RespTraceNum      string `json:"respTraceNum"`                // 通联订单号
	RespCode          string `json:"respCode"`                    // 业务返回码
	RespMsg           string `json:"respMsg"`                     // 失败原因
	RetentionLimitAmt int64  `json:"retentionLimitAmt,omitempty"` // 账户留存额度
	Result            string `json:"result,omitempty"`            // 账户留存额度更新结果
}

func (x *Yst2Ka) Tm1043(ctx context.Context, dto *Tm1043Dto) (_ *Tm1043Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `1043`, data); err != nil {
		return
	}

	var result Tm1043Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
