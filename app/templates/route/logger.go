package route

import (
    "time"
	"net/http"

    "github.com/Sirupsen/logrus"
    h "<%= vcs %>/<%= repo %>/<%= project %>/handler"
)

func logger(inner h.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

        logrus.WithFields(logrus.Fields{
            "method": r.Method,
            "uri": r.RequestURI,
            "name": name,
            "duration": time.Since(start),
        }).Info("http server started")
	})
}
