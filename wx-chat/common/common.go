package common

type HttpResponse struct {
	IRet int64  `json:"iRet"`
	SMsg string `json:"sMsg"`
}

func JsonResponse(iRet int64, sMsg string) *HttpResponse {
	return &HttpResponse{
		IRet: iRet,
		SMsg: sMsg,
	}
}
