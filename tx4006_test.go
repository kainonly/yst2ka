package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx4006(t *testing.T) {
	ctx := context.TODO()
	serviceID := ``
	authorizationCode := ``
	if serviceID == `` || authorizationCode == `` {
		t.Skip("请先准备有效的微信支付分服务ID和授权协议号后再执行真实请求测试")
	}

	dto := yst2ka.NewTx4006Dto(Num(`X`, cfg.EnterpriseCode, `0`), `2`, serviceID, ``)
	err := dto.SetBizParamJSON(map[string]string{
		"authorizationCode": authorizationCode,
	})
	assert.NoError(t, err)

	if err != nil {
		return
	}

	r, err := client.Tx4006(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respTraceNum:`, r.RespTraceNum)
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`vspCusid:`, r.VspCusid)
		t.Log(`bizParam:`, r.BizParam)
	}
}
