package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mariacalinoiu/salesApp/src/datasources"
)

func HandleProiecte(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
	var response []byte
	var status int
	var err error

	switch r.Method {
	case http.MethodGet:
		response, status, err = getProiecte(db, logger)
	//case http.MethodPost, http.MethodPut:
	//	response, status, err = insertProiecte(r, db, logger, r.Method == http.MethodPut)
	//case http.MethodDelete:
	//	status, err = deleteOrder(r, db, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /proiecte route")
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

func getProiecte(db datasources.DBClient, logger *log.Logger) ([]byte, int, error) {
	proiecte, err := db.GetProiecte()
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get proiecte")
	}

	response, err := json.Marshal(proiecte)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal proiecte response json")
	}

	return response, http.StatusOK, nil
}
