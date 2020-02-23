package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func getDrives() (r []string) {
	fmt.Println("input : ")
	var get string
	fmt.Scan(&get)
	f, err := os.Open(string(get) + ":\\")
	if err == nil {
		d := string(get) + ":/"
		r = append(r, d)
		f.Close()

	}
	return

}
func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {

		}

	}

}
