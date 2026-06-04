package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx2091(t *testing.T) {
	ctx := context.TODO()
	batchNo := Num(`B`, cfg.PersonCode, `0`)
	reqTraceNum := Num(`X`, cfg.PersonCode, `0`)

	orgRespTraceNum := ``
	receiverSignNum := ``

	dto := yst2ka.NewTx2091Dto(batchNo,
		[]yst2ka.Tx2091Apply{*yst2ka.NewTx2091Apply(
			reqTraceNum,
			[]yst2ka.Tx2091ApplyInfo{*yst2ka.NewTx2091ApplyInfo(100).SetOrgRespTraceNum(orgRespTraceNum)},
			receiverSignNum,
			100,
		)},
	).SetRespUrl(v.Notify(`/tx2091/callback`))

	r, err := client.Tx2091(ctx, dto)
	assert.NoError(t, err)

	t.Log(`batchNo:`, r.BatchNo)
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
}
