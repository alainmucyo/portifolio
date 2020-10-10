package utils

import (
	"github.com/alainmucyo/my_brand/config"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func FileUpload(r *http.Request, requestName string) (string, error) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile(requestName) // Retrieve the file from form data

	if err != nil {
		return "", err
	}
	defer file.Close() // Close the file when we finish

	// This is path which we want to store the file
	dir := "./" + config.STATIC_FOLDER + "/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}
	fileName := strconv.Itoa(rand.Int()) + handler.Filename
	f, err := os.OpenFile(dir+fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	// Copy the file to the destination path
	io.Copy(f, file)

	return "public/" + fileName, nil
}

type MyFileSystem struct {
	http.Dir
}

func (m MyFileSystem) Open(name string) (result http.File, err error) {
	f, err := m.Dir.Open(name)
	if err != nil {
		return
	}

	fi, err := f.Stat()
	if err != nil {
		return
	}
	if fi.IsDir() {
		// Return a response that would have been if directory would not exist:
		return m.Dir.Open("does-not-exist")
	}
	return f, nil
}
