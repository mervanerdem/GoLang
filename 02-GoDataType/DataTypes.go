package main

import "fmt"

func main() {
	var myNum int = 123       //int number
	var myWord string = "Hi!" //string
	var myBool bool = true    //Just true or false

	fmt.Println(myNum, myWord, myBool)

	//For see Which data type used
	fmt.Printf("%T ", myNum)
	fmt.Printf("%T ", myWord)
	fmt.Printf("%T ", myBool)

}

/* Output
123 Hi! true
int string bool
*/
