package server

import (
    "fmt"
    "net/http"

    "<%= vcs %>/<%= repo %>/<%= project %>/route"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Start(db *database.Database, port string) {
    router := route.NewRouter(db)
    http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
