package tracer

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/zzsds/micro-utils/tracer"
)

func TestHTTPWrapper(t *testing.T) {
	trace, ioc, err := tracer.NewJeager("gateway-test", "")
	defer ioc.Close()
	if err != nil {
		log.Fatalln(err)
	}
	opentracing.SetGlobalTracer(trace)

	ts := httptest.NewServer(HTTPWrapper(http.DefaultServeMux))

	defer ts.Close()
	api := ts.URL
	t.Log(api)
	rsp, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	t.Log(rsp.Status, rsp.Request.URL)
}
