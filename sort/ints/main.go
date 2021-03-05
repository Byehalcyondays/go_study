package main

import (
	"fmt"
	"sort"
)

type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func(hs HeroSlice)Len()int{
	return len(hs)
}

func(hs HeroSlice)Less(i,j int)bool{
	return hs[i].Age>hs[j].Age
}

func main(){
	var intSlice = []int{0,-1,10,7,90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}
