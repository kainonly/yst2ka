package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq1062Dto struct {
	OpenBankNo   string `json:"openBankNo"`             // 开户银行编码
	DateStart    string `json:"dateStart"`              // 查询开始日期，格式 yyyyMMdd
	DateEnd      string `json:"dateEnd"`                // 查询结束日期，格式 yyyyMMdd
	QryTradeType string `json:"qryTradeType,omitempty"` // 交易类型
	ReturnRows   string `json:"returnRows"`             // 查询条数
	Page         string `json:"page"`                   // 页数
	ReqTraceNum  string `json:"reqTraceNum"`            // 请求流水号
}

func NewTq1062Dto(openBankNo string, dateStart string, dateEnd string, returnRows string, page string, reqTraceNum string) *Tq1062Dto {
	return &Tq1062Dto{
		OpenBankNo:  openBankNo,
		DateStart:   dateStart,
		DateEnd:     dateEnd,
		ReturnRows:  returnRows,
		Page:        page,
		ReqTraceNum: reqTraceNum,
	}
}

func (x *Tq1062Dto) SetQryTradeType(v string) *Tq1062Dto {
	x.QryTradeType = v
	return x
}

type Tq1062Result struct {
	TotalPage   string              `json:"totalPage"`   // 总页数
	TotalNum    string              `json:"totalNum"`    // 记录总行数
	InExpDetail []Tq1062InExpDetail `json:"inExpDetail"` // 银行账户收支明细
}

type Tq1062InExpDetail struct {
	ChnlTradeCode string `json:"chnlTradeCode"`         // 银行交易流水号
	TransDate     string `json:"transDate"`             // 交易日期，格式 yyyyMMdd
	TradeTime     string `json:"tradeTime"`             // 交易时间，格式 HHMMSS
	FundDirection string `json:"fundDirection"`         // 资金方向，D-账户出金 C-账户入金
	TradeType     string `json:"tradeType,omitempty"`   // 交易类型
	CurAmount     string `json:"curAmount"`             // 当前余额，单位分
	TransAmount   string `json:"transAmount"`           // 交易金额，单位分
	Summary       string `json:"summary,omitempty"`     // 摘要描述
	Remark        string `json:"remark,omitempty"`      // 备注
	OppAcctNo     string `json:"oppAcctNo,omitempty"`   // 对手方账号
	OppAcctName   string `json:"oppAcctName,omitempty"` // 对手方账户名
	SubAcctName   string `json:"subAcctName,omitempty"` // 银行子账户名
	SubAcctNo     string `json:"subAcctNo,omitempty"`   // 银行子账号
}

func (x *Yst2Ka) Tq1062(ctx context.Context, dto *Tq1062Dto) (_ *Tq1062Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tq/handle`, `1062`, data); err != nil {
		return
	}

	var result Tq1062Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
