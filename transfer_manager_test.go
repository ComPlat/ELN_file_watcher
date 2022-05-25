package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	//"github.com/bouk/monkey"
	"log"
	"net/url"
	"os"
	"testing"
)

type Base int

const (
	E Base = iota
	A
	C
	NONE
)

func getFile(request *http.Request, post_name *string) (string, int) {
	if err := request.ParseMultipartForm(32 << 20); err != nil {
		ErrorLogger.Println(err)
		return "", http.StatusBadRequest
	}
	//Access the photo key - First Approach
	_, file, err := request.FormFile(*post_name)
	if err != nil {
		ErrorLogger.Println(err)
		return "", http.StatusBadRequest
	}

	return file.Filename, http.StatusOK
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func mainTransferTest(t *testing.T, zipped bool) {
	cleanTestDir()
	defer cleanTestDir()
	counter := E
	post_name := "file"
	// Prepare Test
	if err := os.MkdirAll("test_dir/A/B", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("test_dir/A/C", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("test_dir/C", 0777); err != nil {
		log.Fatal(err)
	}

	writeIntoFile("test_dir/A/B/a.txt", "Hallo A_B_a")
	writeIntoFile("test_dir/A/b.txt", "Hallo A_c")
	writeIntoFile("test_dir/A/C/c.txt", "Hallo A_C_c")
	writeIntoFile("test_dir/C/d.txt", "Hallo C_d")
	writeIntoFile("test_dir/e.txt", "Hallo e")

	fmt.Println("mocking server")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inn Request .... ")
		fmt.Println("Received:", r.URL)
		filename, statusCode := getFile(r, &post_name)
		switch counter {
		case A:
			if statusCode != http.StatusOK {
				t.Errorf("Status shuold be Ok but it is %d", statusCode)
			}
			if zipped {
				if filename != "A.zip" {
					t.Errorf("File shuold be A.zip but it is %s", filename)
				}
			} else {
				if !contains([]string{"b.txt", "a.txt", "c.txt"}, filename) {
					t.Errorf("File shuold be in [b.txt,a.txt,c.txt] but it is %s", filename)
				}
			}
			break
		case E:
			if statusCode != http.StatusOK {
				t.Errorf("Status shuold be Ok but it is %d", statusCode)
			}
			if filename != "e.txt" {
				t.Errorf("File shuold be e.txt but it is %s", filename)
			}
			break
		case C:
			break
		case NONE:
			t.Errorf("No Reques expexted!")
			break

		}

	}))

	u, err := url.Parse(ts.URL)
	if err != nil {
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	fmt.Println("url: ", u.String())

	args := Args{src: "/home/martin/Desktop/dev/KIT/ELN_file_watcher/test_dir", duration: 3, url: *u, zipped: zipped, post_name: post_name}
	fmt.Println(args)
	defer ts.Close()

	fmt.Println("________________________________NONE_______________")

	counter = NONE
	done_files := make(chan string, 20)
	quit := make(chan int)
	pm := newTransferManager(&args, done_files)
	go pm.doWork(quit)
	time.Sleep(time.Duration(1) * time.Second)

	counter = E
	done_files <- "/home/martin/Desktop/dev/KIT/ELN_file_watcher/test_dir/e.txt"

	time.Sleep(time.Duration(1) * time.Second)

	counter = A
	done_files <- "/home/martin/Desktop/dev/KIT/ELN_file_watcher/test_dir/A"

	time.Sleep(time.Duration(1) * time.Second)
}

func TestDoWorkTransfer(t *testing.T) {
	mainTransferTest(t, false)
}

func TestDoWorkTransferZipped(t *testing.T) {
	mainTransferTest(t, true)
}
