package ordinal_test

import (
	"fmt"

	"github.com/aprice/gordinals/ordinal"
)

func ExampleFor() {
	fmt.Println(ordinal.For(4))
	// Output: th
}
