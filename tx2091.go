package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2091Dto struct {
	BatchNo   string            `json:"batchNo"`           // 商户批次号
	ApplyList []Tx2091ApplyList `json:"applyList"`         // 批量担保消费申请订单列表
	RespUrl   string            `json:"respUrl,omitempty"` // 后台通知地址
}

func NewTx2091Dto(batchNo string, applyList []Tx2091ApplyList) *Tx2091Dto {
	return &Tx2091Dto{
		BatchNo:   batchNo,
		ApplyList: applyList,
	}
}

func (x *Tx2091Dto) SetRespUrl(v string) *Tx2091Dto {
	x.RespUrl = v
	return x
}

type Tx2091ApplyList struct {
	ReqTraceNum  string            `json:"reqTraceNum"`            // 商户订单号
	ApplyInfo    []Tx2091ApplyInfo `json:"applyInfo"`              // 源担保消费申请订单付款信息
	SignNum      string            `json:"signNum"`                // 商户会员编号-收款人
	Amount       int64             `json:"amount"`                 // 确认金额
	Summary      string            `json:"summary,omitempty"`      // 摘要
	ExtendParams string            `json:"extendParams,omitempty"` // 商户扩展参数
}

func NewTx2091ApplyList(reqTraceNum string, applyInfo []Tx2091ApplyInfo, signNum string, amount int64) *Tx2091ApplyList {
	return &Tx2091ApplyList{
		ReqTraceNum: reqTraceNum,
		ApplyInfo:   applyInfo,
		SignNum:     signNum,
		Amount:      amount,
	}
}

func (x *Tx2091ApplyList) SetSummary(v string) *Tx2091ApplyList {
	x.Summary = v
	return x
}

func (x *Tx2091ApplyList) SetExtendParams(v string) *Tx2091ApplyList {
	x.ExtendParams = v
	return x
}

type Tx2091ApplyInfo struct {
	OrgReqTraceNum  string            `json:"orgReqTraceNum,omitempty"`  // 担保消费申请商户订单号
	OrgTransDate    string            `json:"orgTransDate,omitempty"`    // 担保消费申请订单创建日期
	OrgRespTraceNum string            `json:"orgRespTraceNum,omitempty"` // 担保消费申请通联订单号
	OrderAmount     int64             `json:"orderAmount"`               // 金额
	CouponAmount    int64             `json:"couponAmount,omitempty"`    // 平台抽佣金额
	SepDetail       []Tx2091SepDetail `json:"sepDetail,omitempty"`       // 分账列表
}

func NewTx2091ApplyInfo(orderAmount int64) *Tx2091ApplyInfo {
	return &Tx2091ApplyInfo{
		OrderAmount: orderAmount,
	}
}

func (x *Tx2091ApplyInfo) SetOrgReqTraceNum(v string) *Tx2091ApplyInfo {
	x.OrgReqTraceNum = v
	return x
}

func (x *Tx2091ApplyInfo) SetOrgTransDate(v string) *Tx2091ApplyInfo {
	x.OrgTransDate = v
	return x
}

func (x *Tx2091ApplyInfo) SetOrgRespTraceNum(v string) *Tx2091ApplyInfo {
	x.OrgRespTraceNum = v
	return x
}

func (x *Tx2091ApplyInfo) SetCouponAmount(v int64) *Tx2091ApplyInfo {
	x.CouponAmount = v
	return x
}

func (x *Tx2091ApplyInfo) SetSepDetail(v []Tx2091SepDetail) *Tx2091ApplyInfo {
	x.SepDetail = v
	return x
}

type Tx2091SepDetail struct {
	SignNum string `json:"signNum"`          // 商户会员编号
	Amount  int64  `json:"amount"`           // 分账金额
	Remark  string `json:"remark,omitempty"` // 备注
}

func NewTx2091SepDetail(signNum string, amount int64) *Tx2091SepDetail {
	return &Tx2091SepDetail{
		SignNum: signNum,
		Amount:  amount,
	}
}

func (x *Tx2091SepDetail) SetRemark(v string) *Tx2091SepDetail {
	x.Remark = v
	return x
}

type Tx2091Result struct {
	BatchNo  string `json:"batchNo"`  // 商户批次号
	RespCode string `json:"respCode"` // 业务返回码
	RespMsg  string `json:"respMsg"`  // 业务返回说明
}

func (x *Yst2Ka) Tx2091(ctx context.Context, dto *Tx2091Dto) (_ *Tx2091Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tx/handle`, `2091`, data); err != nil {
		return
	}

	var result Tx2091Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
