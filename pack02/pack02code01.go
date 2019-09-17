// Package pack02 定义包名
// 这里包名和文件名不一定要相同，可以是任意的有效字符串
// 当然一旦确定了包名，同级的所有go代码文件都得是同一个包名
package pack02

// P02c01struct01 is a struct
// 结构体，就是一个多类型的映射
// 由于interface{}
type P02c01struct01 struct {
	field01 int
	field02 int8
	field03 string
	// 简单的 map 类型，k v 都是基本类型
	field04 map[string]string
	// 该 map 中 v 类型为 interface{}，表示可以进入任意类型
	field05 map[string]interface{}
	// 该 map 中 v 类型为 struct{}
	// 匿名结构，只要符合这样的结构都能进入
	// 符合的方式为：结构内的成员声明以及类型均匹配
	field06 map[string]struct {
		b int
		c int
	}
}

// P02c01struct02 is a struct as well
type P02c01struct02 struct {
	field01 int16
	field02 []string
	field03 P02c01struct01
}

// P02c01func01 is a func
func P02c01func01() string {
	return "this is func P02c01func01"
}

// P02c01func02 是个私有方法
// 只能被该包内代码调用
// 即import到别的包的go代码中时，对它们不可见
func p02c01func02() string {
	return "this is func p02c01func02"
}
