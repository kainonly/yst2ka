package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tq3004(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTq3004Dto("ES1001", "01").
		SetBeginTime(`2025-12-01 00:00:00`).
		SetEndTime(`2025-12-07 23:59:59`).
		SetQryStart(`1`).
		SetQryCount(`10`)

	r, err := client.Tq3004(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)
	t.Log(`totalCount:`, r.TotalCount)
	t.Log(`acctDetails:`, r.AcctDetails)
}
