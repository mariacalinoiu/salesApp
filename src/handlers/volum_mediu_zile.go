package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"salesApp/src/datasources"
	"salesApp/src/repositories"
)

func HandleVolumMediuZile(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
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
		response, status, err = getVolumMediuZile(db, r, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /volumZile route")
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

func getVolumMediuZile(db datasources.DBClient, r *http.Request, logger *log.Logger) ([]byte, int, error) {
	dataStart, err := getStringParameter(r, "DataStart", false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	dataEnd, err := getStringParameter(r, "DataEnd", false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	articole, err := db.GetVolumLivratZile(dataStart, dataEnd)
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get volumLivratZile")
	}

	response, err := json.Marshal(articole)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal volumLivratZile response json")
	}

	return response, http.StatusOK, nil
}
