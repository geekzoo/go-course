package main

import (
	"fmt"

	lesson "github.com/geekzoo/go-course/001"
	mymath "github.com/geekzoo/go-course/002"
)

func main() {
	built := lesson.BuildIt("1", "2", "3")
	run := lesson.ReturnIT(built)
	ret := mymath.Sum(113, 224)
	fmt.Println(run, ret)
}
