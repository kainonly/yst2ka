package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1010Dto struct {
	ReqTraceNum string `json:"reqTraceNum"` // 商户请求流水号
	SignNum     string `json:"signNum"`     // 商户会员编号
	MemberRole  string `json:"memberRole"`  // 会员角色
	Name        string `json:"name"`        // 姓名
	CerType     string `json:"cerType"`     // 证件类型
	CerNum      string `json:"cerNum"`      // 证件号码
	AcctNum     string `json:"acctNum"`     // 银行卡号
	Phone       string `json:"phone"`       // 银行预留手机
	BindType    string `json:"bindType"`    // 绑卡方式
	ValidDate   string `json:"validDate"`   // 有效期，格式为月年
	Cvv2        string `json:"cvv2"`        // CVV2
}

func NewTm1010Dto(reqTraceNum string, signNum string, name string, cerType string, cerNum string) *Tm1010Dto {
	return &Tm1010Dto{
		ReqTraceNum: reqTraceNum,
		SignNum:     signNum,
		Name:        name,
		CerType:     cerType,
		CerNum:      cerNum,
	}
}

func (x *Tm1010Dto) SetMemberRole(v string) *Tm1010Dto {
	x.MemberRole = v
	return x
}

func (x *Tm1010Dto) SetAcctNum(v string) *Tm1010Dto {
	x.AcctNum = v
	return x
}

func (x *Tm1010Dto) SetPhone(v string) *Tm1010Dto {
	x.Phone = v
	return x
}

func (x *Tm1010Dto) SetBindType(v string) *Tm1010Dto {
	x.BindType = v
	return x
}

func (x *Tm1010Dto) SetValidDate(v string) *Tm1010Dto {
	x.ValidDate = v
	return x
}

func (x *Tm1010Dto) SetCvv2(v string) *Tm1010Dto {
	x.Cvv2 = v
	return x
}

type Tm1010Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号
	SignNum      string `json:"signNum"`      // 商户会员编号
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 业务返回说明
}

func (x *Yst2Ka) Tm1010(ctx context.Context, dto *Tm1010Dto) (_ *Tm1010Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1010`, data); err != nil {
		return
	}

	var result Tm1010Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
