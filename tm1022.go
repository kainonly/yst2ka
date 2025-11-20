package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1022Dto struct {
	ReqTraceNum         string `json:"reqTraceNum"`         // 商户请求流水号
	SignNum             string `json:"signNum"`             // 商户会员编号
	NotifyUrl           string `json:"notifyUrl"`           // 企业会员审核结果通知地址
	LegpCerFront        string `json:"legpCerFront"`        // 法人身份证（肖像面）
	LegpCerBack         string `json:"legpCerBack"`         // 法人身份证（国徽面）
	UnifiedSocialCredit string `json:"unifiedSocialCredit"` // 统一信用证
	OtherPhotocopyType  string `json:"otherPhotocopyType"`  // 其他影印件类型
	PhotocopyToken      string `json:"photocopyToken"`      // 影印件图片文件
}

func NewTm1022Dto(reqTraceNum string, signNum string, notifyUrl string) *Tm1022Dto {
	return &Tm1022Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		NotifyUrl:   notifyUrl,
	}
}

func (x *Tm1022Dto) SetLegpCerFront(v string) *Tm1022Dto {
	x.LegpCerFront = v
	return x
}

func (x *Tm1022Dto) SetLegpCerBack(v string) *Tm1022Dto {
	x.LegpCerBack = v
	return x
}

func (x *Tm1022Dto) SetUnifiedSocialCredit(v string) *Tm1022Dto {
	x.UnifiedSocialCredit = v
	return x
}

func (x *Tm1022Dto) SetOtherPhotocopyType(v string) *Tm1022Dto {
	x.OtherPhotocopyType = v
	return x
}

func (x *Tm1022Dto) SetPhotocopyToken(v string) *Tm1022Dto {
	x.PhotocopyToken = v
	return x
}

type Tm1022Result struct {
	ReqTraceNum     string `json:"reqTraceNum"`               // 商户订单号
	RespTraceNum    string `json:"respTraceNum"`              // 通联订单号
	CloseResult     string `json:"closeResult,omitempty"`     // 订单关闭结果
	CloseFinishTime string `json:"closeFinishTime,omitempty"` // 订单关闭完成时间
	Result          string `json:"result,omitempty"`          // 订单状态
	RespCode        string `json:"respCode"`                  // 业务返回码
	RespMsg         string `json:"respMsg"`                   // 业务返回说明
}

func (x *Yst2Ka) Tm1022(ctx context.Context, dto *Tm1022Dto) (_ *Tm1022Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1022`, data); err != nil {
		return
	}

	var result Tm1022Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
