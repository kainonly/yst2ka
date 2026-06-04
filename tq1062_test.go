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

	dto := yst2ka.NewTq1062Dto(openBankNo, dateStart, dateEnd, `50`, `1`, Num(`Q`, cfg.PersonCode, `0`))
	r, err := client.Tq1062(ctx, dto)
	assert.NoError(t, err)

	t.Log(`totalPage:`, r.TotalPage)
	t.Log(`totalNum:`, r.TotalNum)
	t.Log(`inExpDetail:`, r.InExpDetail)
}
