package sdk

import "fmt"

func main() {
	// 1. A nullable string with a value
	var myNullableString *string
	val := "hello"
	myNullableString = &val

	// To cast, check if the pointer is not nil, then dereference
	var myString string
	if myNullableString != nil {
		myString = *myNullableString
	} else {
		// Handle the nil case, e.g., by providing a default value
		myString = ""
	}

	fmt.Println("The converted string is:", myString)
}
