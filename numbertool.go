package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	tool := flag.String("tool", "help", "Tool you want to use")
	number := flag.Int("number", 0, "Number to check")
	verbose := flag.Bool("v", false, "Show factors")
	flag.Parse()
	switch *tool {
	case "prime":
		isPrime := isPrime(*number)
		if isPrime {
			fmt.Printf("%d is a prime number\n", *number)
		} else {
			fmt.Printf("%d is not a prime number\n", *number)
			if *verbose {
				factors := factors(*number)
				fmt.Printf("The factors of %d are: %v\n", *number, factors)
			}
		}
	case "factors":
		factors := factors(*number)
		fmt.Printf("The factors of %d are: %v\n", *number, factors)
	case "help":
	default:
		fmt.Println("-number\n\tSpecify the number you are checking. Default: 0")
		fmt.Println("-tool\n\tSpecify the tool you want to use.")
		fmt.Println("\t\tAvailable tools: ")
		fmt.Println("\t\tprime\n\t\t\t checks if the number is prime")
		fmt.Println("\t\t -v\n\t\t\t Show factors")
		fmt.Println("\t\tfactors\n\t\t\t breaks the number into prime factors")
		fmt.Println("\t\thelp\n\t\t\t see help instructions")
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	base := int(math.Sqrt(float64(num)))
	for i := 2; i <= base; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func factors(num int) []int {
	if num < 0 {
		return append(factors(num*-1), -1)
	}
	if num < 2 {
		return []int{num}
	}
	var results []int
	curr := num
	for curr > 1 {
		for i := 2; i <= num; i++ {
			for curr%i == 0 {
				curr = curr / i
				results = append(results, i)
			}
		}
	}
	return results
}
