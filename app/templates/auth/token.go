package auth

import (
    "os"
    "io/ioutil"
    "strings"
    "encoding/csv"
)

var tokens []Token

type Token struct  {
    Token   string
}

func (token *Token) CheckCreds() bool {
    for _,cred := range tokens {
        if cred.Equals(token) {
            return true
        }
    }
    return false
}

func (one *Token) Equals(two *Token) bool {
    return one.Token == two.Token
}

func InitTokens(path string) error {
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
    tokens = []Token{}
    for _,r := range records {
        token := Token{}
        token.Token = r[0]
        tokens = append(tokens, token)
    }
    return nil
}
