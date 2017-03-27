package server

import (
    "fmt"
    "net/http"

    "github.com/gorilla/handlers"
    "<%= vcs %>/<%= repo %>/<%= project %>/route"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Start(db *database.Database, port string, tlsCertFilePath string, tlsPrivateKeyFilePath string) {
    router := route.NewRouter(db)

    headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})

    if tlsCertFilePath == "" || tlsPrivateKeyFilePath == "" {
        http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOk, originsOk, methodsOk)(router))
    } else {
        err := http.ListenAndServeTLS(fmt.Sprintf(":%s", port), tlsCertFilePath, tlsPrivateKeyFilePath, handlers.CORS(headersOk, originsOk, methodsOk)(router))
        if err != nil {
            panic(err.Error())
        }
    }
}
