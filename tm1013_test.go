package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1013(t *testing.T) {
	ctx := context.TODO()
	signNum := Num(`E`, `H5`, `0`)
	dto := yst2ka.NewTm1013Dto(
		Num(`X`, `H5`, `1`),
		signNum,
		`上海测试企业有限公司`,
		v.Notify(`/tm1013/callback`),
	).SetJumpURL(v.Notify(`/tm1013/return`))

	r, err := client.Tm1013(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`signNum:`, r.SignNum)
		t.Log(`regInviteLink:`, r.RegInviteLink)
	}
}
