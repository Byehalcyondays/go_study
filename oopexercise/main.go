package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student) say() string {
	info := fmt.Sprintf("student name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		s.name, s.gender, s.age, s.id, s.score)
	return info
}

func main() {
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     1,
		score:  18.6,
	}
	fmt.Println(stu.say())
}
