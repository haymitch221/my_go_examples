package main

import (
	"./temp1"
	"fmt"
)

func main() {
	stA := &temp1.StructA{}
	stA.Name = "Alex"
	fmt.Println(temp1.FuncTypeB01(stA, temp1.FuncTypeA01))
	fmt.Println(temp1.FuncTypeB01(stA, temp1.FuncTypeA02))
	fmt.Println(temp1.FuncTypeB01(stA, temp1.FuncTypeA03))
	fmt.Println(temp1.FuncTypeB01(stA, func(p *temp1.StructA) string{
		return "<A-lambda>" + p.Name + "</A-lambda>"
	}))
}