package main

import (
	"flag"
	"log"
	"net/url"
	"time"
)

type Args struct {
	src, user, pass string
	dst             url.URL
	duration        time.Duration
	zipped          bool
}

// GetCmdArgs Get/Parse command line arguments manager
func GetCmdArgs() Args {
	var fp, dst, user, pass string
	var duration int
	var zipped bool

	flag.StringVar(&fp, "src", "", "Source directory to be watched.")
	flag.StringVar(&dst, "dst", "", "WebDAV destination URL. If the destination is on the lsdf, the URL should be as follows:\nhttps://os-webdav.lsdf.kit.edu/<OE>/<inst>/projects/<PROJECTNAME>/\n            <OE>-Organisationseinheit, z.B. kit.\n            <inst>-Institut-Name, z.B. ioc, scc, ikp, imk-asf etc.\n            <USERNAME>-User-Name z.B. xy1234, bs_abcd etc.\n            <PROJRCTNAME>-Projekt-Name")
	flag.StringVar(&user, "user", "", "WebDAV user")
	flag.StringVar(&pass, "pass", "", "WebDAV Password")
	flag.IntVar(&duration, "duration", 300, "Duration in seconds, i.e., how long a file must not be changed before sent.")
	/// Only considered if result are stored in a folder.
	/// If zipped is set the result folder will be transferred as zip file
	flag.BoolVar(&zipped, "zip", false, "Only considered if result are stored in a folder. If zipped is set the result folder will be transferred as zip file.")
	flag.Parse()

	if dst == "" || fp == "" {
		err := "'dst' and 'src' must not be empty!"
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	u, err := url.Parse(dst)
	if err != nil {
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	return Args{src: fp, dst: *u, user: user, pass: pass, duration: time.Duration(duration) * time.Second, zipped: zipped}

}
