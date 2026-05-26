package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tx4023(t *testing.T) {
	ctx := context.TODO()
	orgRespTraceNum := ``
	bizLink := ``
	if orgRespTraceNum == `` || bizLink == `` {
		t.Skip("请先准备有效的原通联订单号和业务跳转链接后再执行真实请求测试")
	}

	dto := yst2ka.NewTx4023Dto(orgRespTraceNum, bizLink, `01`)
	r, err := client.Tx4023(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`respCode:`, r.RespCode)
		t.Log(`respMsg:`, r.RespMsg)
		t.Log(`shareToken:`, r.ShareToken)
		t.Log(`expireDate:`, r.ExpireDate)
	}
}
