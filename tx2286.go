package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2286Dto struct {
	ReqTraceNum     string `json:"reqTraceNum"`            // 商户订单号
	OrgRespTraceNum string `json:"orgRespTraceNum"`        // 原垫资发放通联订单号
	OrderAmount     int64  `json:"orderAmount"`            // 还款金额，单位分
	RespURL         string `json:"respUrl,omitempty"`      // 后台通知地址
	Remark          string `json:"remark,omitempty"`       // 备注
	ExtendParams    string `json:"extendParams,omitempty"` // 扩展参数
}

func NewTx2286Dto(reqTraceNum string, orgRespTraceNum string, orderAmount int64) *Tx2286Dto {
	return &Tx2286Dto{
		ReqTraceNum:     reqTraceNum,
		OrgRespTraceNum: orgRespTraceNum,
		OrderAmount:     orderAmount,
	}
}

func (x *Tx2286Dto) SetRespURL(v string) *Tx2286Dto {
	x.RespURL = v
	return x
}

func (x *Tx2286Dto) SetRemark(v string) *Tx2286Dto {
	x.Remark = v
	return x
}

func (x *Tx2286Dto) SetExtendParams(v string) *Tx2286Dto {
	x.ExtendParams = v
	return x
}

type Tx2286Result struct {
	Result       string `json:"result,omitempty"`       // 订单状态
	RespTraceNum string `json:"respTraceNum"`           // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户订单号
	ExtendParams string `json:"extendParams,omitempty"` // 扩展参数
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
}

func (x *Yst2Ka) Tx2286(ctx context.Context, dto *Tx2286Dto) (_ *Tx2286Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2286`, data); err != nil {
		return
	}

	var result Tx2286Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
