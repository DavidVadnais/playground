/******************************************************************************

David Vadnais go playground

Following:
*   https://www.w3schools.com/go/index.php

*******************************************************************************/

package main

import (
"fmt"
)

// main
func main() {
    conditions()
}

func conditions() {
    /////////////////// basics
    /*
    Less than <
    Less than or equal <=
    Greater than >
    Greater than or equal >=
    Equal to ==
    Not equal to !=
    */
    
    /*
    Logical  &&
    Logical  ||
    Logical  !
    */
    
    /////////////////// if
    if 20 > 4 && 5 == 5 {
        fmt.Println("true")
    } else if ( 1 < 0 ) {
        
    } else {
        fmt.Println(20>5)
    }
    

    /////////////////// switch
    day := 2

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    }
    
    day = 99

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("the other thing")
    }
    
    day = 5

    switch day {
    case 1,3,5:
        fmt.Println("Odd weekday")
    case 2,4:
        fmt.Println("Even weekday")
    case 6,7:
        fmt.Println("Weekend")
    default:
        fmt.Println("Invalid day of day number")
    }
    
    /////////////////// loop
    //basics
    for i:=0; i < 5; i++ {
        if i == 2 {
            continue
        }
        
        fmt.Println(i)
        
        if i == 4 {
            break
        }
    }
    
    // range -- like for each
    fruits := [3]string{"apple", "orange", "banana"}
    for idx, val := range fruits {
        fmt.Printf("%v\t%v\n", idx, val)
    }
}
