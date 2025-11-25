package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1027_BasicInfo(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTm1027Dto(`bf10006`, "1")

	var r yst2ka.Tm1027Result[yst2ka.PersonInfo]
	err := client.Tm1027(ctx, dto, &r)
	assert.NoError(t, err)

	t.Log(r)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)
}
