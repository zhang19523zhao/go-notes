package service

import (
	"fmt"
	"rpcserver/data"
)

type Calculator struct {
}

func (c *Calculator) ADD(request *data.CalculatorRequest, response *data.CalculatorResponse) error {
	response.Result = request.Right + request.Left
	fmt.Printf("in data result: %d ", response.Result)
	return nil
}
