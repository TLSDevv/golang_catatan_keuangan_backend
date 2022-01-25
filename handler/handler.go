package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	//service
}

func NewAPI() *API {
	return &API{}
}

func (a API) Start(host, port string) {
	r := mux.NewRouter()

	fmt.Println(host, port)

	server := http.Server{
		Addr:    endPoint(host, port),
		Handler: r,
	}

	fmt.Printf("Listening %s to port %s", host, port)
	err := server.ListenAndServe()

	if err != nil {
		logrus.Error("error Listen serve ", err)
		return
	}
}

func endPoint(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
