package main

import (
	"fmt"
	"log"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// bits count 4 3 2 1 0
// binary num 1 1 0 1 1

func main() {
	var target int64

	const onePos = 10
	err := setOneBit(onePos, &target)
	if err != nil {
		log.Fatalf("can't set one bit %s\n", err)
	}

	fmt.Printf("masked target with  one %2d's bit: (%#063b)\n", onePos, target)

	const zeroPos = 4
	target = int64(^uint(0) >> 1)
	err = setZeroBit(zeroPos, &target)
	if err != nil {
		log.Fatalf("can't set zero bit %s\n", err)
	}

	fmt.Printf("masked target with zero %2d's bit: (%#b)\n", zeroPos, target)

}

func setOneBit(i uint, target *int64) error {
	if i > 62 {
		return fmt.Errorf("i:%d should be in range [0,62]  \n", i)
	}
	*target |= 1 << i
	return nil
}

func setZeroBit(i uint, target *int64) error {
	if i > 62 {
		return fmt.Errorf("i:%d should be in range [0,62]  \n", i)
	}
	*target &^= 1 << i
	return nil
}
