package main

import (
	"github.com/MohabMohamed/mohab.dev/src/routes"
	"github.com/MohabMohamed/mohab.dev/src/util"
)

func main() {
	util.InitLogger()
	defer util.TerminateLogger()

	routes.Init()
	routes.Serve()
}
