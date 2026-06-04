package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx4016(t *testing.T) {
	ctx := context.TODO()
	serviceID := ``
	creditAgreementID := ``

	dto := yst2ka.NewTx4016Dto(Num(`X`, cfg.EnterpriseCode, `0`), `5`, serviceID, ``)
	err := dto.SetBizParamJSON(map[string]string{
		"creditAgreementId": creditAgreementID,
	})
	assert.NoError(t, err)

	if err != nil {
		return
	}

	r, err := client.Tx4016(ctx, dto)
	assert.NoError(t, err)

	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`bizSubCode:`, r.BizSubCode)
	t.Log(`bizSubMsg:`, r.BizSubMsg)
	t.Log(`vspCusid:`, r.VspCusid)
	t.Log(`bizParam:`, r.BizParam)
}
