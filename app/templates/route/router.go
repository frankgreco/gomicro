package route

import (
	"net/http"

	"github.com/gorilla/mux"
    h "<%= vcs %>/<%= repo %>/<%= project %>/handler"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func NewRouter(db *database.Database) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range resourceRoutes {
		var handler http.Handler

        handler = logger(h.Handler{db, route.HandlerFunc}, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
