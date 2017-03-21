package cmd

import(
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "<%= project %>",
    Short: "<%= project %> is a RESTful microservice",
    Long: `<%= project %> is a RESTful microservice that performs CRUD operations on the <%= nounSingularUpper %> resources`,
}
