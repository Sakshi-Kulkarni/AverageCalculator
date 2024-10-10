package main

import (
	"fmt"
	"os"
	"GO-TRAINING/average"  
)

func main() {
	var nums float64

	calci := &average.Entities{}
	
	fmt.Println("Enter numbers: ",nums)

	for {
		_, err := fmt.Fscan(os.Stdin, &nums)
		if err != nil {
			break   
		}
		calci.Input(nums)  
	}

 	var n int
	var avg float64
	var err error 

	fmt.Println("Choose the average to calculate:")
	fmt.Println("1: Average of even numbers")
	fmt.Println("2: Average of odd numbers")
	fmt.Println("3: Average of all numbers")
	fmt.Scan(&n)

 	switch n {
	case 1:
		avg, err = calci.CalculateEvenAvg()
	case 2:
		avg, err = calci.CalculateOddAvg()
	case 3:
		avg, err = calci.CalculateAll()
	default:
		fmt.Fprintln(os.Stderr, "Invalid choice")  
		return
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)  
	} else {
		fmt.Printf("The average is %.2f\n", avg)  
	}
}
