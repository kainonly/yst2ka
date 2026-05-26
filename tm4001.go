package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm4001Dto struct {
	OrgID       string `json:"orgId,omitempty"`       // 收银宝集团商户号
	CusID       string `json:"cusId"`                 // 收银宝商户号
	TermNo      string `json:"termNo"`                // 终端号
	Operation   string `json:"operation"`             // 操作类型，00-新增 01-修改 02-注销 03-查询
	DeviceType  string `json:"deviceType,omitempty"`  // 设备类型
	TermSn      string `json:"termSn,omitempty"`      // 终端序列号
	TermState   string `json:"termState,omitempty"`   // 终端状态
	TermAddress string `json:"termAddress,omitempty"` // 终端地址
	QueryType   string `json:"queryType,omitempty"`   // 查询类型
}

func NewTm4001Dto(cusID string, termNo string, operation string) *Tm4001Dto {
	return &Tm4001Dto{
		CusID:     cusID,
		TermNo:    termNo,
		Operation: operation,
	}
}

func (x *Tm4001Dto) SetOrgID(v string) *Tm4001Dto {
	x.OrgID = v
	return x
}

func (x *Tm4001Dto) SetDeviceType(v string) *Tm4001Dto {
	x.DeviceType = v
	return x
}

func (x *Tm4001Dto) SetTermSn(v string) *Tm4001Dto {
	x.TermSn = v
	return x
}

func (x *Tm4001Dto) SetTermState(v string) *Tm4001Dto {
	x.TermState = v
	return x
}

func (x *Tm4001Dto) SetTermAddress(v string) *Tm4001Dto {
	x.TermAddress = v
	return x
}

func (x *Tm4001Dto) SetQueryType(v string) *Tm4001Dto {
	x.QueryType = v
	return x
}

type Tm4001Result struct {
	RetCode     string `json:"retCode,omitempty"`     // 收银宝终端处理结果
	RetMsg      string `json:"retMsg,omitempty"`      // 终端报备状态
	OrgID       string `json:"orgId,omitempty"`       // 收银宝集团商户号
	CusID       string `json:"cusId,omitempty"`       // 收银宝商户号
	AppID       string `json:"appid,omitempty"`       // 收银宝应用号
	TermNo      string `json:"termNo,omitempty"`      // 终端号
	DeviceType  string `json:"deviceType,omitempty"`  // 设备类型
	TermSn      string `json:"termSn,omitempty"`      // 终端序列号
	TermState   string `json:"termState,omitempty"`   // 终端状态
	TermAddress string `json:"termAddress,omitempty"` // 终端地址
	ErrorCode   string `json:"errorCode,omitempty"`   // 错误代码
	ErrorMsg    string `json:"errorMsg,omitempty"`    // 错误信息
	WxState     string `json:"wxState,omitempty"`     // 微信报备状态
	WxMsg       string `json:"wxMsg,omitempty"`       // 微信报备信息
	AlState     string `json:"alState,omitempty"`     // 支付宝报备状态
	AlMsg       string `json:"alMsg,omitempty"`       // 支付宝报备信息
	UnState     string `json:"unState,omitempty"`     // 银联报备状态
	UnMsg       string `json:"unMsg,omitempty"`       // 银联报备信息
	RespCode    string `json:"respCode"`              // 响应码
	RespMsg     string `json:"respMsg"`               // 响应信息
}

func (x *Yst2Ka) Tm4001(ctx context.Context, dto *Tm4001Dto) (_ *Tm4001Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `4001`, data); err != nil {
		return
	}

	var result Tm4001Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
