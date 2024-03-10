package internal

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type StaticConfig struct {
	Route string
	StaticDir string
}


func (sc StaticConfig) Serve(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == sc.Route {
		fmt.Fprint(w, "Browse files typing the filename in the url")
		return
	}

	fileName := strings.TrimPrefix(r.URL.Path, sc.Route)

	http.ServeFile(w, r, path.Join(sc.StaticDir, fileName))
}