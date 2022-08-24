package service

import "rpcserver/data"

type Calculator struct {
}

func (c *Calculator) ADD(request *data.CalculatorRequest, response *data.CalculatorResponse) error {
	response.Result = request.Right + request.Left
	return nil
}
