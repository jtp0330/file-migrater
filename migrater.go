package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func migrate(dir string, new_dir string) {
	fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}
	//check if dir exists, if no, create it
	if _, err := os.Stat(new_dir); os.IsNotExist(err) {
		err = os.Mkdir(new_dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range files {
		os.Rename(dir+"/"+file.Name(), new_dir+"/"+file.Name())
	}
}
