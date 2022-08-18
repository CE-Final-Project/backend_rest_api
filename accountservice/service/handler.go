package service

import (
	"encoding/json"
	"github.com/ce-final-project/backend_rest_api/accountservice/dbclient"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var DBClient dbclient.IMongo

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountId = mux.Vars(r)["accountId"]

	account, err := DBClient.FindAccount(accountId)

	// if error return 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
