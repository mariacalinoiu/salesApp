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

func HandleParteneri(w http.ResponseWriter, r *http.Request, db datasources.DBClient, logger *log.Logger) {
	var response []byte
	var status int
	var err error

	switch r.Method {
	case http.MethodGet:
		response, status, err = getParteneri(db, logger)
	//case http.MethodPost, http.MethodPut:
	//	response, status, err = insertParteneri(r, db, logger, r.Method == http.MethodPut)
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

//func insertParteneri(r *http.Request, db datasources.DBClient, logger *log.Logger, update bool) ([]byte, int, error) {
//	order, err := extractPartenerParams(r)
//	if err != nil {
//		return nil, http.StatusBadRequest, errors.New("order information sent on request body does not match required format")
//	}
//
//	//if update {
//	//	err = db.EditOrder(order)
//	//} else {
//	orderID, err = db.InsertOrder(order)
//	//}
//	if err != nil {
//		logger.Printf("Internal error: %s", err.Error())
//		return nil, http.StatusInternalServerError, errors.New("could not save Order")
//	}
//
//	response, err := json.Marshal(orderID)
//	if err != nil {
//		return nil, http.StatusInternalServerError, errors.New("could not marshal orderID response json")
//	}
//
//	return response, http.StatusOK, nil
//}

//func deleteOrder(r *http.Request, db datasources.DBClient, logger *log.Logger) (int, error) {
//	params, ok := r.URL.Query()["orderID"]
//
//	if !ok || len(params[0]) < 1 {
//		return http.StatusBadRequest, errors.New("mandatory parameter 'orderID' not found")
//	}
//
//	orderID, err := strconv.Atoi(params[0])
//	if err != nil {
//		return http.StatusBadRequest, errors.New("could not convert parameter 'orderID' to integer")
//	}
//	err = db.DeleteOrder(orderID)
//	if err != nil {
//		logger.Printf("Internal error: %s", err.Error())
//		return http.StatusInternalServerError, errors.New("could not delete Order")
//	}
//
//	return http.StatusOK, nil
//}
