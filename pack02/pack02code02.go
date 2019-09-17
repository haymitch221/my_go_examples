package pack02

// P02c02func01 is a func
func P02c02func01() string {
	return "this is func P02c02func01" + "; " + p02c01func02()
}

// NewP02c01struct02 是个函数
// 用变量的形式定义的函数
// 因为是大写开头的原因，仍然是共有的
var NewP02c01struct02 = func() P02c01struct02 {
	// 在方法中可以声明结构体
	// 声明的作用域只在该方法中
	type innerStruct struct {
		field01 int64
	}
	// 在方法中可以声明函数，类似javascript的做法
	innerFunc4createInnerStruct := func() innerStruct {
		return innerStruct{field01: 323}
	}
	innerFunc4createInnerStructPtr := func() *innerStruct {
		return &innerStruct{field01: 4564}
	}
	is := innerFunc4createInnerStruct()
	is2 := innerFunc4createInnerStructPtr()
	is3 := innerStruct{}
	// 初始化结构体实例时，随意即可，不一定全都赋值
	ins := P02c01struct02{
		field01: 123,
		field02: []string{"11", "remove me", "333"},
		field03: P02c01struct01{
			field01: 5,
			field02: 6,
			field03: "hahh",
			field04: map[string]string{"aa": "99aaa", "bb": "fevbbbb", "cc": "ccccccxxc"},
			field05: map[string]interface{}{"aa": is, "bb": *is2, "cc": is3, "dd": struct{ a int }{a: 3}, "ee": map[int]int{3: 5, 6: 222}},
			// 这里很有意思，匿名结构体；详细略；即使忘了再想想 go 某方面的思想
			field06: map[string]struct {
				b int
				c int
			}{"aa": struct {
				b int
				c int
			}{b: 3}, "bb": {b: 23232}, "cc": {b: 545}},
		},
	}
	ins.field02 = append(ins.field02, "44544444")
	ins.field02 = append(ins.field02[:1], ins.field02[2:]...)
	return ins
}
