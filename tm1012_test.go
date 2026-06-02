package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1012(t *testing.T) {
	ctx := context.TODO()
	signNum := `T1000`
	dto := yst2ka.NewTm1012Dto(Num(`X`, `H5`, `0`), signNum, `张三`, v.Notify(`/register`)).
		SetMemberRole(`收款方`).
		SetJumpURL(v.Notify(`/register-success`))

	r, err := client.Tm1012(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`regInviteLink:`, r.RegInviteLink)
}
