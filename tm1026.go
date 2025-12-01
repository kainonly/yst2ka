package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1026Dto struct {
	ReqTraceNum string `json:"reqTraceNum"`       // 商户请求流水号
	QryType     string `json:"qryType,omitempty"` // 查询类型
	CusId       string `json:"cusId,omitempty"`   // 商户号 查询类型为2-商户带结算资金时必填, 用于查询指定收银宝商户的带结算资金
	QryDate     string `json:"qryDate,omitempty"` // 期末余额查询日期 qryType=4时上送 格式为yyyyMMdd 不送默认今天,D日查询D-1日的期末余额
}

func NewTm1026Dto(reqTraceNum string) *Tm1026Dto {
	return &Tm1026Dto{
		ReqTraceNum: reqTraceNum,
	}
}

func (x *Tm1026Dto) SetQryType(v string) *Tm1026Dto {
	x.QryType = v
	return x
}

func (x *Tm1026Dto) SetCusId(v string) *Tm1026Dto {
	x.CusId = v
	return x
}

func (x *Tm1026Dto) SetQryDate(v string) *Tm1026Dto {
	x.QryDate = v
	return x
}

type Tm1026Result struct {
	TotalAmt         string `json:"totalAmt,omitempty"`         // 通联头寸余额,单位:分 查询成功,返回
	CusId            string `json:"cusId,omitempty"`            // 通联头寸账户号 查询类型1:返回收付通商户号 查询类型2:返回请求的商户号
	BankCardNo       string `json:"bankCardNo,omitempty"`       // 银行账号 资金管理模式为银行管理、银行托管、自主管理时返回
	BankTotalAmt     string `json:"bankTotalAmt,omitempty"`     // 银行头寸余额,单位分 qryType=4时返回
	YesterdayBalance string `json:"yesterdayBalance,omitempty"` // 银行期末余额 资金管理模式为银行管理、自主管理、银行托管时返回
	RespCode         string `json:"respCode"`                   // 业务返回码 00000:代表成功
	RespMsg          string `json:"respMsg,omitempty"`          // 业务返回说明
}

func (x *Yst2Ka) Tm1026(ctx context.Context, dto *Tm1026Dto) (_ *Tm1026Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1026`, data); err != nil {
		return
	}

	var result Tm1026Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
