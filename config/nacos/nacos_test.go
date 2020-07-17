package nacos

import (
	"testing"

	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
)

var client = NewAutoSource(
	// 默认使用文件
	file.WithPath("config.toml"),
	WithEndpoint("acm.aliyun.com:8080"),
	WithNamespace("werewsd-0d1c-4002-8854-0c08c47fa3e3"),
	WithAccountKey("*"),
	WithSecretKey("*"),
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
				"name": "test_user",
				"user": "root",
				"password": "123456",
				"host": "127.0.0.1:3336",
				"debug": true
			},
			"jwt": {
				"publicKey": "123123213"
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
