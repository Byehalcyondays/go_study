package main

import "fmt"

func main()  {
	//var level int = 1
	for i := 1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d x %d = %d        ",i,j,i*j)
		}
		fmt.Println()
	}
}
