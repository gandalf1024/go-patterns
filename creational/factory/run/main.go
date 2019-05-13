package main

import "go-patterns/creational/factory/data"

//工厂模式  根据参数返回不同的对象
func main() {
	s := data.NewStore(data.MemoryStorage)
	f, _ := s.Open("file")

	f.Write([]byte("data"))
	defer f.Close()
}
