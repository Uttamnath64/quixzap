package main

import "github.com/Uttamnath64/quick-connect/api/application"

func main() {
	// application start
	apps := application.New()
	if apps.Initialize() {
		apps.Run()
	}
}
