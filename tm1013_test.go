package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1013(t *testing.T) {
	ctx := context.TODO()
	signNum := `T1006`
	dto := yst2ka.NewTm1013Dto(Num(`X`, `H5`, `1`), signNum, `竹溪县子怡鞋店`, v.Notify(`/register`)).
		SetMemberRole(`收款方`).
		SetEnterpriseNature(`2`).
		SetJumpURL(v.Notify(`/register-success`))

	r, err := client.Tm1013(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`regInviteLink:`, r.RegInviteLink)
}
