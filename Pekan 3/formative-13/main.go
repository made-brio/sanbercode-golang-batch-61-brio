package main

import "formative13/routers"

func main() {
	var port = ":8085"

	routers.StartServer().Run(port)
}
