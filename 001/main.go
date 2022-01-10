//Package lesson 001 provide functions examples
package lesson

import "fmt"

//ReturnIT returns a string passed from another function ReturnIT("My String")
func ReturnIT(a string) string {
	return "from " + a
}

//BuildIt builds the things
func BuildIt(a, b, c string) string {
	done := fmt.Sprintf("%v %v %v", a, b, c)
	return done
}
