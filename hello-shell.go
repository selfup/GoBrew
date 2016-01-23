package main

import "github.com/codeskyblue/go-sh"

func main() {
    sh.Command("echo", "hello world").Command("cut", "-f2").Run()
}
