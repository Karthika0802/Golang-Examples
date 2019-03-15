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

    //modification happens here

    var mod [5][5]int


    for q := 0; q < 5; q++{ 
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
     fmt.Println("Modified Version:")

     for g := 0; g < 5; g++{ 
        fmt.Println(" ") 
        for h := 0; h < 5; h++{ 
             fmt.Print(mod[g][h])
             fmt.Print(" ")
         } // end for h:=
    } // end for g:=

} // end main
        
