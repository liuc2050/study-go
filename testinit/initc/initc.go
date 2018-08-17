package initc

import (
	"fmt"
)

var C = "c"

func init() {
	fmt.Print(C)
}

func init() {
	fmt.Printf("init2%s\n", C)
}
