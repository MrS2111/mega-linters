// Package main is the entry point of the program.
package main

import "fmt"

func main1() {
    var x int
    fmt.Println("Value of x:", x)

    // Attempting to divide by zero
    result := 10 / x
    fmt.Println("Result of division:", result)
}