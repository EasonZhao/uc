package controllers

const (
	STATUS_OK  = "ok"
	STATUS_ERR = "error"
)

type RPCResult struct {
	Status string
	Data   map[string]interface{}
}

func NewRPCResult(st string) *RPCResult {
	result := RPCResult{}
	result.Status = st
	result.Data = make(map[string]interface{})
	return &result
}
