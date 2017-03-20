package cmd

import(
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "<%= project %>",
    Short: "Hugo is a very fast static site generator",
    Long: `<%= project %> is a RESTful microservice that performs CRUD operations on the resource Person
    Complete documentation can be found at https://github.com/frankgreco/gohttp`,
}
