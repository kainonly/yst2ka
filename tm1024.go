package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1024Dto struct {
	ReqTraceNum     string `json:"reqTraceNum"`               // 请求流水号
	SignNum         string `json:"signNum"`                   // 商户会员编号
	OpType          string `json:"opType"`                    // 操作类型，set-绑定收银宝商户 query-查询
	MemberRole      string `json:"memberRole,omitempty"`      // 会员角色，未注册或未实名场景按文档上送
	SybMerchantCode string `json:"sybMerchantCode,omitempty"` // 收银宝商户号，opType=set 时必填
}

func NewTm1024Dto(reqTraceNum string, signNum string, opType string) *Tm1024Dto {
	return &Tm1024Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		OpType:      opType,
	}
}

func (x *Tm1024Dto) SetMemberRole(v string) *Tm1024Dto {
	x.MemberRole = v
	return x
}

func (x *Tm1024Dto) SetSybMerchantCode(v string) *Tm1024Dto {
	x.SybMerchantCode = v
	return x
}

type Tm1024Result struct {
	RespTraceNum         string              `json:"respTraceNum,omitempty"`         // 响应流水号
	SignNum              string              `json:"signNum,omitempty"`              // 商户会员编号
	RespCode             string              `json:"respCode"`                       // 业务返回码
	RespMsg              string              `json:"respMsg"`                        // 失败原因
	SybMerchantCodeArray []Tm1024SybMerchant `json:"sybMerchantCodeArray,omitempty"` // 已绑定收银宝商户列表
}

type Tm1024SybMerchant struct {
	SybMerchantCode string `json:"sybMerchantCode"` // 收银宝商户号
	BindTime        string `json:"bindTime"`        // 绑定时间，格式 yyyy-MM-dd HH:mm:ss
}

func (x *Yst2Ka) Tm1024(ctx context.Context, dto *Tm1024Dto) (_ *Tm1024Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/tm/handle`, `1024`, data); err != nil {
		return
	}

	var result Tm1024Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
