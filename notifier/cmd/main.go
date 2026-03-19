package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lastlife77/go-notifier/docs"

	"github.com/gorilla/mux"
	"github.com/lastlife77/go-notifier/internal/broker/rabbit"
	"github.com/lastlife77/go-notifier/internal/handler"
	"github.com/lastlife77/go-notifier/internal/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	brok, err := rabbit.New(
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASS"),
		os.Getenv("RABBITMQ_PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(brok)

	r := mux.NewRouter()
	r.Use(middleware.Log)

	r.HandleFunc("/notify", h.CreateNotify).Methods("POST")
	r.HandleFunc("/notify/{id}", h.GetNotifyStatus).Methods("GET")
	r.HandleFunc("/notify/{id}", h.DeleteNotify).Methods("DELETE")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("NOTIFIER_PORT")), r))
}
