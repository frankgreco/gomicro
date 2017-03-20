package route

import (
    "net/http"

    "<%= vcs %>/<%= repo %>/<%= project %>/handler"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
    HandlerFunc func(*database.Database, http.ResponseWriter, *http.Request) error
}

type Routes []Route

var resourceRoutes = Routes{
	Route{
		"Retrieve<%= nounPluralUpper %>",
		"GET",
		"/<%= nounPluralLower %>",
		handler.Retrieve<%= nounPluralUpper %>,
	},
	Route{
		"Create<%= nounPluralUpper %>",
		"POST",
		"/<%= nounPluralLower %>",
		handler.Create<%= nounSingularUpper %>,
	},
    Route{
		"Delete<%= nounPluralUpper %>",
		"DELETE",
		"/<%= nounPluralLower %>",
		handler.Delete<%= nounPluralUpper %>,
	},
    Route{
		"Retrieve<%= nounSingularUpper %>",
		"GET",
		"/<%= nounSingularLower %>/{id:\\w+}",
		handler.Retrieve<%= nounSingularUpper %>,
	},
    Route{
		"Update<%= nounSingularUpper %>",
		"PUT",
		"/<%= nounSingularLower %>/{id:\\w+}",
		handler.Update<%= nounSingularUpper %>,
	},
    Route{
		"Delete<%= nounSingularUpper %>",
		"DELETE",
		"/<%= nounSingularLower %>/{id:\\w+}",
		handler.Delete<%= nounSingularUpper %>,
	},
    Route{
		"Health",
		"GET",
		"/health",
		handler.Health,
	},
	Route{
		"Docs",
		"GET",
		"/docs",
		handler.Docs,
	},
}
