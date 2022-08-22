package api

import (
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	js "github.com/ce-final-project/backend_rest_api/accountService/serializer/json"
	ms "github.com/ce-final-project/backend_rest_api/accountService/serializer/msgpack"
	"log"
	"net/http"
)

type AccountHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Path(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	accountService core.AccountService
}

func NewHandler(accountService core.AccountService) AccountHandler {
	return &handler{accountService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) core.AccountSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Account{}
	}
	return &js.Account{}
}

// TODO Implement all handler
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) Path(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
