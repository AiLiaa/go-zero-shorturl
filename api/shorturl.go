package main

import (
	"flag"
	"fmt"

	"github.com/AiLiaa/go-zero-shorturl/api/internal/config"
	"github.com/AiLiaa/go-zero-shorturl/api/internal/handler"
	"github.com/AiLiaa/go-zero-shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	//根据配置文件加载配置
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//创建HTTP服务器实例
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//将HTTP请求与相应的处理逻辑关联起来
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//启动服务器来提供相应的网络服务
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
