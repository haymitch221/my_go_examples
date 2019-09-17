package main

import (
	"fmt"
)

// Level01StructBase 基本结构体
type Level01StructBase struct {
	name  string
	age   int8
	color string
	speed int64
}

// Level02Struct01 猫结构体
// 继承了 Level01StructBase
type Level02Struct01 struct {
	Level01StructBase
	content string
	brand   string
	shape   string
}

// Level02Struct02 猫结构体
// 继承了 Level01StructBase
type Level02Struct02 struct {
	Level01StructBase
	alias string
}

// Level02Struct03 人结构体
// 继承了 Level01StructBase
type Level02Struct03 struct {
	Level01StructBase
	city   string
	volume int16
}

// ItrfA 某个接口
type ItrfA interface {
	Init(sample *Level02Struct01)
	Func01() string
}

// ItrfB 某个接口
type ItrfB interface {
	Init(sample *Level02Struct02)
	Func02(p string) string
}

// ItrfC 某个接口
type ItrfC interface {
	Func03(itrfA ItrfA, itrfB ItrfB) string
}

// Init 方法 for Level02Struct01
func (st *Level02Struct01) Init(sample *Level02Struct01) {
	// 由于传进来的都是引用地址，直接用“*”获取值，然后传入；下同
	// 目的保证引用地址不变，即值可以随便改
	*st = *sample
}

// Init 方法 for Level02Struct02
func (st *Level02Struct02) Init(sample *Level02Struct02) {
	*st = *sample
}

// Func01 方法
// return name
// Level02Struct01 实现 ItrfA
// 如果这里使用 Level01StructBase 并不会让 其子结构体都变成实现 ItrfA
func (st Level02Struct01) Func01() string {
	n := st.content
	st.age++
	st.content += "!"
	return n
}

// Func02 方法
// return name
// Level02Struct02 实现 ItrfB
func (st *Level02Struct02) Func02(p string) string {
	n := st.alias
	st.age++
	return "<" + p + ">" + n + "</" + p + ">"
}

// Func03 方法
// return name
// 根据该方法的调用demo来看，接口作为参数必须为引用！这似乎意味着什么，即使方法声明上并没有 “*” 号
// 也许 interface 和引用有声明关系？
func (st *Level02Struct03) Func03(itrfA ItrfA, itrfB ItrfB) string {
	n := st.city
	st.age++
	return "Func03mix: " + itrfA.Func01() + " & " + itrfB.Func02(n)
}

func main() {
	st03 := Level02Struct03{city: "London"}
	// 接口声明时，用 new() 方法没法给结构体变量赋初值
	// 用 & 可以给结构体变量赋初值
	var st01 ItrfA = &Level02Struct01{content: "yesss"}
	var st02 ItrfB = new(Level02Struct02)
	st02.Init(&Level02Struct02{alias: "John"})
	fmt.Println(st03.Func03(st01, st02))

	// 这个 demo 更适合大多数情况，实际变量声明并不需要关联接口；
	// 变量的引用与接口有关系；只要接口题实现了接口
	stA := Level02Struct01{content: "Okkkk"}
	stB := Level02Struct02{alias: "Alex"}
	fmt.Println(st03.Func03(&stA, &stB))
}
