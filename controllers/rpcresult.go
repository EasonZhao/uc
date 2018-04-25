package controllers

const (
	STATUS_OK  = "ok"
	STATUS_ERR = "error"
)

type RPCResult struct {
	Status string
	Data   map[string]string
}

func NewRPCResult(st string) *RPCResult {
	result := RPCResult{}
	result.Status = st
	result.Data = make(map[string]string)
	return &result
}
