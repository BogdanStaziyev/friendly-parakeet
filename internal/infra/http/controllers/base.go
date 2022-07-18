package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"startUp/internal/domain"
)

func ok(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func success(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Print(err)
	}
}

func internalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		log.Print(err)
	}
}

func created(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if body != nil {
		err := json.NewEncoder(w).Encode(map[string]interface{}{"created": body})

		if err != nil {
			log.Print(err)
		}
	}
}

func badRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		log.Print(err)
	}
}

func notFound(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		log.Print(err)
	}
}

func parseUrlQuery(r *http.Request) (*domain.UrlQueryParams, error) {
	params := domain.UrlQueryParams{}
	q := r.URL.Query()
	if q.Has("page") {
		if page, err := strconv.ParseUint(q.Get("page"), 10, 32); err == nil {
			params.Page = uint(page)
		} else {
			return nil, fmt.Errorf("expected 'page' to be an unsigned integer, '%s' was given: %w", q.Get("page"), err)
		}
	}

	if q.Has("pageSize") {
		if size, err := strconv.ParseUint(q.Get("pageSize"), 10, 32); err == nil {
			params.PageSize = uint(size)
		} else {
			return nil, fmt.Errorf("expected 'pageSize' to be an unsigned integer, '%s' was given: %w", q.Get("pageSize"), err)
		}
	}

	if q.Has("showDeleted") {
		if show, err := strconv.ParseUint(q.Get("showDeleted"), 10, 32); err == nil {
			params.ShowDeleted = show > 0
		} else {
			return nil, fmt.Errorf("expected 'showDeleted' to be an unsigned integer, '%s' was given: %w", q.Get("showDeleted"), err)
		}
	}

	return &params, nil
}

//todo uncommenting if need
//func conflict(w http.ResponseWriter, err error) error {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusConflict)
//
//	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
//}

//func validationError(w http.ResponseWriter, err error) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusUnprocessableEntity)
//	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
//	if err != nil {
//		log.Print(err)
//	}
//}

//func genericError(w http.ResponseWriter, err error) error {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusBadRequest)
//
//	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
//}

//func noContent(w http.ResponseWriter) error {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusNoContent)
//
//	return nil
//}
