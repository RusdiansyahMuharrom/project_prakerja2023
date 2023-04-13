package configuration

import "net/http"

func AllowOrigins() []string {
	url := []string{"http://127.0.0.1:3001,", "http://projectprakerja2023-production.up.railway.app/"}
	return url
}

func AllowMethods() []string {
	method := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}
	return method
}
