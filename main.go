package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/debug/trace"
	"github.com/micro/go-micro/v2/debug/trace/memory"
	"github.com/zzsds/micro-utils/config/nacos"
	"google.golang.org/protobuf/proto"
)

func main() {
	cfg, err := config.NewConfig(config.WithSource(nacos.NewAutoSource(file.WithPath("config.toml"), nacos.WithDataIDKey("api.upload"), nacos.WithGroupKey("dev"))))
	if err != nil {
		panic(err)
	}
	server := micro.NewService(
		micro.Name("go.micro.srv.test"),
		micro.Config(cfg),
		micro.Tracer(memory.NewTracer()),
	)
	server.Init(micro.WrapHandler())
	var cf struct {
		Aliyun struct {
			AccessKeyID string `json:"accessKeyId"`
		}
	}

	server.Options().Config.Scan(&cf)
	server.Server().Handle(
		server.Server().NewHandler(
			NewHello(server),
		))

	fmt.Println(string(server.Options().Config.Bytes()), server.Options().Config.Map(), cf)
	server.Run()
	// fmt.Println(cs.Format, string(cs.Data))
}

type Hello struct {
	name  string
	trace trace.Tracer
}

func NewHello(srv micro.Service) *Hello {
	srv.Server().Options().Tracer.Finish(&trace.Span{
		Name: "jayden",
	})
	return &Hello{
		name:  srv.Name(),
		trace: srv.Server().Options().Tracer,
	}
}

func (h *Hello) Hi(ctx context.Context, req *proto.MarshalOptions, rsp *proto.UnmarshalOptions) error {
	span := &trace.Span{}
	_, span = h.trace.Start(ctx, "Hi")
	fmt.Println(ctx, span)
	return nil
}
