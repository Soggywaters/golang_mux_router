package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func index_handler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Welcome home!")
}


func user_handler(writer http.ResponseWriter, request *http.Request) {
	param_vars := mux.Vars(request)
	writer.WriteHeader(http.StatusOK) //HTTP status

	fmt.Fprintf(writer, "Hello user number %v", param_vars["id"])
}

func main() {
	//Initiate the variable
	route := mux.NewRouter()

	route.HandleFunc("/", index_handler)
	route.HandleFunc("/user/{id}", user_handler)
	http.Handle("/", route)
	http.ListenAndServe(":8000", nil)
}
