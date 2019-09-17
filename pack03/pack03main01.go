package main

import (
	"fmt"

	"../pack02"
)

// main 入口方法
// 引入了 pack02 包
// 实验基本的类型的指针使用
// 以及结构体的简单定义和声明
func main() {
	// 声明一个普通的变量
	var varInt int8 = 3
	// 声明一个指针变量，存储的是引用地址; 等同于 var ptrInt *int8 = &varInt
	var ptrInt = &varInt
	// 指针变量的赋值必须来自于一个变量的引用
	// & 符号即获取!变量!的内存地址（引用地址）
	// 所以不能使用成 ptrInt3 = &&ptrInt, 因为后一个&ptrInt得到的内存地址还没有分配到地址，前一个&当然就不能行驶其功能
	var ptrInt2 = &ptrInt
	var ptrInt3 = &ptrInt2
	// 打印验证一些值
	fmt.Println(ptrInt)
	fmt.Println(ptrInt2)
	fmt.Println(ptrInt3)
	fmt.Println(*ptrInt)
	fmt.Println(*&ptrInt)
	fmt.Println(**&ptrInt)
	fmt.Println(***&ptrInt3)
	// ---
	fmt.Println()
	// ---
	type aStruct struct {
		f1 int
		f2 string
	}
	type bStruct struct {
		f2 int
		f3 float32
		f4 string
	}
	// 方法中定义了一个结构体
	type myStruct struct {
		// 这样表示继承
		aStruct
		// 可以多继承
		bStruct
		// 这只是个成员
		as aStruct
		// 这也是个成员
		bs bStruct
		// 这也是个成员，只不过是匿名结构体的形式出现的
		inner1 struct {
			ee int
		}
		// 这里并没有覆盖bStruct的f4，只是覆盖了bStruct的f4在myStruct直接访问的权利
		f4 int64
		// 成员的类型是interface{}表示可以是所有的类型，go中所有的类型都实现了所有无方法的接口
		a interface{}
		b interface{}
		c interface{}
	}
	// 定义结构体的变量实例 方式1: 使用 structXXX{xxx}
	// 对于继承的成员，必须按照嵌套的方式（不要漏了:右端的类型）定义成员值，如下
	ins1 := myStruct{aStruct: aStruct{f1: 4343}, a: 32, b: "d4545", c: []interface{}{2, "aa", 44, 3}}
	// 实际调用继承的成员，以下方式仍然有效
	fmt.Println(ins1.f1)
	// 定义结构体的变量实例 方式2: 使用 new(typeXXX)
	// 这里的type不仅仅是定义的结构体，可以是 int、string 等基本类型
	// 更重要的是！！！ 该方式返回的是引用！！！
	// 以下等同于 &myStruct{}
	ins2 := new(myStruct)
	// 访问继承类型的成员，因为有多继承，只要不是有歧义的都能这样访问
	ins2.f1 = 7871
	// 由于f2在这里多继承是有歧义的，像如下所以这样访问（ins2.f2是找不到的）
	// 视作内部有一个和继承类型同名的成员，当然这种方式对于不歧义的继承成员是访问的同一个成员
	ins2.aStruct.f2 = "f2qdada"
	// 根据上面描述，以下等价于fmt.Println(ins2.f1)
	fmt.Println(ins2.aStruct.f1)
	ins2.f4 = 4444444434
	ins2.bStruct.f4 = "f44444"
	ins2.bs.f4 = "bs.f44444"
	ins2.a = "rww"
	ins2.b = map[interface{}]interface{}{"a3": "33f", 2: 324}
	// 所以结构体变量直接访问成员(insXXX.fieldXXX)的规则为：
	// - 首先看当前结构体是否有定义，有则访问；
	// - 无则按照树结构寻找继承的结构体中是否有，有则访问；
	// - 如果树结构中有冲突（多继承有不同结构体有同名成员），则只能按照树结构方式逐级调用
	// 另外到此可见 interface{} 却是指代任意类型（类似Java的Object）；
	// 但是除非直到为什么要使用，否则最好别用
	fmt.Println(ins1)
	fmt.Println(ins2)

	fmt.Println(pack02.P02c01func01())
	fmt.Println(pack02.P02c02func01())
	fmt.Println(pack02.NewP02c01struct02())
}
