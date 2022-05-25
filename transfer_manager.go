package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// TransferManager reacts on the channel done_files.
// If folder of file is ready to send it sends it via HTTP to <CMD arg -url>.
// It also initializes the zipping if <CMD arg -zipped> is set.
type TransferManager struct {
	args       *Args
	done_files chan string
}

// doWork runs in a endless loop. It reacts on the channel done_files.
// If folder of file is ready to send it sends it via HTTP to <CMD arg -url>.
// It also initializes the zipping if <CMD arg -zipped> is set
// It terminates as soon as a value is pushed into quit. Run in extra goroutine.
func (m TransferManager) doWork(quit chan int) {
	InfoLogger.Println("Started transfer process.")
	for {

		select {
		case <-quit:
			InfoLogger.Println("Quit transfer process.")
			return
		case to_send := <-m.done_files:
			if file, err := os.Stat(to_send); err != nil {
				ErrorLogger.Println(err)
			} else if !file.IsDir() {
				if err := m.send_file(to_send, file); err != nil {
					ErrorLogger.Println(err)
				}
			} else if m.args.zipped {
				zip_paht, err := zipFolder(to_send)
				if err != nil {
					ErrorLogger.Println(err)
				}
				if file, err := os.Stat(zip_paht); err != nil {
					ErrorLogger.Println(err)
				} else {
					if err = m.send_file(zip_paht, file); err != nil {
						ErrorLogger.Println(err)
					}
				}

			} else {
				err := filepath.Walk(to_send, func(path string, info os.FileInfo, err error) error {
					if err == nil && !info.IsDir() {
						err = m.send_file(path, info)
					}

					return err

				})

				if err != nil {
					ErrorLogger.Println(err)
				}
			}

		}
	}
}

// send_file sends a file via HTTP
func (m TransferManager) send_file(path_to_file string, fileInfo os.FileInfo) error {
	var urlPath string
	if relpath, err := filepath.Rel(m.args.src, path_to_file); err == nil {
		relpath = strings.Replace(relpath, string(os.PathSeparator), "/", -1)
		tmpUrlPath := m.args.url
		tmpUrlPath.Path = path.Join(tmpUrlPath.Path, relpath)
		urlPath = tmpUrlPath.String()
		InfoLogger.Println("Sending...", relpath)
	} else {
		return err
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile(m.args.post_name, fileInfo.Name())
	if err != nil {
		return err
	}
	file, err := os.Open(path_to_file)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	if err = writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
func newTransferManager(args *Args, done_files chan string) TransferManager {
	return TransferManager{args: args, done_files: done_files}
}
