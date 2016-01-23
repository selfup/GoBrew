package main

import "github.com/codeskyblue/go-sh"

func main() {
    sh.Command("echo", "Starting Brew Awesomeness").Run()
    sh.Command("brew", "update").Run()
    sh.Command("brew", "upgrade").Run()
    sh.Command("brew", "cleanup").Run()
    sh.Command("brew", "doctor").Run()
    sh.Command("echo", "Brew Awesomeness Is Done").Run()
}
