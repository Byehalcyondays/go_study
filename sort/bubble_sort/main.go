package main

import "fmt"

func BubbleSort(arr *[]int)  {
	fmt.Println("排序前：",(*arr))
	temp:=0
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-1-i;j++{
			if(*arr)[j]>(*arr)[j+1]{
				temp = (*arr)[j]
				(*arr)[j]=(*arr)[j+1]
				(*arr)[j+1]=temp
			}
		}
	}
	fmt.Println("排序后：",(*arr))
}
func main()  {
	arr:=[]int{24,69,80,57,13}
	BubbleSort(&arr)

}
