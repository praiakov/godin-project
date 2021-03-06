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

func listUsers(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errorMessage := "Error reading users"

		data, err := service.ListUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

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

func deleteUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.DeleteUser(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorMessage := "Error update user"

		var input struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			TotalMonth int    `json:"total_month"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		vars := mux.Vars(r)

		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.UpdateUser(input.Name, input.Email, input.TotalMonth, id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/users", n.With(
		negroni.Wrap(listUsers(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(getUser(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(deleteUser(service)),
	)).Methods("DELETE", "OPTIONS")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(updateUser(service)),
	)).Methods("PUT", "OPTIONS")
}
