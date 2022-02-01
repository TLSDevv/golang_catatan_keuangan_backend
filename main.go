package main

import (
	"fmt"
	"strings"
)

func main() {
	// cmd.Execute()
	test := "test"
	test += " yey"

	var sb strings.Builder

	sb.WriteString("test yey")
	sb.WriteString("\ntest again")
	fmt.Println(test)
	fmt.Println(sb.String())

	params := map[string]interface{}{}
	params["test"] = 59
	params["test2"] = 69
	params["test3"] = "wei"
	params["test4"] = "wei"
	params["test5"] = "wei"

	value, isExist := params["test6"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("not exist")
	}
}
