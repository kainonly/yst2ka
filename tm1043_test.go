package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1043(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTm1043Dto(Num(`X`, cfg.PersonCode, `0`), `query`, cfg.PersonCode)
	r, err := client.Tm1043(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`reqTraceNum:`, r.ReqTraceNum)
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`retentionLimitAmt:`, r.RetentionLimitAmt)
		t.Log(`result:`, r.Result)
	}
}
