package persian_test

import (
	"fmt"

	"github.com/aslrousta/persian"
)

func ExampleWrap() {
	fmt.Printf("Int:   %v\n", persian.Wrap(1234))
	fmt.Printf("Float: %v\n", persian.Wrap(1234.5678))
	// Output:
	// Int:   ۱۲۳۴
	// Float: ۱۲۳۴٫۵۶۷۸
}

func ExampleWrapWithSep() {
	fmt.Printf("Int:   %v\n", persian.WrapWithSep(1234))
	fmt.Printf("Float: %v\n", persian.WrapWithSep(1234.5678))
	// Output:
	// Int:   ۱٬۲۳۴
	// Float: ۱٬۲۳۴٫۵۶۷۸
}
