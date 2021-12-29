package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/praiakov/godin/api/handler"
	"github.com/praiakov/godin/api/middleware"
	"github.com/praiakov/godin/config"
	"github.com/praiakov/godin/infrastructure/repository"
	"github.com/praiakov/godin/usecase/user"
)

func main() {
	connection := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", config.DB_USER, config.DB_DATABASE, config.DB_PASSWORD)
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewUserMyPostgres(db)
	userService := user.NewService(userRepo)

	r := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
	)

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	//user
	handler.MakeUserHandlers(r, *n, userService)
	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
