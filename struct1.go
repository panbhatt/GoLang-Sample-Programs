package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name string
	Number int
	Boss *Employee
	Hired time.Time
}

func main(){

	var emp Employee

	emp.Name = "Pankaj"
	emp.Number = 1
	emp.Hired = time.Now()

	fmt.Printf("%T %[1]v\n",emp)

	fmt.Printf("Employee -> %T %+[1]v\n",emp) // JSON like output


	boss := Employee {

		Name : "Stephen Neal",
		Number : 0, 
		Hired : time.Now(),
	}

	emp.Boss = &boss

	fmt.Printf("Boss -> %T %+[1]v\n",boss)
	fmt.Printf("Employee -> %T %+[1]v\n",emp)

	// Creating a map of the organ
	ization. 
	m := make(map[string]*Employee)

	m["bhatp48"] = &emp 
	m["sneal"] = &boss; 

	fmt.Printf("Whole Map -> %+V", m)

	// You can't get address of an STRUCT VALUE in a MAP


}