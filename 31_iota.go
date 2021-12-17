package main

import (
	"fmt"
)

type shoe int 

const (
	tennis shoe = iota 
	racquet 
	nba 
	nhl
)

type Flags int

const (
	FlagUp Flags = 1 << iota
	FlagDown 
	FlagBroadCast
	FlagPointToPoint
	FlagMultiCast
)

type ByteSize int64

const (
	_ = iota
	KiB ByteSize = 1 << ( 10 * iota) // 2*10
	MiB   // 2 ^ 20
	GiB
	TiB
)

func main() {
	fmt.Println(tennis, racquet, nba , nhl)

	fmt.Println(FlagUp, FlagDown, FlagBroadCast, FlagPointToPoint,  FlagMultiCast)

	fmt.Println(KiB, MiB, GiB, TiB )
}