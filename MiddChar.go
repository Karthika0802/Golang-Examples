// Sorts a list of strings by the middle character, case sensetive. 

package main

import (
	"fmt"
//	"strings"
)

func main() {
	arr := [6]string{"CabiACB", "cobITTF", "Rapter", "Jolly", "Fiver", "CabiACB"} //The strings




	for i := 0; i<6 ; i++{  // Uses selection sort to find the lowest value
		lowest := i //Lowest value stored here
		for j := i+1; j<6; j++{

			if arr[i][len(arr[i])/2+1] > arr[j][len(arr[j])/2+1]{ //If compared value is lower then set current lowest to compared
				lowest = j

			}
		}
		temp := arr[i] //Swap
		arr[i] = arr[lowest]
		arr[lowest] = temp
		
	}

	for b:=0; b<6; b++{
		fmt.Println(arr[b])
	}


}