package main

import "fmt"

func main() {

	/*
		//1- 6 sting karekterden oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.
		myArr := [6]string{"Ben", "Sen", "O", "Biz", "Siz", "Onlar"} //Array kısmı
		for index, value := range myArr {
			fmt.Println(index, value)
		}

		fmt.Println("------------------------")

		mySlc := []int{100, 200, 300, 400, 500}
		for i, v := range mySlc {
			fmt.Println(i, v)
		} */

	/* 	//2-[0,1,2,3,4,5,6,7,8] slice'ından [0,1,2,3] ,[4,5,6] , [6,7,8], [2,3,6,7] slice'larını oluşturunuz.

	   	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	   	mySlc2 := mySlc[:4]
	   	mySlc3 := mySlc[4:7]
	   	mySlc4 := mySlc[6:]
	   	mySlc5 := append(mySlc[2:4], mySlc[6:8]...)

	   	fmt.Println("My Slc", mySlc)
	   	fmt.Println("My Slc2", mySlc2)
	   	fmt.Println("My Slc3", mySlc3)
	   	fmt.Println("My Slc4", mySlc4)
	   	fmt.Println("My Slc5", mySlc5)
	*/
	/* 	//3-slicelar için copy metodu ve assign(=) farkını açıklayın.
	   	mySlc := []int{0, 1, 2}
	   	mySlc2 := make([]int, 2)
	   	mySlc3 := make([]int, 3)

	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2)
	   	fmt.Println(mySlc3)

	   	fmt.Println("-----------")

	   	copy(mySlc2, mySlc) //copy
	   	mySlc3 = mySlc      //append (=)

	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2)
	   	fmt.Println(mySlc3)

	   	fmt.Println("-----------")

	   	mySlc2[1] = 3 //kopyalanan
	   	mySlc3[0] = 7 // atanan

	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2)
	   	fmt.Println(mySlc3)
	*/

	//4-map gösterimi ile yazar ve yazar kitabının isimlerini for range ile gösterin.

	myMap := map[string][]string{
		"Yaşar Kemal":    []string{"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali": []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
	}
	/* fmt.Println(myMap)
	fmt.Println(myMap["Mehmet"])
	fmt.Println(myMap["Mehmet"][0]) */

	for key, value := range myMap {
		fmt.Println("Yazar:", key)
		fmt.Println("Kitapları:")
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}

	}

}
