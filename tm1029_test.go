package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1029(t *testing.T) {
	ctx := context.TODO()

	num := Num(`X`, cfg.PersonCode, `0`)
	dto := yst2ka.NewTm1029Dto(num, cfg.PersonCode, v.Notify(`/tm1029/callback`)).
		SetJumpUrl(v.Notify(`/tm1029/return`))

	r, err := client.Tm1029(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`signNum:`, r.SignNum)
		t.Log(`openAcctStatus:`, r.OpenAcctStatus)
		t.Log(`openAcctUrl:`, r.OpenAcctUrl)
	}
}
