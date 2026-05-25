package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3004Dto struct {
	SignNum      string `json:"signNum"`      // 商户会员编号
	AcctType     string `json:"acctType"`     // 账户类型
	QryTransCode string `json:"qryTransCode"` // 订单类型
	RespTraceNum string `json:"respTraceNum"` // 通联订单号
	BeginTime    string `json:"beginTime"`    // 开始日期，格式 yyyy-MM-dd HH:mm:ss
	EndTime      string `json:"endTime"`      // 结束日期，格式 yyyy-MM-dd HH:mm:ss
	QryStart     string `json:"qryStart"`     // 起始位置（从 1 开始）
	QryCount     string `json:"qryCount"`     // 查询条数（仅支持 100 以内）
}

func NewTq3004Dto(signNum string, acctType string) *Tq3004Dto {
	return &Tq3004Dto{
		SignNum:  signNum,
		AcctType: acctType,
	}
}

func (x *Tq3004Dto) SetQryTransCode(qryTransCode string) *Tq3004Dto {
	x.QryTransCode = qryTransCode
	return x
}

func (x *Tq3004Dto) SetRespTraceNum(respTraceNum string) *Tq3004Dto {
	x.RespTraceNum = respTraceNum
	return x
}

func (x *Tq3004Dto) SetBeginTime(beginTime string) *Tq3004Dto {
	x.BeginTime = beginTime
	return x
}

func (x *Tq3004Dto) SetEndTime(endTime string) *Tq3004Dto {
	x.EndTime = endTime
	return x
}

func (x *Tq3004Dto) SetQryStart(qryStart string) *Tq3004Dto {
	x.QryStart = qryStart
	return x
}

func (x *Tq3004Dto) SetQryCount(qryCount string) *Tq3004Dto {
	x.QryCount = qryCount
	return x
}

type Tq3004Result struct {
	SignNum     string             `json:"signNum"`     // 商户会员编号
	AcctDetails []Tq3004AcctDetail `json:"acctDetails"` // 账户明细记录
	RespCode    string             `json:"respCode"`    // 业务返回码
	RespMsg     string             `json:"respMsg"`     // 业务返回说明
	TotalCount  int                `json:"totalCount"`  // 总笔数
}

type Tq3004AcctDetail struct {
	RespTraceNum string `json:"respTraceNum"` // 通联订单号
	FinishTime   string `json:"finishTime"`   // 交易完成时间，格式 yyyy-MM-dd HH:mm:ss
	TransCode    string `json:"transCode"`    // 订单类型
	AccountType  string `json:"accountType"`  // 账户交易类型
	TxAviAmt     string `json:"txAviAmt"`     // 可用变更金额（正值增加，负值减少）
	TxFrzAmt     string `json:"txFrzAmt"`     // 在途变更金额（正值增加，负值减少）
	PreAviAmt    string `json:"preAviAmt"`    // 动账前可用金额
	PreFrzAmt    string `json:"preFrzAmt"`    // 动账前在途金额
	AfterAviAmt  string `json:"afterAviAmt"`  // 动账后可用金额
	AfterFrzAmt  string `json:"afterFrzAmt"`  // 动账后在途余额
}

func (x *Yst2Ka) Tq3004(ctx context.Context, dto *Tq3004Dto) (_ *Tq3004Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tq/handle`, `3004`, data); err != nil {
		return
	}

	var result Tq3004Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
