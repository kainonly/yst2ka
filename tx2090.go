package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2090Dto struct {
	ReqTraceNum     string           `json:"reqTraceNum"`               // 商户订单号
	OrgReqTraceNum  string           `json:"orgReqTraceNum,omitempty"`  // 担保消费申请商户订单号
	OrgTransDate    string           `json:"orgTransDate,omitempty"`    // 担保消费申请订单创建日期
	OrgRespTraceNum string           `json:"orgRespTraceNum,omitempty"` // 担保消费申请通联订单号
	ReceiverList    []Tx2090Receiver `json:"receiverList"`              // 收款人列表
	RespUrl         string           `json:"respUrl,omitempty"`         // 后台通知地址
	Summary         string           `json:"summary,omitempty"`         // 摘要
	ExtendParams    string           `json:"extendParams,omitempty"`    // 扩展参数
}

func NewTx2090Dto(reqTraceNum string, receiverList []Tx2090Receiver) *Tx2090Dto {
	return &Tx2090Dto{
		ReqTraceNum:  reqTraceNum,
		ReceiverList: receiverList,
	}
}

func (x *Tx2090Dto) SetOrgReqTraceNum(v string) *Tx2090Dto {
	x.OrgReqTraceNum = v
	return x
}

func (x *Tx2090Dto) SetOrgTransDate(v string) *Tx2090Dto {
	x.OrgTransDate = v
	return x
}

func (x *Tx2090Dto) SetOrgRespTraceNum(v string) *Tx2090Dto {
	x.OrgRespTraceNum = v
	return x
}

func (x *Tx2090Dto) SetRespUrl(v string) *Tx2090Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2090Dto) SetSummary(v string) *Tx2090Dto {
	x.Summary = v
	return x
}

func (x *Tx2090Dto) SetExtendParams(v string) *Tx2090Dto {
	x.ExtendParams = v
	return x
}

type Tx2090Receiver struct {
	SignNum      string            `json:"signNum"`                // 商户会员编号
	Amount       int64             `json:"amount"`                 // 金额
	CouponAmount int64             `json:"couponAmount,omitempty"` // 平台抽佣金额
	SepDetail    []Tx2090SepDetail `json:"sepDetail,omitempty"`    // 分账列表
}

func NewTx2090Receiver(signNum string, amount int64) *Tx2090Receiver {
	return &Tx2090Receiver{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2090Receiver) SetCouponAmount(v int64) *Tx2090Receiver {
	x.CouponAmount = v
	return x
}

func (x *Tx2090Receiver) SetSepDetail(v []Tx2090SepDetail) *Tx2090Receiver {
	x.SepDetail = v
	return x
}

type Tx2090SepDetail struct {
	SignNum string `json:"signNum"`          // 商户会员编号
	Amount  int64  `json:"amount"`           // 分账金额
	Remark  string `json:"remark,omitempty"` // 备注
}

func NewTx2090SepDetail(signNum string, amount int64) *Tx2090SepDetail {
	return &Tx2090SepDetail{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2090SepDetail) SetRemark(v string) *Tx2090SepDetail {
	x.Remark = v
	return x
}

type Tx2090Result struct {
	Result       string `json:"result,omitempty"` // 订单状态
	ReqTraceNum  string `json:"reqTraceNum"`      // 商户订单号
	RespTraceNum string `json:"respTraceNum"`     // 通联订单号
	RespCode     string `json:"respCode"`         // 业务返回码
	RespMsg      string `json:"respMsg"`          // 业务返回说明
}

func (x *Yst2Ka) Tx2090(ctx context.Context, dto *Tx2090Dto) (_ *Tx2090Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2090`, data); err != nil {
		return
	}

	var result Tx2090Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
