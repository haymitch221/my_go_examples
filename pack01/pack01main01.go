package main

import (
	"flag"
	"fmt"
)

// 这里定义了一个全局变量，string，有赋值
// 是因为 import 了 flag 来读取 终端 go run xxx.go 后的输入
// 该变量保存输入值，由 flag 的方法来操作
var what = "hello(declare in code)"

// 这里定义了一个全局变量，string，未赋值
// 用途和上面的 what 是一样的
var who string

// P01m01func01 这是个方法，公共的，有一个参数，有一个返回值
// return string 返回值是 string 类型
func P01m01func01(what, who string) string {
	return "sure, " + what + " " + who + "! ... ok, done"
}

// package main 的 init 方法
// 会在main之前被调用
// 定义程序的初始化逻辑，实际就是个方法被调而已
// 可能有些事情会发生，不确定 (TODO：can it be certain?）
func init() {
	// flag 的逻辑，终端 go run xxx.go 后的输入
	// 需要绑定用于保存来源参数的变量
	// 需要指定来源参数名，参数默认值，以及描述说明(暂不知道描述说明有什么用，可能没用吧，不清楚)
	// 档案这里的赋值（包括默认）都会覆盖掉变量本身的初始赋值（shut up, you know why）
	flag.StringVar(&what, "what", "hello(default)", "what will do")
	flag.StringVar(&who, "who", "Kid(default)", "what will do to someone")
}

// package main 的 main 方法
// 会在 init 之后被调用
// 定义程序入口，正式开始
func main() {
	// 这个方法的目的是停止接受来源参数
	// 然后转化保存到绑定的变量中
	flag.Parse()
	fmt.Println(P01m01func01(what, who))
}
