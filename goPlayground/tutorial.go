/******************************************************************************

David Vadnais go playground

Following:
*   https://www.w3schools.com/go/index.php

*******************************************************************************/

package main

import "fmt"

//constants of different data types
const PI float32 = 3.14 
const DEBUG bool = true
const RATE_DAVID_OUT_OF_TEN int = 10
const DAVIDS_NAME = "David"

// main
func main() {
    helloWorld()
    arraysAndMore()
    slicesNewConcept()
}

func helloWorld() {
    /////////////////// hello world Println
    var helloWorld string = "Hello World" //trpe declared
    fmt.Println(helloWorld)
    
    var helloWorldTwo  = "Hello World 2" //trpe infered
    fmt.Println(helloWorldTwo)
    
    helloWorldThree := "Hello World 3" // type infered 
    fmt.Println(helloWorldThree)
    
    /////////////////// print no \n
    fmt.Print(helloWorld)
    fmt.Print(helloWorld, "\n")
    
    var i_2,j_2 = 10,20

    fmt.Print(i_2, j_2, "\n")
    
    ///////////////////  printf
    var i string = "Hello"
    var j int = 15

    fmt.Printf("i has value: {%v} and type: {%T}\n", i, i)
    fmt.Printf("j has value: {%v} and type: {%T}", j, j) 
    
    
    /////////////////// numbers
    var a, b, c, d int = 7, 3, 3, 7

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}

func arraysAndMore() {
    /////////////////// arrays declared length
    var arr1 = [3]int{1,2,3}
    arr2 := [5]int{4,5,6,7,8}

    fmt.Println(arr1)
    fmt.Println(arr2)
    
    /////////////////// arrays infered length
    var arr3 = [...]int{1,2,3}
    arr4 := [...]int{4,5,6,7,8}

    fmt.Println(arr3)
    fmt.Println(arr4)
    
    /////////////////// access elements
    fmt.Println(arr4[0])
    fmt.Println(arr4[3])
    
    arr4[0] = 50
    fmt.Println(arr4[0])
    
    /////////////////// partial init
    arr5 := [5]int{} //not initialized
    arr6 := [5]int{1,2} //partially initialized
    arr7 := [5]int{1:10,2:40}//Initialize Only Specific Elements
    fmt.Println(arr5)
    fmt.Println(arr6)
    fmt.Println(arr7)
    
    /////////////////// len
    fmt.Println(len(arr7))
    
}

/**
"Slices are similar to arrays, but are more powerful and flexible.
Like arrays, slices are also used to store multiple values of the same type in a single variable.
However, unlike arrays, the length of a slice can grow and shrink as you see fit."

so like python hi = hello[]
**/
func slicesNewConcept() {
    /////////////////// basics
    myslice1 := []int{}
    myslice2 := []int{1,2,3}

    fmt.Println(len(myslice1))
    fmt.Println(cap(myslice1))
    fmt.Println(myslice1)

    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))
    fmt.Println(myslice2)
    
    /////////////////// Create a Slice From an Array
    arr1 := [6]int{10, 11, 12, 13, 14,15}
    myslice := arr1[2:4]

    fmt.Printf("myslice = %v\n", myslice)
    fmt.Printf("length = %d\n", len(myslice))
    fmt.Printf("capacity = %d\n", cap(myslice))
    
    /////////////////// Create a Slice With The make() Function
    myslice3 := make([]int, 5, 10)
    fmt.Printf("myslice3 = %v\n", myslice3)
    fmt.Printf("length = %d\n", len(myslice3))
    fmt.Printf("capacity = %d\n", cap(myslice3))

    // with omitted capacity
    myslice4 := make([]int, 5)
    fmt.Printf("myslice4 = %v\n", myslice4)
    fmt.Printf("length = %d\n", len(myslice4))
    fmt.Printf("capacity = %d\n", cap(myslice4))
    
    /////////////////// access/change an element same as array
    
    /////////////////// Append 
    // Append Elements To a Slice
    myslice5 := []int{1, 2, 3, 4, 5, 6}
    fmt.Printf("myslice5 = %v\n", myslice5)
    fmt.Printf("length = %d\n", len(myslice5))
    fmt.Printf("capacity = %d\n", cap(myslice5))
    
    myslice5 = append(myslice5, 20, 21)
    fmt.Printf("myslice5 = %v\n", myslice5)
    fmt.Printf("length = %d\n", len(myslice5))
    fmt.Printf("capacity = %d\n", cap(myslice5))
    
    //Append One Slice To Another Slice
    myslice6 := append(myslice5, myslice4...)
    
    fmt.Printf("myslice6=%v\n", myslice6)
    fmt.Printf("length=%d\n", len(myslice6))
    fmt.Printf("capacity=%d\n", cap(myslice6))
}
