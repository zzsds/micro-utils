package tracer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func TestNewJeager(t *testing.T) {
	type args struct {
		name    string
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    opentracing.Tracer
		want1   io.Closer
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				name:    "test",
				address: "127.0.0.1:6831",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := NewJeager(tt.args.name, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJeager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer got1.Close()
			opentracing.SetGlobalTracer(got)
			span := got.StartSpan(fmt.Sprintf("%s trace", os.Args[1]))
			span.SetTag("trace to", os.Args[1])
			defer span.Finish()
			ctx := opentracing.ContextWithSpan(context.Background(), span)
			reqSpan, _ := opentracing.StartSpanFromContext(ctx, "Client_Test1 request")
			defer reqSpan.Finish()

			url := "https://www.baidu.com"
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Errorf("baidu request fail : %v", err)
			}

			ext.SpanKindRPCClient.Set(reqSpan)
			ext.HTTPUrl.Set(reqSpan, url)
			ext.HTTPMethod.Set(reqSpan, "Method")
			ext.HTTPStatusCode.Set(reqSpan, 200)
			ext.DBType.Set(reqSpan, "MySql")
			ext.DBUser.Set(reqSpan, "root")
			reqSpan.Tracer().Inject(
				span.Context(),
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(req.Header),
			)

			reqSpan, ctx = opentracing.StartSpanFromContext(ctx, "Client_Test2 request")
			defer reqSpan.Finish()

			reqSpan, ctx = opentracing.StartSpanFromContext(ctx, "Client_Test3 request")
			defer reqSpan.Finish()

			t.Log("结束")
		})
	}
}
