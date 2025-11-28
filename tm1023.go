package yst2ka

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/kainonly/go/help"
)

type Tm1023Dto struct {
	SignNum  string `json:"signNum"`  // 商户会员编号
	AcctType string `json:"acctType"` // 账户类型
}

func NewTm1023Dto(signNum string) *Tm1023Dto {
	return &Tm1023Dto{
		SignNum: signNum,
	}
}

func (x *Tm1023Dto) SetAcctType(v string) *Tm1023Dto {
	x.AcctType = v
	return x
}

type Tm1023Result[T any] struct {
	RespCode      string `json:"respCode"`      // 业务返回码
	RespMsg       string `json:"respMsg"`       // 业务返回说明
	SignNum       string `json:"signNum"`       // 商户会员编号
	BalanceDetail T      `json:"balanceDetail"` // 此应用下该会员号对应的所有账户余额
}

type BalanceDetail struct {
	AcctNum           string `json:"acctNum"`                     // 账户号
	AcctType          string `json:"acctType"`                    // 账户类型 01-簿记账户 11-支付账户 02-应用营销账户 03-应用担保账户 04-应用预充手续费 09-应用储值卡账户 10-储值卡待结算户
	TotalAmt          int64  `json:"totalAmt"`                    // 总余额 可用+在途
	TransitAmt        int64  `json:"transitAmt"`                  // 在途余额
	AvailableAmt      int64  `json:"availableAmt"`                // 可用余额 日终零点将"在途余额"更新至"可用余额"
	YesAmt            int64  `json:"yesAmt"`                      // 昨日期末余额 可用+在途
	RetentionLimitAmt int64  `json:"retentionLimitAmt,omitempty"` // 账户留存额度 通过1043-账户留存额度管理接口设置后返回
}

func (x *BalanceDetail) SetRetentionLimitAmt(v int64) *BalanceDetail {
	x.RetentionLimitAmt = v
	return x
}

func (x *Yst2Ka) Tm1023(ctx context.Context, dto *Tm1023Dto, i any) (err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `1023`, data); err != nil {
		return
	}

	fmt.Println(bizData)

	if err = sonic.UnmarshalString(bizData, i); err != nil {
		return
	}
	return
}

func (x *Yst2Ka) GetPlatformBalanceDetail(ctx context.Context) (detail BalanceDetail, err error) {
	var r *Tm1023Result[BalanceDetail]
	if err = x.Tm1023(ctx, NewTm1023Dto(`yunBizUserId_B2C`), &r); err != nil {
		return
	}
	detail = r.BalanceDetail
	return
}

func (x *Yst2Ka) GetMemberBalanceDetails(ctx context.Context, signNum string) (details []BalanceDetail, err error) {
	var r *Tm1023Result[string]
	if err = x.Tm1023(ctx, NewTm1023Dto(signNum), &r); err != nil {
		return
	}
	ValidCodes := map[string]bool{
		"00000": true,
		"66666": true,
		"66667": true,
	}
	if !ValidCodes[r.RespCode] {
		err = help.E(0, r.RespMsg)
		return
	}

	details = make([]BalanceDetail, 0)
	if err = sonic.UnmarshalString(r.BalanceDetail, &details); err != nil {
		return
	}
	return
}
