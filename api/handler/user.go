package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/praiakov/godin/api/presenter"
	"github.com/praiakov/godin/entity"
	"github.com/praiakov/godin/usecase/user"
)

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"

		var input struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			TotalMonth int    `json:"total_month"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateUser(input.Name, input.Email, input.TotalMonth)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.User{
			ID:         id,
			Name:       input.Name,
			Email:      input.Email,
			TotalMonth: input.TotalMonth,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error readind user"
		vars := mux.Vars(r)

		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.UserDetail{
			ID:         data.ID,
			Name:       data.Name,
			Email:      data.Email,
			PaidDate:   data.PaidDate,
			DueDate:    data.DueDate,
			TotalMonth: data.TotalMonth,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(getUser(service)),
	)).Methods("GET", "OPTIONS")
}
