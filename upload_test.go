package yst2ka_test

import (
	"context"
	"os"
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

func TestYst2Ka_FileUploadLegpCerFront(t *testing.T) {
	ctx := context.TODO()

	b, err := os.ReadFile(`./ocr/legpCerFront.jpg`)
	assert.NoError(t, err)

	r, err := x.FileUpload(ctx, yst2ka.FileUploadOption{
		Name:  "kain.legpCerFront",
		Type:  "0",
		Bytes: b,
	})
	assert.NoError(t, err)

	t.Log(r)
	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`fileId:`, r.FileId) // 3320251120151991415567696941058
}

func TestYst2Ka_FileUploadLegpCerBack(t *testing.T) {
	ctx := context.TODO()

	b, err := os.ReadFile(`./ocr/legpCerBack.jpg`)
	assert.NoError(t, err)

	r, err := x.FileUpload(ctx, yst2ka.FileUploadOption{
		Name:  "kain.legpCerBack",
		Type:  "0",
		Bytes: b,
	})
	assert.NoError(t, err)

	t.Log(r)
	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`fileId:`, r.FileId) // 3320251120151991413871184859138
}

func TestYst2Ka_FileUploadUnifiedSocialCredit(t *testing.T) {
	ctx := context.TODO()

	b, err := os.ReadFile(`./ocr/unifiedSocialCredit.jpg`)
	assert.NoError(t, err)

	r, err := x.FileUpload(ctx, yst2ka.FileUploadOption{
		Name:  "kain.unifiedSocialCredit",
		Type:  "0",
		Bytes: b,
	})
	assert.NoError(t, err)

	t.Log(r)
	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`fileId:`, r.FileId) // 3320251120151991415604724256769
}
