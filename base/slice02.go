package main

import "fmt"

func main()  {
	outerSlice:= make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice[0])
	modifySlice03(outerSlice)
	fmt.Println(outerSlice)
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice[0])
	fmt.Println(len(outerSlice),cap(outerSlice))
	fmt.Println(outerSlice)
}

func modifySlice03(innerSlice []string)  {
	innerSlice = append(innerSlice, "c")
	fmt.Printf("%p %v   %p\n", &innerSlice, innerSlice, &innerSlice[0])
	innerSlice[0] = "1"
	innerSlice[1] = "2"

	fmt.Println(innerSlice)

}
