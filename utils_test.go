package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetRootDir(t *testing.T) {
	{
		got := getRootDir("")
		if got != "" {
			t.Errorf("getRootDir(\"\") = %s; want \"\"", got)
		}
	}
	{
		got := getRootDir("/")
		if got != "/" {
			t.Errorf("getRootDir(\"/\") = %s; want \"/\"", got)

		}
	}
	{
		got := getRootDir("/Ap")
		if got != "/" {
			t.Errorf("getRootDir(\"/\") = %s; want \"/\"", got)

		}
	}
	{
		got := getRootDir("/Ap/Bp")
		if got != "/" {
			t.Errorf("getRootDir(\"/\") = %s; want \"/\"", got)

		}
	}
	{
		got := getRootDir("//Bp")
		if got != "/" {
			t.Errorf("getRootDir(\"/\") = %s; want \"/\"", got)

		}
	}
	{
		got := getRootDir("Ap")
		if got != "Ap" {
			t.Errorf("getRootDir(\"/\") = %s; want \"Ap\"", got)

		}
	}
	{
		got := getRootDir("Ap/")
		if got != "Ap" {
			t.Errorf("getRootDir(\"/\") = %s; want \"Ap\"", got)

		}
	}
	{
		got := getRootDir("Ap/Bp")
		if got != "Ap" {
			t.Errorf("getRootDir(\"/\") = %s; want \"Ap\"", got)

		}
	}
}

func cleanTestDir() {
	if err := os.RemoveAll("testDir"); err != nil {
		fmt.Println(err)
	}
	if err := os.Mkdir("testDir", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("testDir/src", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("testDir/dst", 0777); err != nil {
		log.Fatal(err)
	}
}

func writeIntoFile(path string, content string) {
	dst, err := os.Create(path) // dir is directory where you want to save file.
	if err != nil {
		log.Fatal(err)
	}
	defer func(dst *os.File) {
		if err := dst.Close(); err != nil {
			log.Fatal(err)
		}
	}(dst)
	if _, err = dst.Write([]byte(content)); err != nil {
		log.Fatal(err)
	}
}

func TestZipFolder(t *testing.T) {
	cleanTestDir()
	defer cleanTestDir()

	// Prepare Test
	if err := os.MkdirAll("testDir/A/B", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("testDir/A/C", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("testDir/C", 0777); err != nil {
		log.Fatal(err)
	}

	writeIntoFile("testDir/A/B/a.txt", "Hallo A_B_a")
	writeIntoFile("testDir/A/b.txt", "Hallo A_c")
	writeIntoFile("testDir/A/C/c.txt", "Hallo A_C_c")
	writeIntoFile("testDir/C/d.txt", "Hallo C_d")
	writeIntoFile("testDir/e.txt", "Hallo e")
	// Done Prepare

	folderA, err := zipFolder("testDir/A")
	if err != nil {
		return
	}

	folderC, err := zipFolder("testDir/C")
	if err != nil {
		return
	}

	folderE, err := zipFolder("testDir/e.txt")
	if err != nil {
		return
	}

	if _, err := os.Stat(folderA); err != nil {
		t.Errorf("zipFolder(\"testDir/A\") did not work!")

	}

	if _, err := os.Stat(folderC); err != nil {
		t.Errorf("zipFolder(\"testDir/C\") did not work!")

	}

	if _, err := os.Stat(folderE); err != nil {
		t.Errorf("zipFolder(\"testDir/e.txt\") did not work!")

	}
}
