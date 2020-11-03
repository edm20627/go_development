package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hello
}

func ExampleUnordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered output:
	// 1
	// 3
	// 2
}

func ExampleShuffle() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range x {
		fmt.Printf("k=%s v=%d\n", k, v)
	}
	// Unordered output:
	// k=a v=1
	// k=b v=2
	// k=c v=3
}
