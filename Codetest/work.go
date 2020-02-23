package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					*files = append(*files, dir+"/", f.Name())
				}
			}

		}

	}

}
