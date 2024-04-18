package response

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult(status bool, message string, data interface{}) Result {
	return Result{
		Type:    getStatus(status),
		Message: message,
		Data:    data,
	}
}

func getStatus(status bool) string {
	if status {
		return "success"
	} else {
		return "error"
	}
}

func GetResponse(w http.ResponseWriter, isAuthorized bool, status bool, message string, response interface{}) {
	if isAuthorized == false {
		w.WriteHeader(http.StatusUnauthorized)
	} else if status == true {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	result := NewResult(status, message, response)
	data, _ := json.Marshal(result)
	_, _ = w.Write(data)
	return
}
