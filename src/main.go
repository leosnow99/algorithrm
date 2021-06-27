package main

import (
	"algorithm/src/string_type"
	"fmt"
)

func main() {
	tire := string_type.NewTire()
	tire.Insert("a")
	tire.Insert("ab")
	tire.Insert("abc")
	tire.Insert("abcd")
	fmt.Println("tire pre ab: ", tire.PrefixNumber("ab"))
}
