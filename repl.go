package main

import (
	"strings"
)

func cleanInput(text string) []string {
	stringList := strings.Fields(text)
	for i, text := range stringList {
		stringList[i] = strings.ToLower(text)
	}
	return stringList
}
