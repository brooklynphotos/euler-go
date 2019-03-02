package main

// https://projecteuler.net/problem=20

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Number("96").multiple(uint16(3)))
	fmt.Println(factorial(10).sumDigits())
	fmt.Println(factorial(100).sumDigits())
}

func factorial(n uint16) (res Number) {
	res = "1"
	for i := uint16(1); i <= n; i++ {
		res = res.multiple(i)
	}
	return
}

type Number string

func (n Number) sumDigits() (sum int) {
	sum = 0
	for _, c := range n {
		d, _ := strconv.Atoi(string(c))
		sum += d
	}
	return
}

func (n Number) multiple(x uint16) Number {
	var newN Number
	newN = ""
	previousCarry := uint16(0)
	for i := len(n) - 1; i >= 0; i-- {
		d, _ := strconv.Atoi(string(n[i]))
		r := uint16(d)*x + previousCarry
		previousCarry = r / 10
		newD := r % uint16(10)
		add := strconv.Itoa(int(newD))
		newN = Number(add) + newN
	}
	if previousCarry > 0 {
		newN = Number(strconv.Itoa(int(previousCarry))) + newN
	}
	return newN
}
