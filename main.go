package main

import (
	"fmt"
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
	out, err := exec.Command("git", "rev-list", "--topo-order", "master").CombinedOutput()
	crash(err)
	commits := strings.Split(string(out), "\n")
	for i, commit := range commits {
		start := time.Now()
		err := exec.Command("git", "checkout", commit).Run()
		crash(err)

		//err = exec.Command("nix", "build", ".#awsebcli").Run()
		err = exec.Command("nix-build", "-A", "awsebcli").Run()
		duration := time.Since(start)
		if duration > 15*time.Second {
			log.Printf("%d %s", i, time.Since(start).String())
		} else if err != nil {
			log.Printf("%d %s", i, err)
		} else {
			fmt.Print(".")
		}
	}
}
