package models

type <%= nounSingularUpper %> struct {
	Id         int64    `json:"id,omitempty"`
	ParamOne   string   `json:"paramOne"`
    ParamTwo   string   `json:"paramTwo"`
}

type <%= nounPluralUpper %> []<%= nounSingularUpper %>
