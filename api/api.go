package main

import (
    "net/http"
)

func V1PostConversationHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"alive": true}`))
}
