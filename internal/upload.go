package internal

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

type UploadConfig struct {
	UploadDir string
	MaxSize int64
}

func (uploadDir UploadConfig) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(uploadDir.MaxSize)
	if err != nil {
		fmt.Fprintf(w, "Error At ParseMultipartForm %s", err.Error())
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		fmt.Fprintf(w, "Error At FormFile %s", err.Error())
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "Error At ReadAll %s", err.Error())
		return
	}

	hashedName := sha256.Sum256([]byte(header.Filename + time.Now().String()))

	var ext string

	for i := len(header.Filename) - 1; i >= 0; i-- {
		if header.Filename[i] == byte('.') {
			ext = header.Filename[i+1:]
			break
		}
	}

	fileName := fmt.Sprintf("%x.%s", hashedName, string(ext))
	filePath := path.Join(uploadDir.UploadDir, fileName)

	err = os.WriteFile(filePath, data, 0777)
	if err != nil {
		fmt.Fprintf(w, "Error At WriteFile %s", err.Error())
		return
	}

	fmt.Fprint(w, fileName)
}
