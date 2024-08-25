package main

import (
	"main/pkg/globals"
	"main/server"
)


func main() {
	globals.SetupEnvironmentVariables()
	server.Run()
}