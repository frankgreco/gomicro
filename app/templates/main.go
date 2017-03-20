package main

import (
    "fmt"
    "os"

    "github.com/Sirupsen/logrus"
    "<%= vcs %>/<%= repo %>/<%= project %>/cmd"
)

func init() {
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
}

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
