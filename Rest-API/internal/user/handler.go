package user

import (
	"RestAPI/internal/handlers"
	"RestAPI/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// подсказка для себя, т.е. check на соответствие структуры handler интерфейсу Handler
var _ handlers.Handler = &handler{}

// структура хранения логгера и сервиса
type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(http.StatusOK) // 200
	//w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	//w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("this is list of users"))

}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(http.StatusCreated) // 201
	//w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte("this is creating user"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(http.StatusOK) // 200
	//w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte("this is user by uuid"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(http.StatusNoContent) // 204
	//w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte("this is update user"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(http.StatusNoContent) // 204
	//w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte("this is partially update  user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(http.StatusNoContent) // 204
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte("this is delete user"))
}

func (h *handler) Register(router *httprouter.Router) {

	// usersURL
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	// userURL
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}
