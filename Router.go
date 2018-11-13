package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func index_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome home!")
}

func about_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "About us!")
}

func money_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "MOONEY")
}

func main() {
	//initiate the variable
	route := mux.NewRouter()
	route.HandleFunc("/", index_handler)
	route.HandleFunc("/about", about_handler)
	route.HandleFunc("/money", money_handler)

}
