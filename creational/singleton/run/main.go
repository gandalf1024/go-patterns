package main

import (
	"fmt"
	"go-patterns/creational/singleton/singl"
)

//单例模式
func main() {
	s := singl.New()
	s["this"] = "that"
	s2 := singl.New()
	fmt.Println("This is ", s2["this"])
}
