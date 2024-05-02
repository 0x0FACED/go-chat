package main

import (
	"go-chat/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
