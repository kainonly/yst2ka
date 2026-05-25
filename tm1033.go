package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tm1033Dto struct {
	ReqTraceNum    string               `json:"reqTraceNum"`    // 商户请求流水号
	SignNum        string               `json:"signNum"`        // 商户会员编号
	BankAcctDetail Tm1033BankAcctDetail `json:"bankAcctDetail"` // 银行账户信息
}

type Tm1033BankAcctDetail struct {
	AcctAttr           string `json:"acctAttr"`           // 账户类型（1-对公，不填默认1-对公）
	AcctNum            string `json:"acctNum"`            // 账号（SM4加密）
	OpenBankNo         string `json:"openBankNo"`         // 银行代码
	OpenBankBranchName string `json:"openBankBranchName"` // 开户行支行名称（账户类型=1-对公则必填）
	PayBankNumber      string `json:"payBankNumber"`      // 支付行号，12位数字（账户类型=1-对公则必填）
	OpenBankProvince   string `json:"openBankProvince"`   // 开户行所在省（中文）
	OpenBankCity       string `json:"openBankCity"`       // 开户行所在市（中文）
}

func NewTm1033Dto(reqTraceNum string, signNum string, bankAcctDetail Tm1033BankAcctDetail) *Tm1033Dto {
	return &Tm1033Dto{
		ReqTraceNum:    reqTraceNum,
		SignNum:        signNum,
		BankAcctDetail: bankAcctDetail,
	}
}

type Tm1033Result struct {
	RespTraceNum string `json:"respTraceNum"` // 响应流水号（业务正常处理返回）
	RespCode     string `json:"respCode"`     // 业务返回码
	RespMsg      string `json:"respMsg"`      // 失败原因
}

func (x *Yst2Ka) Tm1033(ctx context.Context, dto *Tm1033Dto) (_ *Tm1033Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1033`, data); err != nil {
		return
	}

	var result Tm1033Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
