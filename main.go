package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
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
		log.Printf("%d %s", i, commit)
		start := time.Now()
		err := exec.Command("git", "checkout", commit).Run()
		crash(err)

		//err = exec.Command("nix", "build", ".#awsebcli").Run()
		err = exec.Command("nix-build", "-A", "awsebcli").Run()
		log.Printf("%d %s", i, time.Since(start).String())
		crash(err)
	}
}
