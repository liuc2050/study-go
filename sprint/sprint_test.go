package sprint

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSprint(t *testing.T) {
	var test FileInt = 2
	if ret := Sprint(test); ret != "FileInt 2" {
		t.Errorf("need FileInt 2\n")
	}

	v := reflect.ValueOf(test)
	fmt.Printf("kind == Int: %t\n", v.Kind() == reflect.Int)
}
