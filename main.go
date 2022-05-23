package main

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

// init initializes the logger.
func init() {

	logFile, err := os.OpenFile("efw_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// main starts the ELN file watcher. See README for more information.
func main() {
	now := time.Now()
	InfoLogger.Println("Starting at ", now.Format(time.RFC822))
	defer (func() {
		now := time.Now()
		InfoLogger.Println("Done at ", now.Format(time.RFC822))
	})()

	// Chain as communication channel between the file watcher and the transfer manager
	done_files := make(chan string, 20)
	// For potential (not jet implemented quit conditions)
	quit := make(chan int)
	args := GetCmdArgs()
	InfoLogger.Println("CMD Args: ", args)
	pm := newProcessManager(&args, done_files)
	go pm.doWork(quit)
	tm := newTransferManager(&args, done_files)
	go tm.doWork(quit)

	for {
		time.Sleep(args.duration * 20)
	}
}
