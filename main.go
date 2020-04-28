// Programmer name: Vincent G.
// Date completed:  4/28/2020
// Description: 4.2.1 Lab - Struct Part 1

package main

import "fmt"

//Create a struct that keeps track of animal and stores its name, age, breakfast hour, and dinner hour.

type animal struct {
  name, age, breakfast, dinner string 
}

func main() {
   //create 2 animal objects that store the appropriate data and then print out the data stored

   Gorilla := animal{"Benny", "9 yrs old", "9am", "5pm"}
   fmt.Println(Gorilla)
   Giraffe := animal{"Tim", "3 yrs old", "9am", "5pm"}
   fmt.Println(Giraffe)
}