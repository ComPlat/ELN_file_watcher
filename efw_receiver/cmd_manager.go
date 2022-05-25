package main

import (
	"flag"
	"log"
)

type Args struct {
	dst       string
	post_name string
	url       string
	port      string
}

// GetCmdArgs Get/Parse command line arguments manager
func GetCmdArgs() Args {
	var dst, url_path, post_name, port string

	flag.StringVar(&dst, "dst", "", "Destination directory where received files are stored.")
	flag.StringVar(&url_path, "url", "upload", "Fixed URL path to upload files <upload path>")
	flag.StringVar(&post_name, "post", "file", "The post field name by which the file will be sent")
	flag.StringVar(&port, "port", ":8080", "Server address port. Starts with leading ':'")
	flag.Parse()

	if url_path == "" || dst == "" {
		err := "'url' and 'dst' must not be empty!"
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	return Args{dst: dst, url: url_path, post_name: post_name, port: port}

}
