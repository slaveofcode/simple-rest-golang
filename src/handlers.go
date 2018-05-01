package main

import "html"
import "encoding/json"
import "fmt"
import "net/http"
import "github.com/gorilla/mux"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Saving money"},
		Todo{Name: "Buying House"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoId := vars["TodoId"]
	fmt.Fprintln(w, "Todo Show: ", todoId)
}