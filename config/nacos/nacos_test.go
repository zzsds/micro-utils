package nacos

import (
	"testing"

	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
)

var client = NewAutoSource(
	file.WithPath("config.toml"),
	WithEndpoint("acm.aliyun.com:8080"),
	WithNamespace("a0630038-0d1c-4002-8854-0c08c47fa3e3"),
	WithAccountKey("LTAI4FgL4Ew4kGTSEWQ8gSbo"),
	WithSecretKey("ZElyfnMQ4E4tE8QKJeXdZmgJ54Mgea"),
	WithDataIDKey("srv.user"),
	WithGroupKey("dev"))

func TestRead(t *testing.T) {
	n, err := client.Read()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v, %s", n, string(n.Data))
}

func TestWrite(t *testing.T) {
	err := client.Write(&source.ChangeSet{
		Data: []byte(`{
			"mysql": {
				"name": "welfare_user",
				"user": "root",
				"password": "password",
				"host": "47.114.118.106:13336",
				"debug": true
			},
			"jwt": {
				"publicKey": "3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ=="
			},
			"setting": {
				"distributional": 3,
				"task": {
					"direct": -10,
					"team": -31,
					"permit": -3,
					"audite": -3
				}
			 }
		}`),
	})
	if err != nil {
		t.Log(err)
	}
}
func TestWatch(t *testing.T) {
	w, err := client.Watch()
	if err != nil {
		t.Error(err)
	}
	sc, err := w.Next()
	if err != source.ErrWatcherStopped {
		t.Errorf("expected watcher stopped error, got %v", err)
	}
	t.Logf("%+v %s", sc, []byte(sc.Data))
}
