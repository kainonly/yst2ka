package yst2ka_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_FileUpload(t *testing.T) {
	ctx := context.TODO()

	r, err := x.FileUpload(ctx, yst2ka.FileUploadOption{
		Name:  "hi.text",
		Type:  "0",
		Bytes: []byte(`你好！`),
	})
	assert.NoError(t, err)
	t.Log(r)
}
