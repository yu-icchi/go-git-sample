package main

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func main() {
	println("git clone")

	git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL: "https://github.com/yu-ichiko/go-psd.git",
		Progress: os.Stdout,
	})
}
