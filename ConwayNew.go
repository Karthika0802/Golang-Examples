/*This simulation is called Conways' game of life. The simulation models population change according to certain rules.
Each cell is either inhabited or uninhabited. Each inhabited/uninhabited cell can change its status in each iteration 
The rules are are follows:
1. If a inhabited cell has less than 2 neighbors, it will die from lonliness
2. If a inhabited cell has 2-3 neighbors, it will live
3. If a inhabited cell has over 3 neighbors, it will die from overpopulation
4. If a uninhabited cell has exactly 3 neigbors, it will become inhabited by reproduction
The simulations' original input is determined via randomly generated numbers.
A 1 represents an inhabited cell and a 0 represents an uninhabited one.

The program runs 10 iterations following the original.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UnixNano()) // randomize seed


//the original table



    var arr = [5][5] int{  
        {rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)} ,   
        {rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)} ,     
        {rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)} ,   
        {rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)} ,     
        {rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)}}

 //original chart printed
 
    fmt.Println("Original:") 


 for i := 0; i < 5; i++{ 
    fmt.Println(" ") 
    for j := 0; j < 5; j++{ 
         fmt.Print(arr[i][j])
         fmt.Print(" ")
     }
    }

    var mod [5][5]int //defining a second array to hold modified values

    //modification happens here


     for iter := 1; iter < 11; iter++{ //To create multiple iterations of this process
        for q := 0; q < 5; q++{  //Nested loop to find the values of the next iteration
            for z := 0; z < 5; z++{ 
                
                neigh := arr[q][z] * -1  //Running count of living neighbors. Multiplied by -1 to counter itself from being countered in the loop
                
                for a := -1; a <=1; a++{ //To check values of neighbors
                    for b := -1; b <=1; b++{
                    
                        if q+a > - 1 && z+b > -1 && q+a < 5 && z+b < 5{ //To ensure array doesn't go out of bounds
                        neigh += arr[q+a][z+b] //Adds to running total of living neighbor

                        }
                        
                    }
                }
                

                if arr[q][z] == 1 && neigh < 2 {                   //death from lonliness
                        
                    mod[q][z] = 0

                } else if arr[q][z] == 1 && neigh > 3 {             //death from overpopulation                 
            
                    mod[q][z] = 0

                } else if arr[q][z] == 0 && neigh == 3{             //reproduction at exactly 3 neighboring cells

                    mod[q][z] = 1

                } else {
                    mod[q][z] = arr[q][z]
                }
                    
                        
            }
        }
        fmt.Println("")
        fmt.Println("Modified Iteration: " , iter)

        for g := 0; g < 5; g++{ 
            fmt.Println(" ") 
            for h := 0; h < 5; h++{ 
                fmt.Print(mod[g][h])
                fmt.Print(" ")
            } // end for h:=
        }

        for k := 0; k < 5; k++{ 
            for l := 0; l < 5; l++{ 
                arr[k][l] = mod[k][l]
            } 
        }
    }




} // end main
        
