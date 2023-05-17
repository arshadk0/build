package main

import "build/router"

func main() {
	// Start the server
	router.Router.Run(":8080")
}
