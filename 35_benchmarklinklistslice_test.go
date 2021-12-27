
// go test -bench=. -benchmem 35_benchmarklinklistslice_test.go 

package main

import (
	"testing"
)

type Node struct {
	val int
	next *Node
}

func insert(valTOBeInserted int , head *Node) *Node {

	tail := &Node{valTOBeInserted, nil}

	if head!=nil {
		head.next = tail
	}


	return tail

}

func makeList(n int) *Node {
	
	var head, tail *Node

	head = insert(0,nil)
	tail = insert(1,head)

	for i:=2;i<n;i++{
		tail= insert(n,tail)
	}

	return head
}

func sumList(h *Node) int {

	sum :=0;
	for h.next !=nil {
		sum += h.val
		h = h.next
	}

	return sum
}

func makeSlice(n int) []int{
	r := make([]int, n) 

	for i:=0;i<n;i++ {
		r[i] = i
	}

	return r
}



func sumSlice(sl []int)  int {

	sum := 0
	for _,v := range sl {
		sum += v
	}

	return sum
}

func BenchmarkList(b *testing.B) {
	for i:=0;i<b.N;i++{
		list:= makeList(2000)
		sumList(list)
	}
}


func BenchmarkSlice(b *testing.B) {
	for i:=0;i<b.N;i++{
		sl:= makeSlice(2000)
		sumSlice(sl)
	}
}

func BenchmarkListOnlySum(b *testing.B) {
	list:= makeList(2000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sumList(list)
	}
}


func BenchmarkSliceonlySum(b *testing.B) {
	sl:= makeSlice(2000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sumSlice(sl)
	}
}