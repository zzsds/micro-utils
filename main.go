package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/zzsds/micro-utils/config/nacos"
)

func main() {
	fmt.Println(123)
	// cs, _ := nacos.NewAutoSource(nacos.WithGroupKey("dev")).Read()
	server := micro.NewService(
		micro.Name("test"),
		micro.Config(config.DefaultConfig),
	)
	server.Options().Config.Load(nacos.NewAutoSource(nacos.WithGroupKey("dev")))
	fmt.Println(string(server.Options().Config.Bytes()))

	// fmt.Println(cs.Format, string(cs.Data))
}
