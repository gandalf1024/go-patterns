package main

import (
	"fmt"
	"go-patterns/creational/pool/po"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	p := po.New(2)
	select {
	case obj := <-*p:
		fmt.Println(obj)
		str := "00000"
		var i interface{} = str
		*p <- &i
		wg.Done()
	default:
		wg.Done()
		return
	}
	wg.Wait()
}
