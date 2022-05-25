package main

import (
	//"github.com/bouk/monkey"
	"log"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestDoWorkProcess(t *testing.T) {
	cleanTestDir()
	defer cleanTestDir()

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

	args := Args{src: "/home/martin/Desktop/dev/KIT/ELN_file_watcher/test_dir", duration: 3, url: url.URL{
		Scheme:     "",
		Opaque:     "",
		User:       nil,
		Host:       "",
		Path:       "",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}, zipped: true}

	//start_time := time.Now()

	done_files := make(chan string, 20)
	quit := make(chan int)
	pm := newProcessManager(&args, done_files)
	go pm.doWork(quit)
	quit <- 1
	if len(done_files) > 0 {
		t.Errorf("Done files channel shoud be empty but len(done_files)=%d", len(done_files))
	}

	//wayback := start_time.Add(time.Duration(3) * time.Second)
	//patch := monkey.Patch(time.Now, func() time.Time { return wayback })
	//defer patch.Unpatch()
	time.Sleep(time.Duration(2900) * time.Millisecond)
	go pm.doWork(quit)
	quit <- 1
	if len(done_files) > 0 {
		t.Errorf("Done files channel shoud be empty but len(done_files)=%d", len(done_files))
	}
	time.Sleep(time.Duration(100) * time.Millisecond)

	go pm.doWork(quit)

	if len(done_files) > 3 {
		t.Errorf("Done files channel shoud have > 3 elements. len(done_files)=%d", len(done_files))
	}

	quit <- 1
}
