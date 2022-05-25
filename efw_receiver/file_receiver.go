package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	regex_url   *regexp.Regexp
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	args        Args
)

// init initializes the logger.
func init() {

	logFile, err := os.OpenFile("efw_receiver_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	now := time.Now()
	InfoLogger.Println("Starting at ", now.Format(time.RFC822))
	defer (func() {
		now := time.Now()
		InfoLogger.Println("Done at ", now.Format(time.RFC822))
	})()

	args = GetCmdArgs()

	regex_string := strings.TrimPrefix(strings.TrimSuffix(args.url, "/"), "/")

	regex_url = regexp.MustCompile(fmt.Sprintf("/%s/(.+)", regex_string))
	handler := RegexpHandler{}
	handler.HandleFunc(regex_url, createImage)
	if err := os.MkdirAll(args.dst, os.ModePerm); err != nil {
		log.Fatal(err)
		return
	}

	InfoLogger.Println("Server running at prot ", args.port)
	InfoLogger.Printf("Upload URL  http://<LOCAL IP>%s/%s/**/*\n", args.port, regex_string)
	InfoLogger.Println("Post file variable name: ", args.post_name)
	InfoLogger.Println("Post file saved at: ", args.dst)

	err := http.ListenAndServe(args.port, &handler)
	if err != nil {
		return
	}
}

func createImage(w http.ResponseWriter, request *http.Request) {
	pathList := regex_url.FindAllStringSubmatch(request.RequestURI, -1)
	if len(pathList) < 1 {
		ErrorLogger.Println("No Files sent!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := request.ParseMultipartForm(32 << 20); err != nil {
		ErrorLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, _, err := request.FormFile(args.post_name)
	if err != nil {
		ErrorLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uploadFilePath := filepath.Join(args.dst, pathList[0][1])
	InfoLogger.Println("New File at: ", pathList[0][1])

	if err := os.MkdirAll(filepath.Dir(uploadFilePath), os.ModePerm); err != nil {
		ErrorLogger.Println(err)
		return
	}

	tmpfile, err := os.Create(uploadFilePath)
	defer func(tmpfile *os.File) {
		err := tmpfile.Close()
		if err != nil {
			ErrorLogger.Println(err)
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
