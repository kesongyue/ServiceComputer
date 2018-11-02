package main

import (
	"Go-Agenda/cmd"
	"Go-Agenda/entity"
)

func main() {
	entity.ReadJson()
	entity.ReadLoginJson()
	cmd.Execute()
}
