package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1051(t *testing.T) {
	ctx := context.TODO()
	num := Num(`X`, cfg.EnterpriseCode, `0`)

	payAgreementJson := yst2ka.NewPayAgreementJson(
		`3320240327141772874573646393345`,
		`3320240327141772874286441426946`,
		`3320240327141772874286441426946`,
	)
	dto := yst2ka.NewTm1051Dto[yst2ka.PayAgreementJson](num, cfg.EnterpriseCode, `竹溪县子怡鞋店`, `1`,
		*payAgreementJson, v.Notify(`/tm1052/callback`),
	)
	r, err := client.Tm1051(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
}
