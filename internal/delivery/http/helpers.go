package http

import (
	"net/http"
	"strconv"
)

func getPostFormInt(w http.ResponseWriter, r *http.Request, name string) (int, error) {
	result, err := strconv.Atoi(r.PostFormValue(name))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return 0, err
	}
	return result, nil
}
