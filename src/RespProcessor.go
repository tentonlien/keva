package main

import (
	"fmt"
	"strconv"
)

type RespParam struct {
	params []string
}


func process(raw []string)  {
	//fmt.Println("raw:", raw)
	//parseArray(raw, 0)
	//fmt.Println(restParams)
}


func parseArray(params []string, pos int) (param RespParam) {
	for i := pos; i < len(params); i ++ {
		fmt.Print(params[pos], " ")
		if params[i][0] == '*' {
			length, err := strconv.Atoi(params[i][1:])
			fmt.Println("length:", length)
			if err != nil {
				fmt.Println("Error", err.Error())
			}
			//arr := [length]string
			for k := 0; k < length; k ++ {
				fmt.Print(params[i + k], " ")
			}
		}
	}
	return
}
