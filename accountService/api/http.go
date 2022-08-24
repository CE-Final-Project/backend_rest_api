package api

import (
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	js "github.com/ce-final-project/backend_rest_api/accountService/serializer/json"
	ms "github.com/ce-final-project/backend_rest_api/accountService/serializer/msgpack"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
)

type AccountHandler interface {
	GetAllAccount(w http.ResponseWriter, r *http.Request)
	GetSingleAccount(w http.ResponseWriter, r *http.Request)
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

func (h *handler) GetAllAccount(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	accounts, err := h.accountService.Find("")
	if err != nil {
		if errors.Cause(err) == core.ErrAccountNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err != nil {
		// should never be here but log the error just in case
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetSingleAccount(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	playerId := vars["player_id"]
	account, err := h.accountService.FindOne(playerId)
	if err != nil {
		if errors.Cause(err) == core.ErrAccountNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.serializer(contentType).Encode(&account, w)
	if err != nil {
		// should never be here but log the error just in case
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	account := &core.Account{}
	err = h.serializer(contentType).Decode(requestBody, r)
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
