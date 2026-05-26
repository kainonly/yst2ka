package yst2ka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bytedance/sonic"
)

type Tx4016Dto struct {
	ReqTraceNum string `json:"reqTraceNum"`        // 请求流水号
	VspCusid    string `json:"vspCusid,omitempty"` // 收银宝商户号
	OpType      string `json:"opType"`             // 服务单操作类型
	ServiceID   string `json:"serviceId"`          // 服务ID
	BizParam    string `json:"bizParam"`           // 业务参数
}

func NewTx4016Dto(reqTraceNum string, opType string, serviceID string, bizParam string) *Tx4016Dto {
	return &Tx4016Dto{
		ReqTraceNum: reqTraceNum,
		OpType:      opType,
		ServiceID:   serviceID,
		BizParam:    bizParam,
	}
}

func (x *Tx4016Dto) SetVspCusid(v string) *Tx4016Dto {
	x.VspCusid = v
	return x
}

func (x *Tx4016Dto) SetBizParamJSON(v any) error {
	data, err := sonic.MarshalString(v)
	if err != nil {
		return err
	}
	x.BizParam = data
	return nil
}

type Tx4016Result struct {
	RespTraceNum string `json:"respTraceNum"`         // 通联订单号
	RespCode     string `json:"respCode"`             // 业务返回码
	RespMsg      string `json:"respMsg"`              // 失败原因
	BizSubCode   string `json:"bizsubcode,omitempty"` // 支付宝业务错误码
	BizSubMsg    string `json:"bizsubmsg,omitempty"`  // 支付宝业务错误原因
	VspCusid     string `json:"vspCusid,omitempty"`   // 收银宝商户号
	BizParam     string `json:"bizParam,omitempty"`   // 业务响应参数
}

func (x *Tx4016Result) DecodeBizParam(v any) error {
	if x.BizParam == `` {
		return nil
	}
	return sonic.UnmarshalString(x.BizParam, v)
}

type Tx4016NotifyResult struct {
	VspCusid  string          `json:"vspCusid"`  // 收银宝子商户号
	EventType string          `json:"eventType"` // 通知类型
	BizParam  json.RawMessage `json:"bizParam"`  // 业务参数
}

func (x *Tx4016NotifyResult) DecodeBizParam(v any) error {
	if len(x.BizParam) == 0 {
		return nil
	}
	return sonic.Unmarshal(x.BizParam, v)
}

func (x *Yst2Ka) Tx4016(ctx context.Context, dto *Tx4016Dto) (_ *Tx4016Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `4016`, data); err != nil {
		return
	}

	var result Tx4016Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
