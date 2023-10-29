package rest

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// GetString gets a request parameter as string
func GetString(r *http.Request, param string) (string, error) {
	p, ok := mux.Vars(r)[param]
	if !ok {
		return "", errors.New("param not found")
	}

	return p, nil
}
