package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2094Dto struct {
	ReqTraceNum  string               `json:"reqTraceNum"`            // 商户订单号
	ReceiverList []Tx2094ReceiverList `json:"receiverList"`           // 收款人列表
	RespUrl      string               `json:"respUrl,omitempty"`      // 后台通知地址
	Remark       string               `json:"remark,omitempty"`       // 备注
	ExtendParams string               `json:"extendParams,omitempty"` // 扩展参数
}

func NewTx2094Dto(reqTraceNum string, receiverList []Tx2094ReceiverList) *Tx2094Dto {
	return &Tx2094Dto{
		ReqTraceNum:  reqTraceNum,
		ReceiverList: receiverList,
	}
}

func (x *Tx2094Dto) SetRespUrl(v string) *Tx2094Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2094Dto) SetRemark(v string) *Tx2094Dto {
	x.Remark = v
	return x
}

func (x *Tx2094Dto) SetExtendParams(v string) *Tx2094Dto {
	x.ExtendParams = v
	return x
}

type Tx2094ReceiverList struct {
	SignNum      string            `json:"signNum"`                // 商户会员编号-收款方
	Amount       int64             `json:"amount"`                 // 核销金额
	CouponAmount int64             `json:"couponAmount,omitempty"` // 平台抽佣金额
	SepDetail    []Tx2094SepDetail `json:"sepDetail,omitempty"`    // 储值卡核销分账人列表
}

func NewTx2094ReceiverList(signNum string, amount int64) *Tx2094ReceiverList {
	return &Tx2094ReceiverList{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2094ReceiverList) SetCouponAmount(v int64) *Tx2094ReceiverList {
	x.CouponAmount = v
	return x
}

func (x *Tx2094ReceiverList) SetSepDetail(v []Tx2094SepDetail) *Tx2094ReceiverList {
	x.SepDetail = v
	return x
}

type Tx2094SepDetail struct {
	SignNum string `json:"signNum"`          // 商户会员编号-分账收款人
	Amount  int64  `json:"amount"`           // 分账金额
	Remark  string `json:"remark,omitempty"` // 备注
}

func NewTx2094SepDetail(signNum string, amount int64) *Tx2094SepDetail {
	return &Tx2094SepDetail{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2094SepDetail) SetRemark(v string) *Tx2094SepDetail {
	x.Remark = v
	return x
}

type Tx2094Result struct {
	Result       string `json:"result,omitempty"`       // 订单状态
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户订单号
	RespTraceNum string `json:"respTraceNum"`           // 通联订单号
	ExtendParams string `json:"extendParams,omitempty"` // 扩展参数
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
}

func (x *Yst2Ka) Tx2094(ctx context.Context, dto *Tx2094Dto) (_ *Tx2094Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2094`, data); err != nil {
		return
	}

	var result Tx2094Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
