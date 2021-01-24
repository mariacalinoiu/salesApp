package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"salesApp/src/datasources"
)

func HandleUnitatiDeMasura(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
	var response []byte
	var status int
	var err error

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.Method {
	case http.MethodGet:
		response, status, err = getUnitatiDeMasura(db, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /um route")
	}

	if err != nil {
		logger.Printf("Error: %s; Status: %d %s", err.Error(), status, http.StatusText(status))
		http.Error(w, err.Error(), status)

		return
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

func getUnitatiDeMasura(db datasources.DBClient, logger *log.Logger) ([]byte, int, error) {
	unitatiDeMasura, err := db.GetUnitatiDeMasura()
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get unitatiDeMasura")
	}

	response, err := json.Marshal(unitatiDeMasura)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal unitatiDeMasura response json")
	}

	return response, http.StatusOK, nil
}
