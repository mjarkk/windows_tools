package helpers

import (
	"log"
	"os/exec"
	"os/user"
	"path"
)

func OpenUserDir(dir string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	OpenDir(path.Join(usr.HomeDir, dir))
}

func OpenDir(dir string) {
	exec.Command("explorer", dir).CombinedOutput()
}
