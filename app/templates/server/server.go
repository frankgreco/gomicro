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

    <%if https { %>
    http.ListenAndServeTLS(fmt.Sprintf(":%s", ":443"), "certs/server.crt", "certs/server.key", handlers.CORS(headersOk, originsOk, methodsOk)(router))
    http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", port), handlers.CORS(headersOk, originsOk, methodsOk)(router))
    <% } %>
    <%if http { %>
    http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOk, originsOk, methodsOk)(router))
    <% } %>

}
