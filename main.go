package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/victor-renan/fileupload/internal"
)

const PORT int = 3000
const UPLOAD_PATH string = "uploads"
const MAX_SIZE int64 = 3 * (2 >> 20)

func main() {
	uploadRoute := "/uploads/"
	http.HandleFunc(uploadRoute, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodPut:
			internal.UploadConfig{
				UploadDir: UPLOAD_PATH,
				MaxSize:   MAX_SIZE,
			}.Upload(w, r)

		case http.MethodDelete:
			internal.DeleteConfig{
				UploadDir:   UPLOAD_PATH,
				DeleteRoute: uploadRoute,
			}.Delete(w, r)

		case http.MethodGet:
			internal.StaticConfig{
				Route:     uploadRoute,
				StaticDir: UPLOAD_PATH,
			}.Serve(w, r)

		}
	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)

	if err != nil {
		log.Fatalf("Error At Serve: %s\n", err.Error())
	}
}
