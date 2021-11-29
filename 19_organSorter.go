package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name string
	Weight int
}

type Organs []Organ

func (os Organs) Len() int { return len(os)}

func (os Organs) Swap(i int,j int) {
	os[i], os[j] = os[j], os[i]
}

func main() {

s := []Organ{{ "brain", 1340}, {"heart", 565}, { "dil", 444}}

fmt.Println(s) // no sorting here

sort.Sort(OrgansByName{s})
fmt.Println("Name Sorted -> ", s)
sort.Sort(OrgansByWeight{s})
fmt.Println("Weight Sorted -> ", s)


}

// Both Len and Swap function would be availabl ehere. 

type OrgansByName struct { Organs}
type OrgansByWeight struct { Organs}

func (on OrgansByName) Less(i,j int) bool {
	return on.Organs[i].Name < on.Organs[j].Name
}

func (on OrgansByWeight) Less(i,j int) bool {
	return on.Organs[i].Weight < on.Organs[j].Weight
}






