package main

import (
	"fmt"
)

// main 入口方法
// 目的：实验指针较详细的相关用法
// 持续补充
func main() {
	// 变量的形式定义的函数
	// 方法内定义函数，只能用这种方式
	myFunc := func() {
		// 方法内定义结构体
		type struct1 struct {
			a int
			b string
			c []int8
		}
		// 方法内定义结构体
		type struct2 struct {
			f  int64
			g  int16
			s1 struct1
		}
		// 方法内定义的方法，必须变量的形式定义
		// 传入参数为值传递；即调用时传入参数，这里得到的只是其copy
		// 在该方法中无论做什么都不会对实际调用的值有影响
		func1 := func(ins struct2) {
			ins.f++
			ins.s1.b = "func1 changed"
			fmt.Println("in func1 : ", ins)
		}
		// 方法内定义方法，必须变量的形式
		// 指针作为参数，类型前加“*”；让传入变量的引用作为参数（使用&），而不是值的复制
		// 在该方法中对引用参数的值做写操作，会对实际调用的值有影响
		func2 := func(ins *struct2) {
			// 以下这行似乎很奇怪，别在意，理解为语法糖即可；实际等同于 (*ins).f++ ；后面类似
			ins.f++
			ins.s1.b = "func2 changed"
			fmt.Println("in func2 : ", *ins)
		}
		// 方法内定义方法，必须变量的形式
		// 指针作为参数，类型前加“*”；让传入变量的引用作为参数（使用&），而不是值的复制
		// 尝试在该方法中，修改引用；实验发现不会对实际被调用的值造成影响
		func3 := func(ins *struct2) {
			// 以下这行对实际被调用的变量产生的影响
			ins.f++
			// 以下这行不会对实际被调用的变量产生的影响，且因为这个操作，后续的所有操作都不会
			ins = &struct2{f: 666, g: 6666, s1: struct1{a: 233, b: "new-one"}}
			// 由于引用已经改变，以下这行不会对实际被调用的变量产生的影响
			ins.f++
			fmt.Println("in func3 : ", *ins)
		}
		// 方法内定义方法，必须变量的形式
		// 指针作为参数，类型前加“*”；让传入变量的引用作为参数（使用&），而不是值的复制
		// 尝试在该方法中，修改引用对应的值；实验发现会对实际被调用的值造成影响
		func4 := func(ins *struct2) {
			// 以下这行对实际被调用的变量产生的影响
			ins.f++
			// 因为引用传递的原因，以下这行会对实际被调用的变量产生的影响，因为是对引用赋值的操作，实际引用地址并没有改变
			*ins = struct2{f: 666, g: 6666, s1: struct1{a: 233, b: "new-one"}}
			// 由于引用没有改变，以下这行继续对实际被调用的变量产生的影响
			ins.f++
			fmt.Println("in func4 : ", *ins)
		}
		// 声明一个变量
		s2ins := struct2{
			f: 324,
			g: 23,
			s1: struct1{
				a: 12,
				b: "init-b",
				c: []int8{32, 3, 4, 7},
			},
		}
		// 验证我们的输出
		fmt.Println("begin    : ", s2ins)
		func1(s2ins)
		fmt.Println("out func1: ", s2ins)
		func2(&s2ins)
		fmt.Println("out func2: ", s2ins)
		func3(&s2ins)
		fmt.Println("out func3: ", s2ins)
		func4(&s2ins)
		fmt.Println("out func4: ", s2ins)
	}
	myFunc()
}
