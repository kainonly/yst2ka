package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tq3002(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTq3002Dto("20260525170000300200000001")

	r, err := client.Tq3002(ctx, dto)
	assert.NoError(t, err)
	t.Log(`result:`, r)
}
