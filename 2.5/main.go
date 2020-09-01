// +build ignore

package main

import popcount "gopl.exercise/2.4"

func main() {
	println(popcount.PopCountExpr(64))
	println(popcount.PopCountLoop(64))
	println(popcount.PopCountShift(64))
	println(popcount.PopCountBitwiseAnd(64))
}
