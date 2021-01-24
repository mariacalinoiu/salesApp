package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"salesApp/src/datasources"
)

func HandleSucursale(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
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
		response, status, err = getSucursale(db, logger)
	//case http.MethodPost, http.MethodPut:
	//	response, status, err = insertSucursale(r, db, logger, r.Method == http.MethodPut)
	//case http.MethodDelete:
	//	status, err = deleteOrder(r, db, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /sucursale route")
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

func getSucursale(db datasources.DBClient, logger *log.Logger) ([]byte, int, error) {
	sucursale, err := db.GetSucursale()
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get sucursale")
	}

	response, err := json.Marshal(sucursale)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal sucursale response json")
	}

	return response, http.StatusOK, nil
}
