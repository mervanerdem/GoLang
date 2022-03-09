package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 	myArr := [...]int32{1, 2, 3, 4}
	   	fmt.Println(myArr)

	   	var mySlc []int32
	   	fmt.Println(mySlc)
	   	mySlc = make([]int32, 2)
	   	mySlc[1] = 2
	   	fmt.Println(mySlc)
	   	mySlc = make([]int32, 3)
	   	fmt.Println(mySlc) */

	slc1 := []string{"1, 2", "1, 2, 3"}
	fmt.Println(slc1)
	/*
		fmt.Println(string(slc1[0][0]))
		fmt.Println(string(slc1[1][4]))
	*/

	a := strings.Split(slc1[0], ", ") // array içindeki belli kısmı silme
	b := strings.Split(slc1[1], ", ") //, ve boşluk silme
	fmt.Println(a)
	fmt.Println(b)
	slc1 = append(slc1, "4") //sonuna yeni bir eleman ekleme
	fmt.Println(slc1[2])

}
