// you can get all in one array code in arrays.go file by running this command " go run arrays.go "

package main

import "fmt"

func main() {
	//b := [5]int{1, 2, 3, 4, 5}   //in this way you can add controlled array , where you can assign number of indexes we can add data
	b := [...]int{1, 2, 3, 4, 5, 3, 5, 8, 3, 43, 5} //in this way you can uncontrolled array , can add as many as data in teh array
	fmt.Println(b)

	//check the length of array

}
