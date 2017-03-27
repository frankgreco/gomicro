package auth

import (
    "os"
    "io/ioutil"
    "strings"
    "encoding/csv"
)

var basics []Basic

type Basic struct  {
    Username string
    Password string
}

func (basic *Basic) CheckCreds() bool {
    for _,cred := range basics {
        if cred.Equals(basic) {
            return true
        }
    }
    return false
}

func (one *Basic) Equals(two *Basic) bool {
    return (one.Username == two.Username) && (one.Password == two.Password)
}

func InitBasic(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    data, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }
    file.Close()
    reader := csv.NewReader(strings.NewReader(string(data)))
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }
    basics = []Basic{}
    for _,r := range records {
        basic := Basic{}
        basic.Username = r[0]
        basic.Password = r[1]
        basics = append(basics, basic)
    }
    return nil
}
