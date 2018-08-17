package main

import "fmt"

func main() {
	fmt.Println("main")
	Double(4)
}

func Double(x int) (ret int) {
	defer PrintRet(&x, &ret)
	defer func() func() {
		fmt.Printf("Begin (%d) = %d\n", x, ret)
		return func() {
			fmt.Printf("End (%d) = %d\n", x, ret)
		}
	}()()
	defer GetPrintFunc("Double", &x, &ret)()

	return x + x
}

func GetPrintFunc(funcName string, x, ret *int) func() {
	fmt.Printf("Begin %s(%d) = %d\n", funcName, *x, *ret)
	return func() {
		fmt.Printf("End %s(%d) = %d\n", funcName, *x, *ret)
	}
}

func PrintRet(x, ret *int) {
	fmt.Printf("PrintRet Double(%d) = %d\n", *x, *ret)
}
