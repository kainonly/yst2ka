package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3002Dto struct {
	RespTraceNum string `json:"respTraceNum"` // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum"`  // 商户订单号
	OriTransDate string `json:"oriTransDate"` // 订单创建日期，格式 yyyyMMdd
}

func NewTq3002Dto(respTraceNum string) *Tq3002Dto {
	return &Tq3002Dto{
		RespTraceNum: respTraceNum,
	}
}

func (x *Tq3002Dto) SetReqTraceNum(reqTraceNum string) *Tq3002Dto {
	x.ReqTraceNum = reqTraceNum
	return x
}

func (x *Tq3002Dto) SetOriTransDate(oriTransDate string) *Tq3002Dto {
	x.OriTransDate = oriTransDate
	return x
}

type Tq3002Result struct {
	ReqTraceNum      string         `json:"reqTraceNum"`      // 商户订单号
	RespTraceNum     string         `json:"respTraceNum"`     // 通联订单号
	Result           string         `json:"result"`           // 订单状态（0:进行中 1:交易成功 2:交易失败）
	TxDesc           string         `json:"txDesc"`           // 订单状态说明/失败错误信息
	OrgReqTraceNum   string         `json:"orgReqTraceNum"`   // 原商户订单号（退款时返回）
	OrgRespTraceNum  string         `json:"orgRespTraceNum"`  // 云商通原订单号（退款时返回）
	OrderAmount      int64          `json:"orderAmount"`      // 订单金额
	PayAmount        int64          `json:"payAmount"`        // 支付金额
	FinishTime       string         `json:"finishTime"`       // 订单支付完成时间，格式 yyyy-MM-dd HH:mm:ss
	SignNum          string         `json:"signNum"`          // 商户会员编号（付款人）
	ExtendParams     string         `json:"extendParams"`     // 扩展参数
	ChannelParamInfo map[string]any `json:"channelParamInfo"` // 渠道参数信息（支付详情）
	RespCode         string         `json:"respCode"`         // 业务返回码
	RespMsg          string         `json:"respMsg"`          // 业务返回说明
	IsPreConsume     string         `json:"isPreConsume"`     // 是否微信订单预消费（0:否 1:是）
}

func (x *Yst2Ka) Tq3002(ctx context.Context, dto *Tq3002Dto) (_ *Tq3002Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tq/handle`, `3002`, data); err != nil {
		return
	}

	var result Tq3002Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
