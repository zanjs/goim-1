package main

import (
	"goim/logic/controller"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	controller.Engine.Run(":8080")
}
