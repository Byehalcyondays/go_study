package main

import "fmt"

func main() {
	var totallevel int = 9
	for i := 1; i <= totallevel; i++ {
		for j := 1; j <= totallevel-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == totallevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
