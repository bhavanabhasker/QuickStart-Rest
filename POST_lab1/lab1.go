package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name string `json:"name"`
}
type Result struct {
	Greeting string `json:"greeting"`
}

func main() {

	httprouter := mux.NewRouter()
	var handler http.HandlerFunc
	handler = Hello
	httprouter.
		Methods("POST").
		Path("/hello").
		Name("Hello").
		Handler(handler)

	log.Fatal(http.ListenAndServe(":1234", httprouter))
}
func Hello(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	var result Result
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	result.Greeting += "Hello,"
	result.Greeting += todo.Name
	result.Greeting += "!"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

}
