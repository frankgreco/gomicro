package models

type <%= nounSingularUpper %> struct {
    ID         uint64   `json:"id,omitempty"`
	ParamOne   string   `json:"paramOne"`
    ParamTwo   string   `json:"paramTwo"`
}

type <%= nounPluralUpper %> []<%= nounSingularUpper %>

func (<%= nounSingularUpper %>) TableName() string {
	return "<%= nounPluralLower %>"
}
