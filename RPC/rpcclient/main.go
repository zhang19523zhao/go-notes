package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpcclient/data"
)

func main() {
	addr := ":9999"
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	request := &data.CalculatorRequest{1, 3}
	respones := &data.CalculatorResponse{}
	defer conn.Close()
	conn.Call("Calculator.ADD", request, respones)
	fmt.Println(respones.Result)
}
