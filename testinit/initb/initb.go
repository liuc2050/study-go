package initb

import (
	"fmt"
)

var B = c[10]
var c [256]byte = func() (b [256]byte) {
	for i := range b {
		b[i] = b[i/2] + byte(i&1)
	}
	return
}()

func init() {
	fmt.Println(B)
}
