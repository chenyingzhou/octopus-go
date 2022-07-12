package main

import (
	"github.com/chenyingzhou/octopus-go/driver"
	"os"
	"os/signal"
)

func main() {
	driver.CanalStart()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
