package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"salesApp/src/datasources"
	"salesApp/src/repositories"
)

func HandleParteneri(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
	var response []byte
	var status int
	var err error

	switch r.Method {
	case http.MethodGet:
		response, status, err = getParteneri(db, logger)
	case http.MethodPost, http.MethodPut:
		status, err = insertPartener(r, db, logger, r.Method == http.MethodPut)
	//case http.MethodDelete:
	//	status, err = deleteOrder(r, db, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /parteneri route")
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

func getParteneri(db datasources.DBClient, logger *log.Logger) ([]byte, int, error) {
	parteneri, err := db.GetParteneri()
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get parteneri")
	}

	response, err := json.Marshal(parteneri)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal parteneri response json")
	}

	return response, http.StatusOK, nil
}

func extractPartenerParams(r *http.Request) (repositories.Partener, error) {
	var unmarshalledPartener repositories.Partener

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return repositories.Partener{}, err
	}

	err = json.Unmarshal(body, &unmarshalledPartener)
	if err != nil {
		return repositories.Partener{}, err
	}

	return unmarshalledPartener, nil
}

func insertPartener(r *http.Request, db datasources.DBClient, logger *log.Logger, update bool) (int, error) {
	partener, err := extractPartenerParams(r)
	if err != nil {
		return http.StatusBadRequest, errors.New("order information sent on request body does not match required format")
	}

	//if update {
	//	err = db.EditPartener(proiect)
	//} else {
	err = db.InsertPartener(partener)
	//}
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return http.StatusInternalServerError, errors.New("could not save articol")
	}

	return http.StatusOK, nil
}
