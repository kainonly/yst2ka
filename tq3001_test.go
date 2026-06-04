package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tq3001(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTq3001Dto("20260525170000300100000001")

	r, err := client.Tq3001(ctx, dto)
	assert.NoError(t, err)

	t.Log(`result:`, r)
}
