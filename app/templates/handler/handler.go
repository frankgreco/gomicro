package handler

import (
    "net/http"
    "encoding/json"

    "github.com/Sirupsen/logrus"
    "<%= vcs %>/<%= repo %>/<%= project %>/utils"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

type Handler struct {
	*database.Database
	H func(db *database.Database, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Database, w, r)
	if err != nil {
        w.Header().Set("Content-Type", "application/json")
		switch e := err.(type) {
		case utils.Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
            logrus.WithFields(logrus.Fields{
                "method": r.Method,
                "uri": r.RequestURI,
            }).Error(e.Error())
            w.WriteHeader(e.Status())
            if err := json.NewEncoder(w).Encode(utils.JsonErr{e.Status(), e.Error()}); err != nil {
                panic(err.Error())
            }
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
            logrus.WithFields(logrus.Fields{
                "method": r.Method,
                "uri": r.RequestURI,
            }).Error("unknown error")
            w.WriteHeader(http.StatusInternalServerError)
            if err := json.NewEncoder(w).Encode(utils.JsonErr{http.StatusInternalServerError, "unknown error"}); err != nil {
                panic(err.Error())
            }
		}
	}
}
