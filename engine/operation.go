package engine

import (
	"ConfCenter/basic"
	"ConfCenter/config"
	"encoding/json"
	"net/http"
)

/*
	get：查询
	post：插入
	patch：修改
*/

func OperationService(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		err, b := insertService(w, r)
		if err != nil || !b {
			return
		}
		return
	case r.Method == http.MethodGet:
		err := GetService(w, r)
		if err != nil {
			return
		}
		return
	case r.Method == http.MethodPatch:
		err, b := PatchService(w, r)
		if err != nil || !b {
			return
		}
		return
	default:
		errResult.SendErrorResponse(w, config.ErrorMethodFailed)
		return
	}

}

//这里注意，服务名字是唯一的
func insertService(w http.ResponseWriter, r *http.Request) (error, bool) {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), service)
	switch {
	case err != nil:
		config.Log.Info("post json Unmarshal err", err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err, false
		//这里要查询servicename在数据库中有没有
	case !service.GetService():
		config.Log.Info("the service name is  existed", service.ServiceName)
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		return nil, false
	default:
		err := service.InsertService()
		if err != nil {
			config.Log.Error("insert opration err ", err)
			errResult.SendErrorResponse(w, config.DbError)
			return err, false
		}
		result.Response(w)
		return nil, true
	}
}

func GetService(w http.ResponseWriter, r *http.Request) error {
	err, s := service.GetAllService()
	if err != nil {
		config.Log.Error("insert opration err ", err)
		errResult.SendErrorResponse(w, config.DbError)
		return err
	}
	res := make(map[string]interface{}, 1)
	res["result"] = s

	massage, err := json.Marshal(res)
	if err != nil {
		config.Log.Info("get json Unmarshal err", err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	}
	normalResult.Resp = string(massage)
	normalResult.Code = 200
	result.NormalResponse(w, normalResult)
	return nil
}

func PatchService(w http.ResponseWriter, r *http.Request) (error, bool) {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), service)
	switch {
	case err != nil:
		config.Log.Info("post json Unmarshal err", err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err, false
	case !service.GetService():
		config.Log.Info("the service name is  existed", service.ServiceName)
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		return nil, false
	default:
		err := service.UpdateService()
		if err != nil {
			errResult.SendErrorResponse(w, config.OperationDbErr)
			return err, false
		}
		normalResult.Resp = "更新成功"
		normalResult.Code = 200
		result.NormalResponse(w, normalResult)
		return nil, true
	}
}
