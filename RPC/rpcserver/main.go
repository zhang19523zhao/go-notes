package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcserver/service"
)

func main() {
	//注册服务
	err := rpc.Register(&service.Calculator{})
	if err != nil {
		log.Fatalln(err)
	}

	addr := ":9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen error: %s\n", err)
	}
	log.Printf("listen on: %s", addr)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("conn err %s", err)
			continue
		}
		log.Printf("clinet connection: %s", conn.RemoteAddr())

		//使用例程启动jsonrpc处理客户请求
		go jsonrpc.ServeConn(conn)
	}
}
