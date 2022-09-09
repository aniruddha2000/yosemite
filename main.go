package main

import (
	"flag"
	"log"

	"github.com/aniruddha2000/yosemite/app/client"
	"github.com/aniruddha2000/yosemite/app/service"
)

func main() {
	var nameSpace string

	flag.StringVar(&nameSpace, "ns", "test-ns",
		"namespace name on which the checking is going to take place")

	log.Printf("Checking Pods for namespace %s\n", nameSpace)
	c := client.NewClient()
	c.C = service.Init()

	c.CheckPodEnv(nameSpace)
}
