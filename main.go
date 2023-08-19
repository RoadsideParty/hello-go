package main

import (
	"fmt"
	"go-study/router"
)

func main() {
	myRouter := router.GetRouter()
	err := myRouter.Run(":8080")
	if err != nil {
		fmt.Print(err)
	}
}
