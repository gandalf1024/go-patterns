package main

import "go-patterns/structural/decorator/build"

func Double(n int) int {
	return n * 2
}

//装饰器模式
//允许动态扩展现有对象的功能而不改变其内部。
func main() {
	f := build.LogDecorate(Double)
	f(5)
}
