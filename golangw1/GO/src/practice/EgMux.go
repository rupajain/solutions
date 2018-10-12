package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "log"
    "net/http"
)

func main() {
    port := ":8082"
    var router = mux.NewRouter()
    router.HandleFunc("/m/{msg}", handleMessage).Methods("GET")
    router.HandleFunc("/n/{num}", handleNumber).Methods("GET")

    headersOk := handlers.AllowedHeaders([]string{"Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

    fmt.Printf("Server is running at http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    message := vars["msg"]
    response := map[string]string{"message": message}
    w.Header().Set("Content-Type", "application/json") // this
    json.NewEncoder(w).Encode(response)
}

func handleNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    number := vars["num"]
    response := map[string]string{"number": number}
    w.Header().Set("Content-Type", "application/json") // and this
    json.NewEncoder(w).Encode(response)
}