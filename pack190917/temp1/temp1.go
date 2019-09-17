package temp1

// type 的进阶用法窥见 go 的函数式编程特性
// 函数式编程的多态
// 由于Go是强类型，有type概念
// 函数也是一个复合类型，由调用者参数、传入参数、返回类型以及实现逻辑决定
// 而在Go中，前三个特征一致（模式匹配的规则，人能意识到的，同时go编译器也能识别的）的一组函数视作同一个声明（给这个声明命名为“抽象的函数类型声明”，或“函数接口”）
// 对比在Javascript中，没有抽象的函数类型声明
// 因为Javascript是弱类型，函数的类型可以随易指定，但需要人为做类型检查

// 将函数式编程，看作程序迈向自我意识的过程，而不是面向对象式
// ...

type StructA struct{
	Name string
}

// FuncTypeA 定义的函数接口
type FuncTypeA func(p *StructA) string

// FuncTypeB 定义的函数接口
type FuncTypeB func(p *StructA, f FuncTypeA) string

// FuncTypeB01 实现的函数接口
func FuncTypeB01(p *StructA, f FuncTypeA) string {
	return "<B-01>" + f(p) + "</B-01>"
}

// FuncTypeA01 实现的函数接口
func FuncTypeA01(p *StructA) string {
	return "<A-01>" + p.Name + "</A-01>"
}

// FuncTypeA02 实现的函数接口
func FuncTypeA02(p *StructA) string {
	return "<A-02>" + p.Name + "</A-02>"
}

// FuncTypeA03 实现的函数接口
func FuncTypeA03(p *StructA) string {
	return "<A-03>" + p.Name + "</A-03>"
}
