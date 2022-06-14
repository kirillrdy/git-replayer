package main

import (
	"log"
	"os/exec"
)

func crash(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	out, err := exec.Command("git", "rev-list", "master").CombinedOutput()
	crash(err)
	log.Print(out)
}
