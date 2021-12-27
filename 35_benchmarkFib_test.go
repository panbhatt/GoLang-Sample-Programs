// o test -bench=. 35_benchmarkFib_test.go
// go test -bench=.  -benchtime=4s -cpu=2,4,8 35_benchmarkFib_test.go

package main

import (
	"testing"
)

func Fib(n int, recursive bool) int {

	switch n {
	case 0:
		return 0
	case 1: 
		return 1
	default:

		if recursive {
			return Fib(n-1,true) + Fib(n-2, true)
		} else {
			a, b := 0,1
			for i:=1;i<n;i++ {
				a, b= b, a+b
			}
			return b
		}
	}



}

func BenchmarkFib20T(b *testing.B) {

	for i:=0;i<b.N;i++{
		Fib(20,true)
	}

}

func BenchmarkFib20F(b *testing.B) {
	for i:=0;i<b.N;i++{
		Fib(20,false)
	}
}
