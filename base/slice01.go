package main

import "fmt"

type Tes struct {
	Age string
}

func main()  {
	outSlice := []string{"1","2"}
	//modifySlice(&outSlice)
	modifySlice2(outSlice)
	fmt.Println(outSlice)


}

func modifySlice(inSlice *[]string)  {
	*inSlice = append(*inSlice,"c")
	(*inSlice)[0] = "a"
	(*inSlice)[1] = "b"

	fmt.Println(*inSlice)
}

func modifySlice2(inSlice []string)  {
	// 先扩容，指向新地址
	// inSlice = append(inSlice, "g")
	inSlice[0] = "d"
	inSlice[1] = "f"
	// 后扩容
	inSlice = append(inSlice, "g")
	fmt.Println(inSlice)
}
