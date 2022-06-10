package controllers

import (
	"encoding/json"
	"net/http"
)

func success(w http.ResponseWriter, body interface{}) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(body)
}

func internalServerError(w http.ResponseWriter, err error) error  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}