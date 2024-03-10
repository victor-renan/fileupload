package internal

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

type DeleteConfig struct {
	DeleteRoute string
	UploadDir string
}

func (dc DeleteConfig) Delete(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == dc.DeleteRoute {
		fmt.Fprint(w, "Type the file name")
		return
	}

	fileName := strings.TrimPrefix(r.URL.Path, dc.DeleteRoute)

	err := os.Remove(path.Join(dc.UploadDir, fileName))
	if (err != nil) {
		fmt.Fprintf(w, "Error At Remove %s",	err.Error())
		return
	}
}
