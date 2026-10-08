package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cnf := config{}
	cnf.initConfig()
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			inputList := cleanInput(input)
			command, ok := cnf.commandRegistry[inputList[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			args := inputList[1:]
			err := command.callback(&cnf, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
