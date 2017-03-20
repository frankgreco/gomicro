package server

import (
    "fmt"
    "net/http"

    "github.com/gorilla/handlers"
    "<%= vcs %>/<%= repo %>/<%= project %>/route"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Start(db *database.Database, port string) {
    router := route.NewRouter(db)

    headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})

    http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOk, originsOk, methodsOk)(router))
}
