package main

import (
	"mini-project/config"
	"mini-project/routers"
)

func init() {
	config.InitDB()

}

func main() {
	e := routers.Router()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8888"))
}
