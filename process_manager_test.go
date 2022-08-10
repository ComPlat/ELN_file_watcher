package main

import (
	//"github.com/bouk/monkey"
	"log"
	"os"
	"testing"
	"time"
)

func TestDoWorkProcess(t *testing.T) {
	cleanTestDir()
	defer cleanTestDir()

	// Prepare Test
	if err := os.MkdirAll("testDir/src/A/B", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("testDir/src/A/C", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("testDir/src/C", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	writeIntoFile("testDir/src/A/B/a.txt", "Hallo A_B_a")
	writeIntoFile("testDir/src/A/b.txt", "Hallo A_c")
	writeIntoFile("testDir/src/A/C/c.txt", "Hallo A_C_c")
	writeIntoFile("testDir/src/C/d.txt", "Hallo C_d")
	writeIntoFile("testDir/src/e.txt", "Hallo e")

	args := Args{src: "/home/martin/Desktop/dev/KIT/ELN_file_watcher/testDir/src", duration: 3, user: "admin", pass: "admin", sendType: "zip"}

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
