package main

import (
	"fmt"

	_ "github.com/liuc2050/studygo/testinit/inita"
	_ "github.com/liuc2050/studygo/testinit/initb"
	_ "github.com/liuc2050/studygo/testinit/initc"
)

func main() {
	var s string
	fmt.Println(s[0] == 'x')
	fmt.Println("main")
}
