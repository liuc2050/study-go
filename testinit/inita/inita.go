package inita

import (
	"fmt"

	_ "github.com/liuc2050/studygo/testinit/initc"
)

var A = "a"

func init() {
	fmt.Println(A)
}
