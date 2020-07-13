package nacos

import (
	"testing"
)

func TestRead(t *testing.T) {
	nacos := NewSource(
		WithEndpoint("acm.aliyun.com:8080"),
		WithNamespace("a0630038-0d1c-4002-8854-0c08c47fa3e3"),
		WithAccountKey("LTAI4FgL4Ew4kGTSEWQ8gSbo"),
		WitchSecretKey("ZElyfnMQ4E4tE8QKJeXdZmgJ54Mgea"),
		WitchDataIDKey("srv.user"),
		WitchGroupKey("dev"))
	n, err := nacos.Read()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", n)
}
