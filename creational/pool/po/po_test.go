package po

import (
	"fmt"
	"testing"
)

func Test_po(t *testing.T) {
	pool := New(2)
	cc := <-*pool
	fmt.Println(cc)
}
