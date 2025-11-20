package yst2ka_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/kainonly/yst2ka"
)

var x *yst2ka.Yst2Ka
var privateKeyStr = `MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgiaZmB+feACtziE8SYjVZsaQwLNLRiyO8ebSupeoWIF2gCgYIKoEcz1UBgi2hRANCAATwEo0zq6KaB992PToWeJH52LmfS0sFovnB8/LMaoIAOTlFJtA3YgjWXKlO3KT+GqOCfCC4xE60isCr28tqy7hM`
var allinpayPublicKeyStr = `MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEu9LNkJlyLtjJxtQWIGlcZ/hyHt5eZ7LEH1nfOiK1H9HsE1cMPu5KK5jZVTtAyc7lPMXixUMirf6A3tMbuMbgqg==`
var secretKey = `878427523d3525e070298d44481b8d2e`

func TestMain(m *testing.M) {
	var err error
	if x, err = yst2ka.NewYst2(yst2ka.Option{
		BaseURL:           `https://ibstest.allinpay.com/yst/yst-service-api`,
		PrivateKey:        privateKeyStr,
		AllinpayPublicKey: allinpayPublicKeyStr,
		AppID:             "21762000921804636162",
		SecretKey:         secretKey,
	}); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func Num(prefix string, code string, kind string) string {
	now := time.Now()
	formatter := now.Format("20060102150405")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rand.Intn(999) + 1
	paddedStr := fmt.Sprintf("%03d", num)
	return fmt.Sprintf("%s%s-%s%s%s", prefix, code, formatter, paddedStr, kind)
}
