package auth

import (
    "strings"
    "errors"
	"net/http"
    "encoding/base64"
)

func extractAuthHeader(r *http.Request) ([]string, error) {
    s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
    if len(s) != 2 {
        return nil, errors.New("authorization header not properly formatted")
    }
    return s, nil
}

func decode(s string) (*string, error) {
    b, err := base64.StdEncoding.DecodeString(s)
    if err != nil {
        return nil, errors.New("authorization header not properly formatted")
    }
    c := string(b)
    return &c, nil
}

func Check(w http.ResponseWriter, r *http.Request) bool {
    s, err := extractAuthHeader(r)
    if err != nil {
        return false
    }
    switch s[0] {
    case "Basic":
        b, err := decode(s[1])
        if err != nil {
            return false
        }
        pair := strings.SplitN(*b, ":", 2)
        if len(pair) != 2 {
            return false
        }
        basic := &Basic{
            Username: pair[0],
            Password: pair[1],
        }
        return basic.CheckCreds()
    case "Bearer":
        token := &Token{
            Token: s[1],
        }
        return token.CheckCreds()
    default:
        return false
    }
}
