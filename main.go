// main package as loader for the application
package main

import (
	"github.com/okyws/dashboard-backend/config"
	"github.com/okyws/dashboard-backend/routes"
)

func main() {
	config.InitCommand()
	config.InitiateLog()
	routes.RunServer()

	defer config.CloseLog()
}
