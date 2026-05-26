package yst2ka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bytedance/sonic"
)

type Tx4006Dto struct {
	ReqTraceNum string `json:"reqTraceNum"`        // 请求流水号
	VspCusid    string `json:"vspCusid,omitempty"` // 收银宝商户号
	OpType      string `json:"opType"`             // 服务单操作类型
	ServiceID   string `json:"serviceId"`          // 服务ID
	BizParam    string `json:"bizParam"`           // 业务参数
}

func NewTx4006Dto(reqTraceNum string, opType string, serviceID string, bizParam string) *Tx4006Dto {
	return &Tx4006Dto{
		ReqTraceNum: reqTraceNum,
		OpType:      opType,
		ServiceID:   serviceID,
		BizParam:    bizParam,
	}
}

func (x *Tx4006Dto) SetVspCusid(v string) *Tx4006Dto {
	x.VspCusid = v
	return x
}

func (x *Tx4006Dto) SetBizParamJSON(v any) error {
	data, err := sonic.MarshalString(v)
	if err != nil {
		return err
	}
	x.BizParam = data
	return nil
}

type Tx4006Result struct {
	RespTraceNum string `json:"respTraceNum"`       // 通联订单号
	RespCode     string `json:"respCode"`           // 业务返回码
	RespMsg      string `json:"respMsg"`            // 失败原因
	VspCusid     string `json:"vspCusid,omitempty"` // 收银宝商户号
	BizParam     string `json:"bizParam,omitempty"` // 业务响应参数
}

func (x *Tx4006Result) DecodeBizParam(v any) error {
	if x.BizParam == `` {
		return nil
	}
	return sonic.UnmarshalString(x.BizParam, v)
}

type Tx4006NotifyResult struct {
	VspCusid  string          `json:"vspCusid"`  // 收银宝子商户号
	EventType string          `json:"eventType"` // 通知类型
	Summary   string          `json:"summary"`   // 回调摘要
	BizParam  json.RawMessage `json:"bizParam"`  // 业务参数
}

func (x *Tx4006NotifyResult) DecodeBizParam(v any) error {
	if len(x.BizParam) == 0 {
		return nil
	}
	return sonic.Unmarshal(x.BizParam, v)
}

func (x *Yst2Ka) Tx4006(ctx context.Context, dto *Tx4006Dto) (_ *Tx4006Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `4006`, data); err != nil {
		return
	}

	var result Tx4006Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
