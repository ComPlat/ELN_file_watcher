package main

import (
	"flag"
	"log"
	"net/url"
	"time"
)

type Args struct {
	src      string
	duration time.Duration
	url      url.URL
	zipped   bool
}

// GetCmdArgs Get/Parse command line arguments manager
func GetCmdArgs() Args {
	var fp, url_path string
	var duration int
	var zipped bool
	flag.StringVar(&fp, "src", "", "Src directory to be watched")
	flag.StringVar(&url_path, "url", "", "HTTP url to the file network storage. For example: http://<ip address>:<port>/<upload path>/")
	flag.IntVar(&duration, "duration", 300, "Duration in seconds, i.e., how long a file must not be changed before sent")
	/// Only considered if result are stored in a folder.
	/// If zipped is set the result folder will be transferred as zip file
	flag.BoolVar(&zipped, "zip", false, "Only considered if result are stored in a folder. If zipped is set the result folder will be transferred as zip file")
	flag.Parse()

	if url_path == "" || fp == "" {
		err := "'url' and 'src' must not be empty!"
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	u, err := url.Parse(url_path)
	if err != nil {
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	return Args{src: fp, url: *u, duration: time.Duration(duration) * time.Second}

}
