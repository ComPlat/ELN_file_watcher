package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	args        Args
	tr          *http.Transport = nil
	TempPath    string
)

// init initializes the logger and parses CMD args.
func init() {

	logFile, err := os.OpenFile("efw_log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	InfoLogger = log.New(mw, "-> INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Read in the cert file
}

func initArgs() {
	args = GetCmdArgs()
	isCert := len(args.crt) > 0

	executablePath, err := os.Executable()
	if err != nil {
		ErrorLogger.Println(err)
		panic("")
	}

	TempPath = path.Join(path.Dir(executablePath), ".temp")
	_ = os.MkdirAll(TempPath, os.ModePerm)

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	if isCert {
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
		InsecureSkipVerify: !isCert,
		RootCAs:            rootCAs,
	}

	tr = &http.Transport{TLSClientConfig: config}

}

// main starts the ELN file watcher. See README for more information.
func main() {
	initArgs()
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
	InfoLogger.Printf("\n-----------------------------\nCMD Args:\n dst=%s,\n src=%s,\n duration=%d sec.,\n user=%s,\n type=%s,\n crt= %s \n-----------------------------\n", args.dst.String(), args.src, int(args.duration.Seconds()), args.user, args.sendType, args.crt)
	pm := newProcessManager(&args, done_files)
	go pm.doWork(quit)

	prm := newPrepareManager(&args, done_files)
	go prm.doWork(quit)

	tm := newTransferManager(&args)
	if _, err := tm.connect_to_server(); err != nil {
		ErrorLogger.Println("Error connecting: ", err)
		log.Fatal(err)
	}
	go tm.doWork(quit)

	for {
		time.Sleep(args.duration * 20)
	}
}
