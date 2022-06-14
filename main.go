package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func crash(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	os.Chdir("../nixpkgs")
	out, err := exec.Command("git", "rev-list", "master").CombinedOutput()
	crash(err)
	commits := strings.Split(string(out), "\n")
	for i, commit := range commits {
		log.Printf("%d", i)
		err := exec.Command("git", "checkout", commit).Run()
		crash(err)

		err = exec.Command("nix", "build", ".#awsebcli").Run()
		crash(err)
	}
}
