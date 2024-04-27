package routes

import (
	"net/http"
	handler "webserver/src/app/handler"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	http.Handle("/", r)
	println("Listening on 8080")
	http.ListenAndServe(":8080", r)

}
