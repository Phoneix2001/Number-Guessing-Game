package main

import (
	"fmt"
	"number_guessing_game/api"
)

func main()  {
	server := api.NewServer()
	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
}
