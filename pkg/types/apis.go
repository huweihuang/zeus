package types

type CommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResp struct {
	CommonResp
	Data struct {
		Error string `json:"error"`
	} `json:"data"`
}

type CreateInstanceResp struct {
	CommonResp
	Data struct {
		JobID string `json:"jobId"`
	} `json:"data"`
}

type GetInstanceResp struct {
	CommonResp
	Data *Instance `json:"data"`
}
