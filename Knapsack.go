//The purpose of this program is to, given various objects with a COST and a WEIGHT, maximize the amount of COST a bag can hold. The bag is constrained by weight. 

package main
import "fmt"
func main() {
	cap := 100 //Total weight can still hold
	carried := 0 //Total cost bag is currently holding
	weight := [5]int{5, 10, 15, 17, 35} //Weights of the various objects
	cost := [5]int{25, 60, 100, 120, 150} //Values of various objects

	 for cap >= 5 {
 		 OptRat := 0 //the optimal ratio
 		 for i := 0; i<5; i++{ 
			 if  cap >= weight[i]{ //precondition to see the bag can even hold 
				if cost[i]/weight[i] >= OptRat  { //Checks if the variable being checked is more optimal 
					OptRat = i

				}
			 }
			}
		 
		 cap = cap-weight[OptRat]
		 carried = carried+cost[OptRat]

		}


	fmt.Print("Total value carried: " , carried)

}

