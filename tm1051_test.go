package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

//func TestYst2Ka_Tm1051Acct(t *testing.T) {
//	ctx := context.TODO()
//	code := `bf10006`
//	num := Num(`X`, code, `0`)
//
//	cerNum, err := v.Encrypt(`110102200305048508`)
//	assert.NoError(t, err)
//
//	authPerAgreeInfo := yst2ka.NewAuthPerAgreeInfo(cfg.Phone, `张三`, cerNum, `1`, `3320240327141772874286441426946`)
//	dto := yst2ka.NewTm1051Dto[yst2ka.AcctAgreementJson](num, code, `李一四`, `1`,
//		*yst2ka.NewAcctAgreementJson().
//			SetPayeeAgreeToken(`3320240327141772874573646393345`).
//			SetWithdrawAgreeToken(`3320240327141772874286441426946`).
//			SetAuthPerAgreeInfo(authPerAgreeInfo),
//		`https://notify.kainonly.com:8443`,
//	)
//	r, err := client.Tm1051(ctx, dto)
//	assert.NoError(t, err)
//
//	t.Log(`code:`, r.RespCode)
//	t.Log(`msg:`, r.RespMsg)
//	t.Log(`signNum:`, r.Phone)
//	t.Log(`respTraceNum:`, r.RespTraceNum)
//}

func TestYst2Ka_Tm1051Pay(t *testing.T) {
	ctx := context.TODO()
	code := `bf10006`
	num := Num(`X`, code, `0`)

	payAgreementJson := yst2ka.NewPayAgreementJson(
		`3320240327141772874573646393345`,
		`3320240327141772874286441426946`,
		`3320240327141772874286441426946`,
	)
	dto := yst2ka.NewTm1051Dto[yst2ka.PayAgreementJson](num, code, `李一四`, `2`,
		*payAgreementJson, `https://notify.kainonly.com:8443`,
	)
	r, err := client.Tm1051(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
