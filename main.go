package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "reflect"
)

type JsonResponse struct {
    Data map[string]interface{} `json:"data"`
    Type map[string]string      `json:"type"`
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func parseJSON(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == http.MethodOptions {
        return
    }

    var input map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := JsonResponse{
        Data: input,
        Type: make(map[string]string),
    }

    for key, value := range input {
        response.Type[key] = reflect.TypeOf(value).String()
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/parse", parseJSON)
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}