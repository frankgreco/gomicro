package models

<%if (db == "mongodb") { %>
import (
    "gopkg.in/mgo.v2/bson"
)
<% } %>

type <%= nounSingularUpper %> struct {
	<%if (db == "mysql") { %>
	Id         int64    `json:"id,omitempty"`
	ParamOne   string   `json:"paramOne"`
    ParamTwo   string   `json:"paramTwo"`
	<% } %>
	<%if (db == "mongodb") { %>
	Id         bson.ObjectId    `bson:"id" json:"id,omitempty"`
	ParamOne   string   		`bson:"paramOne" json:"paramOne"`
    ParamTwo   string   		`bson:"paramTwo" json:"paramTwo"`
	<% } %>
}

type <%= nounPluralUpper %> []<%= nounSingularUpper %>
