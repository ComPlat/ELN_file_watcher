package main

import (
	"github.com/StarmanMartin/gowebdav"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// TransferManager reacts on the channel done_files.
// If folder of file is ready to send it sends it via WebDAV (HTTP) to <CMD arg -dst>.
// It also initializes the zipping if <CMD arg -zip> is set.
type TransferManager struct {
	args       *Args
	done_files chan string
}

// doWork runs in a endless loop. It reacts on the channel done_files.
// If folder of file is ready to send it sends it via HWebDAV (HTTP) to <CMD arg -dst>.
// It also initializes the zipping if <CMD arg -zip> is set
// It terminates as soon as a value is pushed into quit. Run in extra goroutine.
func (m *TransferManager) doWork(quit chan int) {
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
					if err := os.Remove(zip_paht); err != nil {
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

func (m *TransferManager) connect_to_server() (*gowebdav.Client, error) {
	user := m.args.user
	password := m.args.pass

	c := gowebdav.NewClient(m.args.dst.String(), user, password, tr)
	c.SetTimeout(5 * time.Second)
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c, nil
}

// send_file sends a file via WebDAV
func (m *TransferManager) send_file(path_to_file string, _ os.FileInfo) error {
	var webdavFilePath, urlPathDir string

	c, err := m.connect_to_server()
	if err != nil {
		return err
	}
	if relpath, err := filepath.Rel(m.args.src, path_to_file); err == nil {
		webdavFilePath = strings.Replace(relpath, string(os.PathSeparator), "/", -1)
		webdavFilePath = strings.TrimPrefix(webdavFilePath, "./")
		urlPathDir = filepath.Dir(webdavFilePath)
		InfoLogger.Println("Sending...", relpath)
	} else {
		return err
	}

	if urlPathDir != "." {
		err := c.MkdirAll(urlPathDir, 0644)
		if err != nil {
			return err
		}
	}
	bytes, err := ioutil.ReadFile(path_to_file)
	if err != nil {
		return err
	}

	err = c.Write(webdavFilePath, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
func newTransferManager(args *Args, done_files chan string) TransferManager {
	return TransferManager{args: args, done_files: done_files}
}
