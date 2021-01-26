package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func getIntParameter(r *http.Request, paramName string) (int, error) {
	stringParam, err := getStringParameter(r, paramName)
	if err != nil {
		return -1, err
	}

	param, err := strconv.Atoi(stringParam)
	if err != nil {
		return -1, fmt.Errorf("could not convert parameter '%s' to integer", paramName)
	}

	return param, nil
}

func getStringParameter(r *http.Request, paramName string) (string, error) {
	params, ok := r.URL.Query()[paramName]

	if !ok || len(params[0]) < 1 {
		return "", fmt.Errorf("mandatory parameter '%s' not found", paramName)
	}

	param := strings.Replace(params[0], `'`, ``, -1)
	param = strings.Replace(param, `"`, ``, -1)
	param = strings.Replace(param, `”`, ``, -1)

	return param, nil
}
