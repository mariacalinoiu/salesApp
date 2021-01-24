package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mariacalinoiu/salesApp/src/datasources"
	"github.com/mariacalinoiu/salesApp/src/repositories"
)

func HandleArticole(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
	var response []byte
	var status int
	var err error

	switch r.Method {
	case http.MethodGet:
		response, status, err = getArticole(db, logger)
	case http.MethodPost, http.MethodPut:
		response, status, err = insertArticol(r, db, logger, r.Method == http.MethodPut)
	//case http.MethodDelete:
	//	status, err = deleteOrder(r, db, logger)
	default:
		status = http.StatusBadRequest
		err = errors.New("wrong method type for /articole route")
	}

	if err != nil {
		logger.Printf("Error: %s; Status: %d %s", err.Error(), status, http.StatusText(status))
		http.Error(w, err.Error(), status)

		return
	}

	if response != nil {
		_, err = w.Write(response)
		if err != nil {
			status = http.StatusInternalServerError
			logger.Printf("Error: %s; Status: %d %s", err.Error(), status, http.StatusText(status))
			http.Error(w, err.Error(), status)

			return
		}
	}

	status = http.StatusOK
	logger.Printf("Status: %d %s", status, http.StatusText(status))
}

func getArticole(db datasources.DBClient, logger *log.Logger) ([]byte, int, error) {
	articole, err := db.GetArticole()
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not get articole")
	}

	response, err := json.Marshal(articole)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not marshal articole response json")
	}

	return response, http.StatusOK, nil
}

func extractArticolParams(r *http.Request) (repositories.Articol, error) {
	var unmarshalledArticol repositories.Articol

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return repositories.Articol{}, err
	}

	err = json.Unmarshal(body, &unmarshalledArticol)
	if err != nil {
		return repositories.Articol{}, err
	}

	return unmarshalledArticol, nil
}

func insertArticol(r *http.Request, db datasources.DBClient, logger *log.Logger, update bool) ([]byte, int, error) {
	articol, err := extractArticolParams(r)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("articol information sent on request body does not match required format")
	}

	//if update {
	//	err = db.EditOrder(articol)
	//} else {
	err = db.InsertArticol(articol)
	//}
	if err != nil {
		logger.Printf("Internal error: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not save articol")
	}

	return nil, http.StatusOK, nil
}
