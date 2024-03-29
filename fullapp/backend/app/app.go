package fullapp

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

// func main() {
// 	r := newRouter()
// 	http.ListenAndServe(":8080", r)
// }