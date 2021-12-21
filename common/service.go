package common

import "common/module"

func ServiceJsonData(req *module.JsonDataReq) *module.JsonDataResp {
	resp := &module.JsonDataResp{}
	resp.IntVal = req.IntVal
	return resp
}
