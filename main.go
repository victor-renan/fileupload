package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

const PORT int = 3000
const MAX_SIZE int64 = 3 * (2 >> 20)
const UPLOAD_PATH string = "uploads"

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			fmt.Fprintf(w, "Only PUT method accepted")
			return
		}

		err := r.ParseMultipartForm(MAX_SIZE)
		if err != nil {
			fmt.Fprintf(w, "At ParseMultipartForm %s", err.Error())
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			fmt.Fprintf(w, "At FormFile %s", err.Error())
			return
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(w, "At ReadAll %s", err.Error())
			return
		}

		hashedName := sha256.Sum256([]byte(header.Filename))
		
		var ext string

		for i := len(header.Filename) - 1; i >= 0; i-- {
			if (header.Filename[i] == byte('.')) {
				ext = header.Filename[i+1:]
				break
			}
		}

		filePath := path.Join(UPLOAD_PATH, fmt.Sprintf("%x.%s", hashedName, string(ext)))

		os.WriteFile(filePath, data, 0777)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)

	if err != nil {
		log.Fatalf("Serve: %s\n", err.Error())
	}
}
