package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2086Dto struct {
	ReceiverSignNum string `json:"receiverSignNum"`        // 商户会员编号-垫资收款方
	ReqTraceNum     string `json:"reqTraceNum"`            // 商户订单号
	OrderAmount     int64  `json:"orderAmount"`            // 垫资发放金额，单位分
	RespURL         string `json:"respUrl,omitempty"`      // 后台通知地址
	Remark          string `json:"remark,omitempty"`       // 备注
	ExtendParams    string `json:"extendParams,omitempty"` // 扩展参数
}

func NewTx2086Dto(receiverSignNum string, reqTraceNum string, orderAmount int64) *Tx2086Dto {
	return &Tx2086Dto{
		ReceiverSignNum: receiverSignNum,
		ReqTraceNum:     reqTraceNum,
		OrderAmount:     orderAmount,
	}
}

func (x *Tx2086Dto) SetRespURL(v string) *Tx2086Dto {
	x.RespURL = v
	return x
}

func (x *Tx2086Dto) SetRemark(v string) *Tx2086Dto {
	x.Remark = v
	return x
}

func (x *Tx2086Dto) SetExtendParams(v string) *Tx2086Dto {
	x.ExtendParams = v
	return x
}

type Tx2086Result struct {
	Result       string `json:"result,omitempty"` // 订单状态
	RespTraceNum string `json:"respTraceNum"`     // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`      // 商户订单号
	OrderAmount  int64  `json:"orderAmount"`      // 垫资发放金额
	RespCode     string `json:"respCode"`         // 业务返回码
	RespMsg      string `json:"respMsg"`          // 业务返回说明
}

func (x *Yst2Ka) Tx2086(ctx context.Context, dto *Tx2086Dto) (_ *Tx2086Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2086`, data); err != nil {
		return
	}

	var result Tx2086Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
