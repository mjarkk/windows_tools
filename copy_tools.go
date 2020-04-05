package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	binFolder := ".\\bin"
	startMenu := path.Join(usr.HomeDir, "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\mjarkk_windows_tools")

	os.RemoveAll(startMenu)
	err = os.Mkdir(startMenu, 0777)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(binFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		from := path.Join(binFolder, file.Name())
		bytes, err := ioutil.ReadFile(from)
		if err != nil {
			log.Fatal(err)
		}

		to := path.Join(startMenu, file.Name())
		ioutil.WriteFile(to, bytes, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}
