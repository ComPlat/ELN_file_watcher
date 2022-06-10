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
	args        Args
)

// init initializes the logger and parses CMD args.
func init() {
	args = GetCmdArgs()

	logFile, err := os.OpenFile("efw_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	InfoLogger = log.New(mw, "\rINFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "\rERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
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
	InfoLogger.Printf("\n-----------------------------\nCMD Args:\n dst=%s,\n src=%s,\n duration=%d sec.,\n user=%s,\n zip=%t\n-----------------------------\n", args.dst.String(), args.src, int(args.duration.Seconds()), args.user, args.zipped)
	pm := newProcessManager(&args, done_files)
	go pm.doWork(quit)
	tm := newTransferManager(&args, done_files)
	go tm.doWork(quit)

	for {
		time.Sleep(args.duration * 20)
	}
}
