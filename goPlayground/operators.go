/******************************************************************************

David Vadnais go playground

Following:
*   https://www.w3schools.com/go/index.php

*******************************************************************************/

package main

import (
"fmt"
"time"
)

// main
func main() {
    goOperators()    
}


func goOperators() {
    /////////////////// basics
    /*  
        +, -, *, /, 
        % modulus
        ++, --  
    */
    var (
        sum1 = 100 + 50 // 150 (100 + 50)
        sum2 = sum1 + 250 // 400 (150 + 250)
        sum3 = sum2 + sum2 // 800 (400 + 400)
    )
    fmt.Println(sum3)
    
    /////////////////// Assignment
    /*
        =, +=, -=, *=, /=,
        %=,
        &=, |=, ^=
        
    */
    sum3 += 100
    fmt.Println("+=, ", sum3)
    
    //bit wise comparison 
    // 900 = 1110000100
    // 100 = 0001100100
    // 900 &= 100 = 100 
    // 100 to base10 = 4
    sum3 &= 100
    fmt.Println("&=, ", sum3)
    
    // sum2 = 400
    // 400 = 0011001000
    // 100 = 0001100100
    // 400 |= 100 = 0111110100
    sum2 |= 100
    fmt.Println("|=, ", sum2)
    
    // sum2 = 500 
    // 500 = 0111110100
    // 100 = 0001100100
    // 500 |= 100 = 0011001000
    sum2 ^= 100
    fmt.Println("^=, ", sum2)

    // Zero fill left shift
    // sum 2 == 400
    // 400 = 0011001000
    // 400 >> 3 = 0000110010
    sum2 >>= 3
    fmt.Println(">>=, ", sum2)
    
    // Signed right shift
    // sum 2 == 400
    // 400 =      0000000011001000
    // 400 << 3 = 0000110010000000
    sum2 = 400
    sum2 <<= 3
    fmt.Println("<<=, ", sum2)
    
    /////////////////// Comparison 
    /*
        >,<,==,!=,>=,<=
    */
    
    /////////////////// Logical
    /*
        &&, ||, !
    */
    
    /////////////////// bitwise
    /*
        $, |, 
        ^ xor, 
        << Zero fill left shift,
        >> Signed right shift
    */
    /////////////////// 
}
