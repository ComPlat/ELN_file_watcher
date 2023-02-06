package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"time"
)

var (
	InfoLogger        *log.Logger
	ErrorLogger       *log.Logger
	args              Args
	tr                *http.Transport = nil
	TempPath, LogPath string
)

// init initializes the logger and parses CMD args.
func init() {
	args = GetCmdArgs()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	wd := usr.HomeDir
	ewfFoldername := fmt.Sprintf("efw_exporter_%s", args.name)
	TempPath = path.Join(wd, ewfFoldername+"/efw_temp")
	if err := os.MkdirAll(TempPath, os.ModePerm); err != nil {
		ErrorLogger.Println(err)
		panic("")
	}

	LogPath = path.Join(wd, ewfFoldername+"/efw_log.txt")

	logFile, err := os.OpenFile(LogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	InfoLogger = log.New(mw, "-> INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Read in the cert file
}

func initArgs() {
	isCert := len(args.crt) > 0

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	if isCert {
		certs, err := ioutil.ReadFile(args.crt)
		if err == nil {
			if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
				ErrorLogger.Println("No certs appended, using system certs only")
			}
		}
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: true,
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
	InfoLogger.Printf("\n-----------------------------\nLogfile: %s\n-----------------------------\nCMD Args:\n name=%s,\n dst=%s,\n src=%s,\n duration=%d sec.,\n user=%s,\n type=%s,\n transfer=%s\n-----------------------------\n", LogPath, args.name, args.dst.String(), args.src, int(args.duration.Seconds()), args.user, args.sendType, args.tType)
	pm := newProcessManager(&args, done_files)
	go pm.doWork(quit)

	prm := newPrepareManager(&args, done_files)
	go prm.doWork(quit)

	tm := newTransferManager(&args)
	if err := tm.connect_to_server(); err != nil {
		ErrorLogger.Println("Error connecting: ", err)
		log.Fatal(err)
	}
	go tm.doWork(quit)

	for {
		time.Sleep(args.duration * 20)
	}
}
