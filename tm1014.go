package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1014Dto struct {
	ReqTraceNum string `json:"reqTraceNum"` // 请求流水号
	SignNum     string `json:"signNum"`     // 商户会员编号
	AcctNum     string `json:"acctNum"`     // 银行卡号（SM4 加密）
	AgreementNo string `json:"agreementNo"` // 协议号
}

func NewTm1014Dto(reqTraceNum string, signNum string, acctNum string) *Tm1014Dto {
	return &Tm1014Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		AcctNum:     acctNum,
	}
}

type Tm1014Result struct {
	SignNum  string `json:"signNum"`  // 商户会员编号
	RespCode string `json:"respCode"` // 业务返回码
	RespMsg  string `json:"respMsg"`  // 业务返回说明
}

func (x *Yst2Ka) Tm1014(ctx context.Context, dto *Tm1014Dto) (_ *Tm1014Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1011`, data); err != nil {
		return
	}

	var result Tm1014Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
