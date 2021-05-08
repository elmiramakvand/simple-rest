package main

import (
	"simple-rest/api"
)

func main() {
	r := api.RunApi()
	//running
	r.Run(":8080")

}
