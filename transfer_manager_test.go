package main

import (
	"fmt"
	"github.com/studio-b12/gowebdav"
	"log"
	"net/url"
	"os"
	"testing"
	"time"
)

func mainTransferTest(_ *testing.T, zipped bool) Args {
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

	fmt.Println("mocking server")

	u, err := url.Parse("http://localhost:8080/")
	if err != nil {
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	fmt.Println("url: ", u.String())

	args := Args{src: "./testDir/src", duration: 3, dst: *u, user: "admin", pass: "admin", zipped: zipped}
	fmt.Println(args)

	fmt.Println("________________________________NONE_______________")

	done_files := make(chan string, 20)
	quit := make(chan int)
	pm := newTransferManager(&args, done_files)
	go pm.doWork(quit)
	time.Sleep(time.Duration(1) * time.Second)

	done_files <- "./testDir/src/e.txt"

	time.Sleep(time.Duration(1) * time.Second)

	done_files <- "./testDir/src/A"

	time.Sleep(time.Duration(1) * time.Second)

	return args

}

func TestDoWorkTransfer(t *testing.T) {
	fmt.Println("Make sure that docker container is receiving. Run: docker-compose up")
	args := mainTransferTest(t, false)
	paths := []string{"/A/B/a.txt", "/A/C/c.txt", "/A/b.txt", "/e.txt"}
	user := args.user
	password := args.pass

	c := gowebdav.NewClient(args.dst.String(), user, password)
	for _, p := range paths {
		if _, err := c.Stat(p); err != nil {
			fmt.Println(err)
			t.Errorf("File shold be received=.%s", p)
		}
	}
	err := c.RemoveAll("/")
	if err != nil {
		fmt.Println(err)
	}
}

func TestDoWorkTransferZipped(t *testing.T) {
	fmt.Println("Make sure that docker container is receiving. Run: docker-compose up")
	args := mainTransferTest(t, true)
	paths := []string{"/A.zip", "/e.txt"}
	user := args.user
	password := args.pass

	c := gowebdav.NewClient(args.dst.String(), user, password)
	for _, p := range paths {
		if _, err := c.Stat(p); err != nil {
			fmt.Println(err)
			t.Errorf("File shold be received=.%s", p)
		}
	}
	err := c.RemoveAll("/")
	if err != nil {
		fmt.Println(err)
	}
}
