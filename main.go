package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	args        Args
	tr          *http.Transport
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

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	if len(args.crt) > 0 {
		certs, err := ioutil.ReadFile(args.crt)
		if err != nil {
			ErrorLogger.Fatalf("Failed to append %q to RootCAs: %v", args.crt, err)
		}

		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			ErrorLogger.Println("No certs appended, using system certs only")
		}
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            rootCAs,
	}

	tr = &http.Transport{TLSClientConfig: config}
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
	InfoLogger.Printf("\n-----------------------------\nCMD Args:\n dst=%s,\n src=%s,\n duration=%d sec.,\n user=%s,\n zip=%t,\n crt= %s \n-----------------------------\n", args.dst.String(), args.src, int(args.duration.Seconds()), args.user, args.zipped, args.crt)
	pm := newProcessManager(&args, done_files)
	go pm.doWork(quit)
	tm := newTransferManager(&args, done_files)
	go tm.doWork(quit)

	for {
		time.Sleep(args.duration * 20)
	}
}
