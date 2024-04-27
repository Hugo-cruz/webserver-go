package routes

import (
	"net/http"
	handler "webserver/src/app/handler"
	dep "webserver/src/dependencies"

	"github.com/gorilla/mux"
)

func Router(dependencies dep.Dependencies) {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/data", handler.HomeHandler)
	http.Handle("/", r)
	println("Listening on 8080")
	http.ListenAndServe(":8080", r)

}
