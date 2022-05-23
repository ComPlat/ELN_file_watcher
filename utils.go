package main

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"
)

// get_root_dir returns the root directory, i.e., the first directory in the path.
// If the path is relative <root>/../<file> it returns the relative root.
// If it is not relative it returns the absolut root,  i.e.: '/'
func get_root_dir(path string) string {
	current := path
	for {
		path = filepath.Dir(path)
		if path == "." || path == "" || path == current {
			return current
		}
		current = path
	}
}

// zip_folder zips a folder and safes the zipped folder with the same name in the same directory.
func zip_folder(path_src string) (string, error) {
	// Create a buffer to write our archive to.
	output_path := path_src + ".zip"
	outFile, err := os.Create(output_path)
	if err != nil {
		return "", err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			ErrorLogger.Println(err)
		}
	}(outFile)

	// Create a new zip archive.
	w := zip.NewWriter(outFile)
	defer func(w *zip.Writer) {
		err := w.Close()
		if err != nil {
			ErrorLogger.Println(err)
		}
	}(w)
	err = filepath.Walk(path_src,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return err
			}

			rel_path, err := filepath.Rel(path_src, path)
			if err != nil {
				return err
			}

			dat, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			f, err := w.Create(rel_path)
			if err != nil {
				return err
			}
			_, err = f.Write([]byte(dat))

			return err

		})

	// Make sure to check the error on Close.

	if err != nil {
		return "", err
	}
	return output_path, nil
}
