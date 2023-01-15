package main

import (
	biz "kitex_demo/kitex_gen/biz/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	svr := biz.NewServer(NewUserServiceImpl(), server.WithServiceAddr(addr)) // 指定server的ip:port
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
