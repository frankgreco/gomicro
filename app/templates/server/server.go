package server

import (
    "fmt"
    "net/http"

    "github.com/gorilla/handlers"
    "<%= vcs %>/<%= repo %>/<%= project %>/route"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Start(db *database.Database, port string<%if (scheme == "https") { %>, tlsCertFilePath string, tlsPrivateKeyFilePath string<% } %>) {
    router := route.NewRouter(db)

    headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})

    <%if (scheme == "https") { %>
    err := http.ListenAndServeTLS(fmt.Sprintf(":%s", "443"), tlsCertFilePath, tlsPrivateKeyFilePath, handlers.CORS(headersOk, originsOk, methodsOk)(router))
    if err != nil {
        panic(err.Error())
    }
    <% } %>
    <%if (scheme == "http") { %>
    http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOk, originsOk, methodsOk)(router))
    <% } %>
}
