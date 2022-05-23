package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

var regex_url = regexp.MustCompile("/upload/(.+)")

func main() {
	handler := RegexpHandler{}
	handler.HandleFunc(regex_url, createImage)
	if err := os.MkdirAll("./target/", os.ModePerm); err != nil {
		log.Fatal(err)
		return
	}

	err := http.ListenAndServe(":8080", &handler)
	if err != nil {
		return
	}
}

func createImage(w http.ResponseWriter, request *http.Request) {
	path_list := regex_url.FindAllStringSubmatch(request.RequestURI, -1)
	if len(path_list) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := request.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, _, err := request.FormFile("upload")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	upload_flie_path := filepath.Join("target", path_list[0][1])
	println(upload_flie_path)

	if err := os.MkdirAll(filepath.Dir(upload_flie_path), os.ModePerm); err != nil {
		log.Fatal(err)
		return
	}

	tmpfile, err := os.Create(upload_flie_path)
	defer func(tmpfile *os.File) {
		err := tmpfile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(tmpfile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return
}
