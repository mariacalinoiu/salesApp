package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"salesApp/src/datasources"
	"salesApp/src/repositories"
)

func HandleLiniiVanzari(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
	var response []byte
	var status int
	var err error

	switch r.Method {
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	case http.MethodGet:
		response, status, err = getLiniiVanzari(r, db, logger)
	//case http.MethodPost, http.MethodPut:
	//	response, status, err = insertVanzari(r, db, logger, r.Method == http.MethodPut)
	//case http.MethodDelete:
	//	status, err = deleteOrder(r, db, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /liniiVanzari route")
	}

	if err != nil {
		logger.Printf("Error: %s; Status: %d %s", err.Error(), status, http.StatusText(status))
		http.Error(w, err.Error(), status)

		return
	}

	if response == nil {
		response, _ = json.Marshal(repositories.WasSuccess{Success: true})
	}

	_, err = w.Write(response)
	if err != nil {
		status = http.StatusInternalServerError
		logger.Printf("Error: %s; Status: %d %s", err.Error(), status, http.StatusText(status))
		http.Error(w, err.Error(), status)

		return
	}

	status = http.StatusOK
	logger.Printf("Status: %d %s", status, http.StatusText(status))
}

func getLiniiVanzari(r *http.Request, db datasources.DBClient, logger *log.Logger) ([]byte, int, error) {
	params, ok := r.URL.Query()["IDIntrare"]

	if !ok || len(params[0]) < 1 {
		return nil, http.StatusBadRequest, errors.New("mandatory parameter 'IDIntrare' not found")
	}

	IDIntrare, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("could not convert parameter 'IDIntrare' to integer")
	}
	vanzari, err := db.GetLiniiVanzare(IDIntrare)
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get linii vanzare")
	}

	response, err := json.Marshal(vanzari)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal linii vanzare response json")
	}

	return response, http.StatusOK, nil
}
