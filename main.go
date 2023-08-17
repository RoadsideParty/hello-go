package main

import "go-study/router"

func main() {
	router := router.GetRouter()
	router.Run(":8080")
}
