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
	if err := os.RemoveAll("test_dir"); err != nil {
		fmt.Println(err)
	}
	if err := os.Mkdir("test_dir", 0777); err != nil {
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
	if err := os.MkdirAll("test_dir/A/B", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("test_dir/A/C", 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("test_dir/C", 0777); err != nil {
		log.Fatal(err)
	}

	writeIntoFile("test_dir/A/B/a.txt", "Hallo A_B_a")
	writeIntoFile("test_dir/A/b.txt", "Hallo A_c")
	writeIntoFile("test_dir/A/C/c.txt", "Hallo A_C_c")
	writeIntoFile("test_dir/C/d.txt", "Hallo C_d")
	writeIntoFile("test_dir/e.txt", "Hallo e")
	// Done Prepare

	folderA, err := zipFolder("test_dir/A")
	if err != nil {
		return
	}

	folderC, err := zipFolder("test_dir/C")
	if err != nil {
		return
	}

	folderE, err := zipFolder("test_dir/e.txt")
	if err != nil {
		return
	}

	if _, err := os.Stat(folderA); err != nil {
		t.Errorf("zipFolder(\"test_dir/A\") did not work!")

	}

	if _, err := os.Stat(folderC); err != nil {
		t.Errorf("zipFolder(\"test_dir/C\") did not work!")

	}

	if _, err := os.Stat(folderE); err != nil {
		t.Errorf("zipFolder(\"test_dir/e.txt\") did not work!")

	}
}
