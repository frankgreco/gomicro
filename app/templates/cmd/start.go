package cmd

import (
    "github.com/spf13/cobra"
    "<%= vcs %>/<%= repo %>/<%= project %>/utils"
    "<%= vcs %>/<%= repo %>/<%= project %>/server"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func init() {
    startCmd.Flags().String("app-insecure-port", "8080", "insecure application port")
    startCmd.Flags().String("app-secure-port", "443", "secure application port")
    <%if (db != "sqlite") { %>
    startCmd.Flags().String("db-port", "3306", "database port")
    startCmd.Flags().String("db-host", "127.0.0.1", "database hostname")
    startCmd.Flags().String("db-user", "admin", "database username")
    startCmd.Flags().String("db-pass", "password", "database password")
    <% } %>
    startCmd.Flags().String("db-name", "calls", "database name")
    <%if (db == "sqlite") { %>
    startCmd.Flags().String("db-location", "./data.db", "database location")
    <% } %>

    RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "start a new server",
    Long:  `start a new server`,
    Run: func(cmd *cobra.Command, args []string) {
        appPort, err := utils.GetPriorityFlagValue(cmd.Flags(), "app-port")
        <%if (db != "sqlite") { %>
        dbPort, err := utils.GetPriorityFlagValue(cmd.Flags(), "db-port")
        dbHost, err := utils.GetPriorityFlagValue(cmd.Flags(), "db-host")
        dbUser, err := utils.GetPriorityFlagValue(cmd.Flags(), "db-user")
        dbPass, err := utils.GetPriorityFlagValue(cmd.Flags(), "db-pass")
        <% } %>
        dbName, err := utils.GetPriorityFlagValue(cmd.Flags(), "db-name")
        <%if (db == "sqlite") { %>
        dbLocation, err := utils.GetPriorityFlagValue(cmd.Flags(), "db-location")
        <% } %>
        if err != nil{
            panic(err.Error())
        }
        db := &database.Database{
            <%if (db != "sqlite") { %>
            Host: dbHost,
            Port: dbPort,
            User: dbUser,
            Pass: dbPass,
            <% } %>
            Name: dbName,
            <%if (db == "sqlite") { %>
            Location: dbLocation
            <% } %>
        }
        server.Start(db, appPort)
    },
}
