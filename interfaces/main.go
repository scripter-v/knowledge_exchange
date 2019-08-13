package main

import "fmt"

type myInterface interface {
	printNums()
}

func printSomething(printer myInterface) {
	printer.printNums()
}

type MyType struct {
	array    []int
	lenArray int
	Prefix   string
}

func (t *MyType) AddNum(num int) {
	t.array = append(t.array, num)
	t.lenArray = len(t.array)
}

func (t *MyType) PrintLen() {
	fmt.Println("len of array is", t.lenArray)
}

func (t *MyType) printNums() {
	fmt.Println(t.array)
}

type myType2 struct {
	sss string
}

func (s *myType2) printNums() {
	fmt.Println(s.sss)
}

type myType int

func (i myType) addNum(num int) int {
	return int(i) + num
}

func main() {
	x := &MyType{}
	y := &MyType{}
	z := &myType2{sss: "abcd"}

	x.AddNum(1)
	x.AddNum(2)
	x.lenArray = -10
	x.PrintLen()

	y.AddNum(0)

	printSomething(x)
	printSomething(z)

	var i myType = 123
	fmt.Println(i.addNum(111))
}
