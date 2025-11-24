package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1031Dto struct {
	ReqTraceNum string `json:"reqTraceNum"` // 商户请求流水号
	SignNum     string `json:"signNum"`     // 商户会员编号
	OriPhone    string `json:"oriPhone"`    // 原手机号码
}

func NewTm1031Dto(reqTraceNum string, signNum string, oriPhone string) *Tm1031Dto {
	return &Tm1031Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		OriPhone:    oriPhone,
	}
}

type Tm1031Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	SignNum      string `json:"signNum"`      // 商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tm1031(ctx context.Context, dto *Tm1031Dto) (_ *Tm1031Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1031`, data); err != nil {
		return
	}

	var result Tm1031Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
