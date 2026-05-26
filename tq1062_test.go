package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tq1062(t *testing.T) {
	ctx := context.TODO()
	openBankNo := ``
	dateStart := ``
	dateEnd := ``
	if openBankNo == `` || dateStart == `` || dateEnd == `` {
		t.Skip("请先准备银行管理模式账号和可查询日期后再执行真实请求测试")
	}

	dto := yst2ka.NewTq1062Dto(openBankNo, dateStart, dateEnd, `50`, `1`, Num(`Q`, cfg.PersonCode, `0`))
	r, err := client.Tq1062(ctx, dto)
	assert.NoError(t, err)

	if err == nil {
		t.Log(`totalPage:`, r.TotalPage)
		t.Log(`totalNum:`, r.TotalNum)
		t.Log(`inExpDetail:`, r.InExpDetail)
	}
}
