package handler

import (
	"net/http"
    "io/ioutil"

    "<%= vcs %>/<%= repo %>/<%= project %>/utils"
	"<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Health(db *database.Database, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	return nil
}

func Docs(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    file, err := ioutil.ReadFile("./swagger.json")
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(file)
	return nil
}
