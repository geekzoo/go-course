package main

import (
	"fmt"
	lesson "geeky/rex_lessons/001"
	mymath "geeky/rex_lessons/002"
)

func main() {
	built := lesson.BuildIt("1", "2", "3")
	run := lesson.ReturnIT(built)
	ret := mymath.Sum(113, 224)
	fmt.Println(run, ret)
}
