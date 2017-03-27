package handler

import (
    "io"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"

    <%if (auth) { %>"<%= vcs %>/<%= repo %>/<%= project %>/auth"<% } %>
    "<%= vcs %>/<%= repo %>/<%= project %>/utils"
    "<%= vcs %>/<%= repo %>/<%= project %>/models"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Retrieve<%= nounPluralUpper %>(db *database.Database, w http.ResponseWriter, r *http.Request) error {

    <%if (auth) { %>
    if(!auth.Check(w, r)) {
        return utils.StatusError{http.StatusUnauthorized, errors.New("unauthorized")}
    }
    <% } %>

    result, err := db.RetrieveAll()
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("Content-Type", "application/json")
    if result != nil {
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(result); err != nil {
            return utils.StatusError{http.StatusInternalServerError, err}
        }
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
    return nil
}

func Create<%= nounSingularUpper %>(db *database.Database, w http.ResponseWriter, r *http.Request) error {

    <%if (auth) { %>
    if(!auth.Check(w, r)) {
        return utils.StatusError{http.StatusUnauthorized, errors.New("unauthorized")}
    }
    <% } %>

    <%= nounSingularLower %> := models.<%= nounSingularUpper %>{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
	if err := r.Body.Close(); err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
	}
	if err := json.Unmarshal(body, &<%= nounSingularLower %>); err != nil {
        return utils.StatusError{http.StatusUnprocessableEntity, err}
	}
    result, err := db.Create(&<%= nounSingularLower %>)
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("location", fmt.Sprintf("/<%= nounSingularLower %>/%s", result.ID))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    return nil
}

func Delete<%= nounPluralUpper %>(db *database.Database, w http.ResponseWriter, r *http.Request) error {

    <%if (auth) { %>
    if(!auth.Check(w, r)) {
        return utils.StatusError{http.StatusUnauthorized, errors.New("unauthorized")}
    }
    <% } %>

    if _, err := db.DeleteAll(); err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNoContent)
    return nil
}
