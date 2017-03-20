package utils

import (
    "os"
    "strings"

    flag "github.com/spf13/pflag"
)

func GetPriorityFlagValue(flags *flag.FlagSet, name string) (string, error) {
    if(strings.Compare(os.Getenv(fromFlagToEnv(name)), "") != 0) {
        return os.Getenv(fromFlagToEnv(name)), nil
    }
    return flags.GetString(name)
}

func fromFlagToEnv(name string) string {
    return strings.Replace(strings.ToUpper(name), "-", "_", -1)
}
