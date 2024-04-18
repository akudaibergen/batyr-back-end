package objects

import (
	"app/pkg/logging"
	"app/pkg/response"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	users          = "/api/users"
	allUsers       = "/api/all-users-list"
	userByUid      = "/api/users/:uid"
	confirmPayment = "/api/confirm-payment/:user_uid"
	cancelPayment  = "/api/cancel-payment/:user_uid"
)

type Handler struct {
	Logger  logging.Logger
	Service Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.POST(users, h.CreateNewUser)
	router.GET(users, h.GetUsers)
	router.GET(allUsers, h.GetAllUsers)
	router.GET(userByUid, h.GetUserByUid)
	router.PUT(confirmPayment, h.ConfirmUserPayments)
	router.PUT(cancelPayment, h.CancelUserPayments)
}

func (h *Handler) CreateNewUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	h.Logger.Info("Create User Response")
	writer.Header().Set("Content-Type", "application/json")
	var payload User
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		response.GetResponse(writer, true, false, "invalid data", nil)
		return
	}

	user, err := h.Service.CreateUser(payload)
	if err != nil {
		response.GetResponse(writer, true, true, "Данные не сохранились. Попробуйте позже.", nil)
		return
	}
	response.GetResponse(writer, true, true, "Данные успешно сохранились.", user)
	return
}

func (h *Handler) GetUsers(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	h.Logger.Info("Create User Response")
	writer.Header().Set("Content-Type", "application/json")

	users, err := h.Service.GetUsers()
	if err != nil {
		response.GetResponse(writer, true, true, "Данные не получены. Попробуйте позже.", nil)
		return
	}
	response.GetResponse(writer, true, true, "Данные успешно получены.", users)
	return
}

func (h *Handler) GetAllUsers(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	h.Logger.Info("Get All User Response")
	writer.Header().Set("Content-Type", "application/json")

	users, err := h.Service.GetAllUsers()
	if err != nil {
		response.GetResponse(writer, true, true, "Данные не получены. Попробуйте позже.", nil)
		return
	}
	response.GetResponse(writer, true, true, "Данные успешно получены.", users)
	return
}

func (h *Handler) GetUserByUid(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	h.Logger.Info("Create User Response")
	writer.Header().Set("Content-Type", "application/json")

	uidStr := params.ByName("uid")
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		response.GetResponse(writer, true, true, "Неверный UUID.", nil)
		return
	}

	user, err := h.Service.GetUserByUid(uid)
	if err != nil {
		response.GetResponse(writer, true, true, "Ошибка сервера", nil)
		return
	}
	response.GetResponse(writer, true, true, "Успешно", user)
	return
}

func (h *Handler) ConfirmUserPayments(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	h.Logger.Info("Create User Response")
	writer.Header().Set("Content-Type", "application/json")

	uidStr := params.ByName("user_uid")
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		response.GetResponse(writer, true, true, "Неверный UUID.", nil)
		return
	}

	err = h.Service.SetUserPayment(uid)
	if err != nil {
		response.GetResponse(writer, true, true, "Ошибка сервера", nil)
		return
	}
	response.GetResponse(writer, true, true, "Платеж прошел успешно", nil)
	return
}

func (h *Handler) CancelUserPayments(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	h.Logger.Info("Create User Response")
	writer.Header().Set("Content-Type", "application/json")

	uidStr := params.ByName("user_uid")
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		response.GetResponse(writer, true, true, "Неверный UUID.", nil)
		return
	}

	err = h.Service.CancelPayments(uid)
	if err != nil {
		response.GetResponse(writer, true, true, "Ошибка сервера", nil)
		return
	}
	response.GetResponse(writer, true, true, "Платеж прошел успешно", nil)
	return
}
