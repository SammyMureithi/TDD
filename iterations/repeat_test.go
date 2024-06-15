package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	repeated := Repeat("a",5)
	expected := "aaaaa"
	
	if repeated != expected {
		t.Errorf("Required %q but got %q", expected, repeated)
	}
	
}

func ExampleRepeat(){
	res:=Repeat("l",10)
	fmt.Println(res)

	// Output: llllllllll
}