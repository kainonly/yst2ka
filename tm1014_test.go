package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestTm1014DtoBuilder(t *testing.T) {
	ctx := context.TODO()

	dto := yst2ka.NewTm1014Dto("XES1001-202605251635000014", "ES1001", "6222021234567890")

	r, err := client.Tm1014(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)
}
