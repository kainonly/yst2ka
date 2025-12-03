package yst2ka_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/kainonly/go/help"
	"github.com/kainonly/yst2ka"
	"gopkg.in/yaml.v3"
)

var v *Values
var cfg *ConfigMap
var client *yst2ka.Yst2Ka

type Values struct {
	Context string                `yaml:"context"`
	Configs map[string]*ConfigMap `yaml:"configs"`
}

func (x *Values) Config() *ConfigMap {
	return x.Configs[x.Context]
}

func (x *Values) Encrypt(plaintext string) (string, error) {
	return help.SM4Encrypt(x.Config().SecretKey, plaintext)
}

func (x *Values) Notify(path string) string {
	return fmt.Sprintf("%s/%s", x.Config().NotifyUrl, path)
}

type ConfigMap struct {
	BaseUrl              string `yaml:"base_url"`
	NotifyUrl            string `yaml:"notify_url"`
	PrivateKeyStr        string `yaml:"private_key_str"`
	AllinpayPublicKeyStr string `yaml:"allinpay_public_key_str"`
	AppID                string `yaml:"app_id"`
	SecretKey            string `yaml:"secret_key"`
	PersonCode           string `yaml:"person_code"`
	EnterpriseCode       string `yaml:"enterprise_code"`
	Phone                string `yaml:"phone"`
}

func TestMain(m *testing.M) {
	var err error
	var b []byte
	if b, err = os.ReadFile("./config/values.yml"); err != nil {
		return
	}
	if err = yaml.Unmarshal(b, &v); err != nil {
		return
	}
	cfg = v.Config()
	if client, err = yst2ka.NewYst2Ka(yst2ka.Option{
		BaseURL:           cfg.BaseUrl,
		PrivateKey:        cfg.PrivateKeyStr,
		AllinpayPublicKey: cfg.AllinpayPublicKeyStr,
		AppID:             cfg.AppID,
		SecretKey:         cfg.SecretKey,
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
