package main

import (
	"flag"
	"log"
	"net/url"
	"strconv"
	"time"
)

type Args struct {
	src, user, pass, crt string
	dst                  url.URL
	duration             time.Duration
	sendType             string
	tType, name          string
}

// GetCmdArgs Get/Parse command line arguments manager
func GetCmdArgs() Args {
	var fp, dst, user, pass, crt, durationStr, tType, name string
	var duration int
	var sendType string
	var err error

	flag.StringVar(&fp, "src", "{{ src }}", "Source directory to be watched.")
	flag.StringVar(&name, "name", "{{ name }}", "Name of the EFW instance.")
	flag.StringVar(&dst, "dst", "{{ dst }}", "WebDAV destination URL. If the destination is on the lsdf, the URL should be as follows:\nhttps://os-webdav.lsdf.kit.edu/<OE>/<inst>/projects/<PROJECTNAME>/\n            <OE>-Organisationseinheit, z.B. kit.\n            <inst>-Institut-Name, z.B. ioc, scc, ikp, imk-asf etc.\n            <USERNAME>-User-Name z.B. xy1234, bs_abcd etc.\n            <PROJRCTNAME>-Projekt-Name")
	flag.StringVar(&user, "user", "{{ user }}", "WebDAV or SFTP user")
	flag.StringVar(&pass, "pass", "{{ password }}", "WebDAV or SFTP Password")
	flag.StringVar(&durationStr, "duration", "{{ duration }}", "Duration in seconds, i.e., how long a file must not be changed before sent.")
	flag.StringVar(&crt, "crt", "{{ crt }}", "Path to server TLS certificate. Only needed if the server has a self signed certificate.")
	/// Only considered if result are stored in a folder.
	/// If zipped is set the result folder will be transferred as zip file
	flag.StringVar(&sendType, "type", "{{ type }}", "Type must be 'file', 'folder' or 'zip'. The 'file' option means that each file is handled individually, the 'folder' option means that entire folders are transmitted only when all files in them are ready. The option 'zip' sends a folder zipped, only when all files in a folder are ready.")
	flag.StringVar(&tType, "transfer", "{{ tType }}", "Type must be 'webdav' or 'sftp'.")
	flag.Parse()

	if duration, err = strconv.Atoi(durationStr); err != nil {
		log.Fatal("Duration must be an integer!")
	}

	if sendType != "file" && sendType != "folder" && sendType != "zip" {
		err := "'type' has to be 'file', 'folder', or 'zip'"
		log.Fatal(err)
	}

	if tType != "webdav" && tType != "sftp" {
		err := "'transfer' has to be 'sftp' or 'webdav'"
		log.Fatal(err)
	}

	if dst == "" || fp == "" || sendType == "" {
		err := "'dst' and 'src' must not be empty!"
		log.Fatal(err)
	}

	u, err := url.Parse(dst)
	if err != nil && tType == "sftp" {
		u, err = url.Parse("ssh://" + dst)
	}

	if err != nil {
		log.Fatal(err)
	}

	return Args{src: fp, dst: *u, user: user, pass: pass, crt: crt, duration: time.Duration(duration) * time.Second, sendType: sendType, tType: tType, name: name}

}
